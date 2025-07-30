package httpx

import (
	"bytes"
	"io"
)

type OctetStream []byte

func (s OctetStream) ContentType() string {
	return "application/octet-stream"
}

func (s OctetStream) Reader() io.Reader {
	return bytes.NewReader(s)
}

func (s OctetStream) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(s)
	return int64(n), err
}
