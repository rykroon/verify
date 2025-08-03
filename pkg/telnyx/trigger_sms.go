package telnyx

import (
	"fmt"

	ds "github.com/rykroon/verify/internal/data_structures"
	"github.com/rykroon/verify/internal/httpx"
)

type triggerSmsVerificationParams struct {
	*ds.ParamBuilder
}

func NewTriggerSmsVerificationParams() triggerSmsVerificationParams {
	return triggerSmsVerificationParams{ds.NewParamBuilder()}
}

func (p *triggerSmsVerificationParams) SetPhoneNumber(phoneNumber string) {
	p.Set("phone_number", phoneNumber)
}

func (p *triggerSmsVerificationParams) SetVerifyProfileId(verifyProfileId string) {
	p.Set("verify_profile_id", verifyProfileId)
}

func (c *Client) TriggerSmsVerification(params triggerSmsVerificationParams) (*DetailResponse[Verification], error) {
	body, err := httpx.NewJsonBody(params)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	req, err := c.newRequest("POST", "/verifications/sms", body)
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
	var result *DetailResponse[Verification]
	if err = respBody.UnmarshalJson(&result); err != nil {
		return nil, err
	}
	return result, nil
}
