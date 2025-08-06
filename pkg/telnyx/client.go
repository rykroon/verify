package telnyx

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	apiToken   string
}

func NewClient(httpClient *http.Client, apiToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient, apiToken}
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	urlStr, err := url.JoinPath("https://api.telnyx.com/v2", path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	return req, nil
}
