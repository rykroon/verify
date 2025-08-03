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
	httpClient   *http.Client
}

func NewClient(apiKeySid, apiKeySecret string) *Client {
	return &Client{apiKeySid, apiKeySecret, http.DefaultClient}
}

func (c *Client) SetHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) newRequest(method, path string, body httpx.RequestBodyProvider) (*http.Request, error) {
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

func (c *Client) do(req *http.Request) (*httpx.Body, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	respBody, err := httpx.ReadBodyFromResponse(resp)
	if err != nil {
		return nil, err
	}

	if httpx.IsError(resp) {
		return nil, fmt.Errorf("http server error %d, %s", resp.StatusCode, respBody.ToString())
	}

	return respBody, nil
}
