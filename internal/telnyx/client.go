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

func (c *Client) sendRequest(method, path string, body httpx.BodyEncoder) (*httpx.Response, error) {
	urlStr, err := url.JoinPath("https://api.telnyx.com/v2", path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if body != nil {
		req, err = httpx.NewRequestWithBody(method, urlStr, body.ContentType(), body.Reader())
	} else {
		req, err = httpx.NewRequestWithParams(method, urlStr, nil)
	}

	if err != nil {
		return nil, err
	}

	httpx.SetBearerToken(req, c.apiToken)

	resp, err := httpx.Do(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("http server error %d, %s", resp.StatusCode, string(respBody))
	}

	return resp, nil
}
