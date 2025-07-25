package telnyx

import "fmt"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	resp, err := c.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
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
