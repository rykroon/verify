package telnyx

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListMessageTemplates() (json.RawMessage, error) {
	req, err := c.newRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
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
