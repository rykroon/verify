package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

func (c *Client) VerifyCode(verificationId, code string) (map[string]any, error) {
	params := httpx.NewJsonBody()
	params.Set("code", code)
	err := params.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode json %w", err)
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.newRequestWithBody("POST", path, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result map[string]any
	if err = resp.ToJson(&result); err != nil {
		return nil, err
	}

	return result, nil
}
