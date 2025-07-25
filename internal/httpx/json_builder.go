package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonBodyBuilder struct {
	data map[string]any
}

func NewJsonBodyBuilder() *JsonBodyBuilder {
	return &JsonBodyBuilder{data: make(map[string]any)}
}

func (b *JsonBodyBuilder) Set(k string, v any) *JsonBodyBuilder {
	b.data[k] = v
	return b
}

func (b *JsonBodyBuilder) SetPath(k string, v any) *JsonBodyBuilder {
	parts := strings.Split(k, ".")

	current := b.data
	for i, part := range parts {
		if i == len(parts)-1 {
			current[part] = v
			return b
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
	return b
}

func (b *JsonBodyBuilder) ToReader() (io.Reader, error) {
	bytes_, err := json.Marshal(b.data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bytes_), nil
}

func (b *JsonBodyBuilder) ContentType() string {
	return "application/json"
}
