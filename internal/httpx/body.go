package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type RequestBody interface {
	Reader() io.Reader
	ContentType() string
}

type Body struct {
	data        []byte
	contentType string
}

func (b *Body) Reader() io.Reader {
	return bytes.NewReader(b.data)
}

func (b *Body) ContentType() string {
	return b.contentType
}

func NewBody(data []byte, contentType string) *Body {
	return &Body{data, contentType}
}

func NewBodyFromSerializer(v any, serializer func(any) ([]byte, error), contentType string) (*Body, error) {
	data, err := serializer(v)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize value: %w", err)
	}
	return NewBody(data, contentType), nil
}

func NewJsonBody(v any) (*Body, error) {
	return NewBodyFromSerializer(v, json.Marshal, "application/json")
}

// Form body
type FormBody struct {
	url.Values
}

func NewFormBody() *FormBody {
	return &FormBody{Values: url.Values{}}
}

func (b *FormBody) Reader() io.Reader {
	return strings.NewReader(string(b.Encode()))
}

func (b *FormBody) ContentType() string {
	return "application/x-www-form-urlencoded"
}

// binary body

type OctetStream []byte

func (b OctetStream) Reader() io.Reader {
	return bytes.NewReader(b)
}

func (b OctetStream) ContentType() string {
	return "application/octet-stream"
}

// plain text body
type PlainText string

func (t PlainText) Reader() io.Reader {
	return strings.NewReader(string(t))
}

func (t PlainText) ContentType() string {
	return "text/plain"
}
