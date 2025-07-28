package httpx

import (
	"io"
	"net/http"
	"net/url"
)

func NewRequest(method, url string, body BodyProvider) (*http.Request, error) {
	var reader io.Reader
	var contentType string
	if body != nil {
		reader = body.Reader()
		contentType = body.ContentType()
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", body.ContentType())
	}
	return req, nil
}

func SetQueryParams(req *http.Request, params url.Values) {
	req.URL.RawQuery = params.Encode()
}

func SetBearerToken(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
