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

func (c *Client) newRequest(method, path string, body httpx.RequestBodyProvider) (*http.Request, error) {
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

func (c *Client) do(req *http.Request) (*httpx.Body, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := httpx.ReadBodyFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	if httpx.IsServerError(resp) {
		return nil, fmt.Errorf("http error: %d, %s", resp.StatusCode, body.ToString())
	}
	if httpx.IsClientError(resp) {
		var telnyxErrorResp TelnyxErrorResponse
		if err := body.UnmarshalJson(&telnyxErrorResp); err != nil {
			return nil, fmt.Errorf("failed to decode http error as json: %w", err)
		}
		return nil, &telnyxErrorResp.Errors[0]
	}

	return body, nil
}
