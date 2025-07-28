package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

func (c *Client) TriggerSmsVerification(phoneNumber, verifyProfileId string) (*VerificationResponse, error) {
	m := map[string]any{"phone_number": phoneNumber, "verify_profile_id": verifyProfileId}
	body, err := httpx.NewJsonBody(m)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	req, err := c.newRequest("POST", "/verifications/sms", body)
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
	var result *VerificationResponse
	if err = resp.ToJson(&result); err != nil {
		return nil, err
	}

	return result, nil
}
