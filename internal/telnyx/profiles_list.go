package telnyx

import "net/http"

func (c *Client) ListVerifyProfiles() (*VerificationProfileListResponse, error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var result *VerificationProfileListResponse
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
