package client

import (
	"bufio"
	"context"
	"fmt"
	"net/http/httptest"
	"net/textproto"
	"strings"

	"github.com/goccy/go-json"
)

type SSE struct {
	Close func() error
	Next  func(response any) error
}

type SSEResponse struct {
	Data       any             `json:"data"`
	Label      string          `json:"label"`
	Path       []any           `json:"path"`
	HasNext    bool            `json:"hasNext"`
	Errors     json.RawMessage `json:"errors"`
	Extensions map[string]any  `json:"extensions"`
}

func errorSSE(err error) *SSE {
	return &SSE{
		Close: func() error { return nil },
		Next: func(response any) error {
			return err
		},
	}
}

func (p *Client) SSE(ctx context.Context, query string, options ...Option) *SSE {
	r, err := p.newRequest(query, options...)
	if err != nil {
		return errorSSE(fmt.Errorf("request: %w", err))
	}
	r = r.WithContext(ctx)

	r.Header.Set("Accept", "text/event-stream")
	r.Header.Set("Cache-Control", "no-cache")
	r.Header.Set("Connection", "keep-alive")

	srv := httptest.NewServer(p.h)
	w := httptest.NewRecorder()
	p.h.ServeHTTP(w, r)

	reader := textproto.NewReader(bufio.NewReader(w.Body))
	line, err := reader.ReadLine()
	if err != nil {
		return errorSSE(fmt.Errorf("response: %w", err))
	}
	if line != ":" {
		return errorSSE(fmt.Errorf("expected :, got %s", line))
	}

	return &SSE{
		Close: func() error {
			srv.Close()
			return nil
		},
		Next: func(response any) error {
			for {
				line, err := reader.ReadLine()
				if err != nil {
					return err
				}
				kv := strings.SplitN(line, ": ", 2)

				switch kv[0] {
				case "":
					continue
				case "event":
					switch kv[1] {
					case "next":
						continue
					case "complete":
						return nil
					default:
						return fmt.Errorf("expected event type: %#v", kv[1])
					}
				case "data":
					var respDataRaw SSEResponse
					if err = json.Unmarshal([]byte(kv[1]), &respDataRaw); err != nil {
						return fmt.Errorf("decode: %w", err)
					}

					// we want to unpack even if there is an error, so we can see partial responses
					unpackErr := unpack(respDataRaw, response, p.dc)

					if respDataRaw.Errors != nil {
						return RawJsonError{respDataRaw.Errors}
					}

					return unpackErr
				default:
					return fmt.Errorf("unexpected sse field %s", kv[0])
				}
			}
		},
	}
}
