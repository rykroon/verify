package utils

import (
	"io"
	"net/http"
)

type CachedResponse struct {
	*http.Response
	Body []byte
}

func (cr *CachedResponse) IsInformational() bool {
	return cr.StatusCode >= 100 && cr.StatusCode < 200
}

func (cr *CachedResponse) IsSuccess() bool {
	return cr.StatusCode >= 200 && cr.StatusCode < 300
}

func (cr *CachedResponse) IsRedirect() bool {
	return cr.StatusCode >= 300 && cr.StatusCode < 400
}

func (cr *CachedResponse) IsError() bool {
	return cr.StatusCode >= 400
}

func (cr *CachedResponse) IsClientError() bool {
	return cr.StatusCode >= 400 && cr.StatusCode < 500
}

func (cr *CachedResponse) IsServerError() bool {
	return cr.StatusCode >= 500
}

func (cr *CachedResponse) ContentType() string {
	return cr.Header.Get("Content-Type")
}

func (cr *CachedResponse) IsJson() bool {
	return cr.ContentType() == "application/json"
}

func DoAndReadAll(client *http.Client, req *http.Request) (*CachedResponse, error) {
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &CachedResponse{resp, data}, nil
}
