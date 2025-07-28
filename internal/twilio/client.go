package twilio

import (
	"fmt"
	"net/http"
	neturl "net/url"

	"github.com/rykroon/verify/internal/httpx"
)

type Client struct {
	apiKeySid    string
	apiKeySecret string
}

func NewClient(apiKeySid, apiKeySecret string) *Client {
	return &Client{apiKeySid: apiKeySid, apiKeySecret: apiKeySecret}
}

func (c *Client) newRequest(method, path string, body httpx.BodyProvider) (*http.Request, error) {
	url, err := neturl.JoinPath("https://verify.twilio.com/v2/", path)
	if err != nil {
		return nil, err
	}
	req, err := httpx.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKeySid, c.apiKeySecret)
	return req, nil
}

func (c *Client) do(req *http.Request) (*httpx.Response, error) {
	resp, err := httpx.Do(http.DefaultClient, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("http server error %d, %s", resp.StatusCode, respBody.ToString())
	}

	return resp, nil
}
