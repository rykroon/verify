package httpx

import (
	"io"
	"net/url"
	"strings"
)

type FormBodyBuilder struct {
	Values url.Values
}

func NewFormBodyBuilder() *FormBodyBuilder {
	return &FormBodyBuilder{Values: url.Values{}}
}

func (b *FormBodyBuilder) Add(k, v string) *FormBodyBuilder {
	b.Values.Add(k, v)
	return b
}

func (b *FormBodyBuilder) Set(k, v string) *FormBodyBuilder {
	b.Values.Set(k, v)
	return b
}

func (b *FormBodyBuilder) ToReader() (io.Reader, error) {
	return strings.NewReader(b.Values.Encode()), nil
}

func (b *FormBodyBuilder) ContentType() string {
	return "application/x-www-form-urlencoded"
}
