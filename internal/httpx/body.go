package httpx

import "io"

type BodyEncoder interface {
	Reader() io.Reader
	ContentType() string
}
