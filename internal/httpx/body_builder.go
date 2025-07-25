package httpx

import "io"

type BodyBuilder interface {
	ToReader() (io.Reader, error)
	ContentType() string
}
