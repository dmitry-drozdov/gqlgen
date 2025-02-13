package graphql

import (
	"fmt"
	"io"

	"github.com/goccy/go-json"
)

func MarshalMap(val map[string]any) Marshaler {
	return WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalMap(v any) (map[string]any, error) {
	if m, ok := v.(map[string]any); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}
