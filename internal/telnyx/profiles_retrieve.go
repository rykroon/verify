package telnyx

import "net/http"

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var result *VerificationProfileResponse
	err = c.processResponse(resp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
