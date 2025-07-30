package httpx

import (
	"io"
	"net/url"
	"strings"
)

type FormBody struct {
	url.Values
}

func NewFormBody() FormBody {
	return FormBody{Values: url.Values{}}
}

func (b FormBody) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (b FormBody) Reader() io.Reader {
	return strings.NewReader(b.Encode())
}

func (b FormBody) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(b.Encode()))
	return int64(n), err
}
