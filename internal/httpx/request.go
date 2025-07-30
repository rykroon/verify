package httpx

import (
	"io"
	"net/http"
	"net/url"
)

type RequestBodyProvider interface {
	Reader() io.Reader
	ContentType() string
}

func ReadBodyFromRequest(req *http.Request) (*Body, error) {
	return ReadBody(req.Header.Get("Content-Type"), req.Body)
}

func NewRequest(method, url string, provider RequestBodyProvider) (*http.Request, error) {
	var body io.Reader
	var contentType string
	if provider != nil {
		body = provider.Reader()
		contentType = provider.ContentType()
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}

func SetQueryParams(req *http.Request, params url.Values) {
	req.URL.RawQuery = params.Encode()
}

func SetBearerToken(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
