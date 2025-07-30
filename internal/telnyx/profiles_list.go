package telnyx

import (
	"fmt"
)

func (c *Client) ListVerifyProfiles() (*ListResponse[VerificationProfile], error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if !respBody.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result *ListResponse[VerificationProfile]
	err = respBody.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
