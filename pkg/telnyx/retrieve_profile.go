package telnyx

import (
	"encoding/json"
	"fmt"
)

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (json.RawMessage, error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
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
