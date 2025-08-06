package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type RetrieveProfileParams struct {
	VerifyProfileId string `json:"verify_profile_id"`
}

func (c *Client) RetrieveVerifyProfile(params *RetrieveProfileParams) (*utils.CachedResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	req, err := c.NewRequest("GET", "verify_profiles/"+params.VerifyProfileId, nil)
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
