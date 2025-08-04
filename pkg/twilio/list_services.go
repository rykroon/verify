package twilio

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListServices() (json.RawMessage, error) {
	req, err := c.newRequest("GET", "/Services", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}

func (c *Client) FetchService(sid string) (json.RawMessage, error) {
	req, err := c.newRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
