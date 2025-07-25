package httpx

import (
	"io"
	"net/url"
	"strings"
)

type FormBody struct {
	url.Values
}

func NewFormBody() *FormBody {
	return &FormBody{Values: url.Values{}}
}

func (b *FormBody) Reader() io.Reader {
	return strings.NewReader(string(b.Values.Encode()))
}

func (b *FormBody) ContentType() string {
	return "application/x-www-form-urlencoded"
}
