package telnyx

import "fmt"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequestWithParams("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}
	var result *VerificationProfileResponse
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
