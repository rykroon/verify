package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type MarshalFunc func(any) ([]byte, error)
type UnmarshalFunc func([]byte, any) error

type Body struct {
	data        []byte
	contentType string
}

func NewBody(data []byte, contentType string) *Body {
	return &Body{data, contentType}
}

func NewBodyUsingMarshal(v any, fn MarshalFunc, contentType string) (*Body, error) {
	data, err := fn(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal value: %w", err)
	}
	return NewBody(data, contentType), nil
}

func NewJsonBody(v any) (*Body, error) {
	return NewBodyUsingMarshal(v, json.Marshal, "application/json")
}

func (b *Body) Reader() io.Reader {
	return bytes.NewReader(b.data)
}

func (b *Body) ContentType() string {
	return b.contentType
}

func (b *Body) ToString() string {
	return string(b.data)
}

func (b *Body) UnmarshalWith(v any, unmarshalFunc UnmarshalFunc) error {
	return unmarshalFunc(b.data, v)
}

func (b *Body) UnmarshalJson(v any) error {
	return b.UnmarshalWith(v, json.Unmarshal)
}

func (r *Response) IsJson() bool {
	return r.ContentType() == "application/json"
}
