package utils

import (
	"io"
	"net/http"
)

type CachedResponse struct {
	*http.Response
	Body []byte
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

type HttpError struct {
	StatusCode  int
	ContentType string
	Body        []byte
}

func IsInformational(resp *http.Response) bool {
	return resp.StatusCode >= 100 && resp.StatusCode < 200
}

func IsSuccess(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func IsRedirect(resp *http.Response) bool {
	return resp.StatusCode >= 300 && resp.StatusCode < 400
}

func IsClientError(resp *http.Response) bool {
	return resp.StatusCode >= 400 && resp.StatusCode < 500
}

func IsServerError(resp *http.Response) bool {
	return resp.StatusCode >= 500
}

func IsError(resp *http.Response) bool {
	return resp.StatusCode >= 400
}
