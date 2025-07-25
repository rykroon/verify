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
	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result map[string]any
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
