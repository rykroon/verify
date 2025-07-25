package telnyx

import "fmt"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	resp, err := c.sendRequest("GET", path, nil)
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
	var result *VerificationProfileResponse
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
