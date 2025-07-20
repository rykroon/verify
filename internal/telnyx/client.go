package telnyx

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url, err := neturl.JoinPath("https://api.telnyx.com/v2", path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	return req, nil
}
