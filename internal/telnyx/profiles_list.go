package telnyx

import (
	"fmt"
)

func (c *Client) ListVerifyProfiles() (*VerificationProfileListResponse, error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result *VerificationProfileListResponse
	err = respBody.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
