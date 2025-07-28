package httpx

import (
	"io"
	"net/http"
	"net/url"
)

func NewRequest(method, url string, body RequestBody) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		reader = body.Reader()
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if body != nil {
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
