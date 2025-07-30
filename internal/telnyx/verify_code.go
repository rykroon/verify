package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

func (c *Client) VerifyCode(verificationId, code string) (map[string]any, error) {
	m := map[string]any{"code": code}
	body, err := httpx.NewJsonBody(m)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.newRequest("POST", path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	if !respBody.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}
	var result map[string]any
	if err = respBody.UnmarshalJson(&result); err != nil {
		return nil, err
	}

	return result, nil
}
