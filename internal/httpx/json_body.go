package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonBody struct {
	data map[string]any
	buf  *bytes.Buffer
}

func NewJsonBody() *JsonBody {
	return &JsonBody{data: make(map[string]any)}
}

func (b *JsonBody) Set(k string, v any) {
	b.data[k] = v
}

func (b *JsonBody) SetPath(k string, v any) {
	parts := strings.Split(k, ".")

	current := b.data
	for i, part := range parts {
		if i == len(parts)-1 {
			current[part] = v
			return
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
}

func (b *JsonBody) Encode() error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(b.data); err != nil {
		return err
	}
	b.buf = buf
	return nil
}

func (b *JsonBody) Reader() io.Reader {
	if b.buf == nil {
		panic("JsonBody must be encoded before use")
	}
	return b.buf
}

func (b *JsonBody) ContentType() string {
	return "application/json"
}
