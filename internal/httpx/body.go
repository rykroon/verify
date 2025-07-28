package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type Serializer func(any) ([]byte, error)
type Deserializer func([]byte, any) error

type BodyProvider interface {
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

func (b *Body) ToString() string {
	return string(b.data)
}

func (b *Body) DeserializeWith(v any, deserializer Deserializer) error {
	return deserializer(b.data, v)
}

func (b *Body) ToJson(v any) error {
	return b.DeserializeWith(v, json.Unmarshal)
}

func NewBody(data []byte, contentType string) *Body {
	return &Body{data, contentType}
}

func NewBodyFromSerializer(v any, serializer Serializer, contentType string) (*Body, error) {
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
	return strings.NewReader(b.Encode())
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
