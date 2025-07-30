package httpx

import (
	"bytes"
	"fmt"
	"io"
)

type Body struct {
	data        []byte
	contentType string
}

func NewBody(data []byte, contentType string) *Body {
	return &Body{data, contentType}
}

func (b *Body) ContentType() string {
	return b.contentType
}

func (b *Body) Reader() io.Reader {
	return bytes.NewReader(b.data)
}

func (b *Body) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.data)
	return int64(n), err
}

func (b *Body) ToString() string {
	return string(b.data)
}

func ReadBody(contentType string, reader io.ReadCloser) (*Body, error) {
	if reader == nil {
		return nil, fmt.Errorf("body is nil")
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &Body{
		data:        data,
		contentType: contentType,
	}, nil
}
