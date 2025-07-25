package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

func (c *Client) VerifyCode(verificationId, code string) (map[string]any, error) {
	params := httpx.NewJsonBodyBuilder().Set("code", code)
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	resp, err := c.sendRequest("POST", path, params)
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

	var result map[string]any
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
