package telnyx

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/rykroon/verify/internal/httpx"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c *Client) newRequest(method, path string, body httpx.BodyProvider) (*http.Request, error) {
	urlStr, err := url.JoinPath("https://api.telnyx.com/v2", path)
	if err != nil {
		return nil, err
	}

	req, err := httpx.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	httpx.SetBearerToken(req, c.apiToken)
	return req, nil
}

func (c *Client) do(req *http.Request) (*httpx.Response, error) {
	resp, err := httpx.Do(http.DefaultClient, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.IsError() {
		body, err := resp.ReadBody()
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %w", err)
		}
		if resp.IsServerError() {
			return nil, fmt.Errorf("http server error %d, %s", resp.StatusCode, body.ToString())
		}
		if resp.IsClientError() {
			return nil, fmt.Errorf("http client error %d, %s", resp.StatusCode, body.ToString())
		}
	}

	return resp, nil
}
