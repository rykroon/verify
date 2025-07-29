package telnyx

import "fmt"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequest("GET", path, nil)
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

	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result *VerificationProfileResponse
	err = respBody.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
