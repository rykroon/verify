package httpx

import (
	"io"
	"net/http"
	"net/url"
)

func NewRequestWithBody(method, url, contentType string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func NewRequestWithParams(method, urlStr string, params url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}
	return req, nil
}

func SetBearerToken(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
