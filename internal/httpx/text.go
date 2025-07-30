package httpx

import (
	"io"
	"strings"
)

type PlainText string

func (pt PlainText) ContentType() string {
	return "text/plain"
}

func (pt PlainText) Reader() io.Reader {
	return strings.NewReader(string(pt))
}

func (pt PlainText) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(pt))
	return int64(n), err
}
