package graphql

import (
	"io"

	"github.com/goccy/go-json"
)

func MarshalAny(v any) Marshaler {
	return WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalAny(v any) (any, error) {
	return v, nil
}
