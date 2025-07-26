package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type TriggerSmsVerificationParams struct {
	*httpx.JsonBody
}

func NewTriggerSmsVerificationParams(phoneNumber, verifyProfileId string) *TriggerSmsVerificationParams {
	p := &TriggerSmsVerificationParams{httpx.NewJsonBody()}
	p.Set("phone_number", phoneNumber)
	p.Set("verify_profile_id", verifyProfileId)
	return p
}

func (p *TriggerSmsVerificationParams) SetCustomCode(customCode string) *TriggerSmsVerificationParams {
	p.Set("custom_code", customCode)
	return p
}

func (p *TriggerSmsVerificationParams) SetTimeoutSecs(timeoutSecs string) *TriggerSmsVerificationParams {
	p.Set("timeout_secs", timeoutSecs)
	return p
}

func (c *Client) TriggerSmsVerification(params *TriggerSmsVerificationParams) (*VerificationResponse, error) {
	err := params.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode json %w", err)
	}
	req, err := c.newRequestWithBody("POST", "/verifications/sms", params)
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
