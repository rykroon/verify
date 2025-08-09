package twilio

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient   *http.Client
	apiKeySid    string
	apiKeySecret string
}

func NewClient(client *http.Client, apiKeySid, apiKeySecret string) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{client, apiKeySid, apiKeySecret}
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	urlStr, err := url.JoinPath("https://verify.twilio.com/v2/", path)
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
	req.SetBasicAuth(c.apiKeySid, c.apiKeySecret)
	return req, nil
}
