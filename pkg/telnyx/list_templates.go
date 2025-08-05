package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) NewListMessageTemplatesRequest() (*http.Request, error) {
	req, err := c.newRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (c *Client) ListMessageTemplates() (json.RawMessage, error) {
	req, err := c.NewListMessageTemplatesRequest()
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
