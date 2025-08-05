package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) NewRetrieveVerifyProfileRequest(verifyProfileId string) (*http.Request, error) {
	req, err := c.NewRequest("GET", "verify_profiles/"+verifyProfileId, nil)
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

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
