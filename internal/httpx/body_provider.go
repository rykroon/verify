package httpx

import (
	"io"
	"net/url"
	"strings"
)

type BodyProvider interface {
	Reader() io.Reader
	ContentType() string
}

type FormBody struct {
	url.Values
}

func NewFormBody() FormBody {
	return FormBody{Values: url.Values{}}
}

func (b FormBody) Reader() io.Reader {
	return strings.NewReader(b.Encode())
}

func (b FormBody) ContentType() string {
	return "application/x-www-form-urlencoded"
}
