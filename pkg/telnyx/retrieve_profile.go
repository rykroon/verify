package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) NewRetrieveVerifyProfileRequest(verifyProfileId string) (*http.Request, error) {
	req, err := c.newRequest("GET", "verify_profiles/"+verifyProfileId, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, err
}

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (json.RawMessage, error) {
	req, err := c.NewRetrieveVerifyProfileRequest(verifyProfileId)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
