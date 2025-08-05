package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rykroon/verify/internal/utils"
)

type Client struct {
	apiKeySid    string
	apiKeySecret string
}

func NewClient(apiKeySid, apiKeySecret string) *Client {
	return &Client{apiKeySid, apiKeySecret}
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

func (c *Client) handleResponse(resp *utils.CachedResponse) (json.RawMessage, error) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(resp.Body))
	}

	var result json.RawMessage

	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
