package telnyx

import (
	"net/http"

	"github.com/rykroon/verify/internal/httpx"
)

func (c *Client) ListVerifyProfiles() (*VerificationProfileListResponse, error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var result *VerificationProfileListResponse
	err = httpx.HandleResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
