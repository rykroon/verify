package telnyx

import (
	"fmt"
)

func (c *Client) ListVerifyProfiles() (*VerificationProfileListResponse, error) {
	resp, err := c.sendRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		if resp.IsServerError() {
			return nil, fmt.Errorf("server error")
		} else if resp.IsClientError() {
			return nil, fmt.Errorf("Telnyx Error: %s", string(resp.BodyBytes))
		}
	}

	var result *VerificationProfileListResponse
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
