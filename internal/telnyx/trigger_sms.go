package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type TriggerSmsVerificationParams struct {
	*httpx.JsonBodyBuilder
}

func NewTriggerSmsVerificationParams(phoneNumber, verifyProfileId string) *TriggerSmsVerificationParams {
	return &TriggerSmsVerificationParams{
		httpx.NewJsonBodyBuilder().Set(
			"phone_number", phoneNumber,
		).Set(
			"verify_profile_id", verifyProfileId,
		),
	}
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
	resp, err := c.sendRequest("POST", "/verifications/sms", params)
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
	var result *VerificationResponse
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
