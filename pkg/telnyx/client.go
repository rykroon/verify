package telnyx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiToken   string
	httpClient *http.Client
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken, http.DefaultClient}
}

func (c *Client) SetHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	return c.newRequest(method, path, body)
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
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

func (c *Client) handleResponse(resp *http.Response) (json.RawMessage, error) {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var result json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
