package httpx

import (
	"io"
	"net/http"
	"net/url"
)

func Do(client *http.Client, request *http.Request) (*Response, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return WrapResponse(resp), nil
}

func SendRequestWithBody(method, url, contentType string, body io.Reader) (*Response, error) {
	req, err := NewRequestWithBody(method, url, contentType, body)
	if err != nil {
		return nil, err
	}
	return Do(http.DefaultClient, req)
}

func SendRequestWithParams(method, urlStr string, params url.Values) (*Response, error) {
	req, err := NewRequestWithParams(method, urlStr, params)
	if err != nil {
		return nil, err
	}
	return Do(http.DefaultClient, req)
}
