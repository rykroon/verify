package telnyx

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rykroon/verify/internal/utils"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken}
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

func checkResponse(cr *utils.CachedResponse) error {
	if utils.IsError(cr.StatusCode) {
		return fmt.Errorf("Telnyx error: %d, %s", cr.StatusCode, string(cr.Body))
	} else if !utils.IsSuccess(cr.StatusCode) {
		return fmt.Errorf("unexpected status code: %d", cr.StatusCode)
	}
	return nil
}
