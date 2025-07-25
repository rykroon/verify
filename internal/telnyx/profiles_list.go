package telnyx

import (
	"fmt"
)

func (c *Client) ListVerifyProfiles() (*VerificationProfileListResponse, error) {
	resp, err := c.sendRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, err
	}

	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result *VerificationProfileListResponse
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
