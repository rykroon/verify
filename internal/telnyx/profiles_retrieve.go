package telnyx

import "fmt"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*DetailResponse[VerificationProfile], error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	if !respBody.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result *DetailResponse[VerificationProfile]
	err = respBody.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
