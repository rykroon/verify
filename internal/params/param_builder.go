package params

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type ParamBuilder struct {
	data map[string]any
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{
		data: make(map[string]any),
	}
}

func (pb *ParamBuilder) ToReader() (io.Reader, error) {
	b, err := json.Marshal(pb.data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (pb *ParamBuilder) Set(key string, value any) *ParamBuilder {
	parts := strings.Split(key, ".")

	current := pb.data
	for i, part := range parts {
		if i == len(parts)-1 {
			current[part] = value
			return pb
		}

		// Traverse or create nested map
		if next, ok := current[part]; ok {
			if m, ok := next.(map[string]any); ok {
				current = m
			} else {
				// Overwrite if not a map
				newMap := make(map[string]any)
				current[part] = newMap
				current = newMap
			}
		} else {
			newMap := make(map[string]any)
			current[part] = newMap
			current = newMap
		}
	}
	return pb
}
