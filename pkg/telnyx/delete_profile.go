package telnyx

import (
	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) DeleteVerifyProfile(verifyProfileId string) (*utils.CachedResponse, error) {
	req, err := c.NewRequest("DELETE", "verify_profiles/"+verifyProfileId, nil)
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
