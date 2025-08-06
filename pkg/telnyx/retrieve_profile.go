package telnyx

import (
	"fmt"
	"net/http"
)

func (c *Client) NewRetrieveVerifyProfileRequest(verifyProfileId string) (*http.Request, error) {
	req, err := c.NewRequest("GET", "verify_profiles/"+verifyProfileId, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, err
}
