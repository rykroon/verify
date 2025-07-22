package telnyx

import (
	"net/http"

	"github.com/rykroon/verify/internal/params"
)

type TriggerSmsVerificationParams struct {
	*params.ParamBuilder
}

func NewTriggerSmsVerificationParams(phoneNumber, verifyProfileId string) *TriggerSmsVerificationParams {
	return &TriggerSmsVerificationParams{
		params.NewParamBuilder().Set(
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
	body, err := params.ToReader()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verifications/sms", body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result *VerificationResponse
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
