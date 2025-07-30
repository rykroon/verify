package telnyx

import "fmt"

func (c *Client) ListMessageTemplates() (map[string]any, error) {
	req, err := c.newRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	return result, nil
}
