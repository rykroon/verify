package utils

import (
	"fmt"
	"io"
	"net/http"
)

type Content struct {
	Data []byte
	Type string
}

func SendRequest(client *http.Client, req *http.Request) (*Content, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http error: status code: %d, response body: %s", resp.StatusCode, string(data))
	}

	return &Content{Data: data, Type: resp.Header.Get("Content-Type")}, nil
}
