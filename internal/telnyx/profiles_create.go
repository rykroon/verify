package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type CreateVerifyProfileParams struct {
	*httpx.JsonBodyBuilder
}

func NewCreateVerifyProfileParams(name string) *CreateVerifyProfileParams {
	p := &CreateVerifyProfileParams{httpx.NewJsonBodyBuilder().Set("name", name)}
	return p
}

func (p *CreateVerifyProfileParams) SetSmsMessagingTemplateId(templateId string) *CreateVerifyProfileParams {
	p.SetPath("sms.messaging_template_id", templateId)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsAppName(appName string) *CreateVerifyProfileParams {
	p.SetPath("sms.app_name", appName)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsCodeLength(codeLength string) *CreateVerifyProfileParams {
	p.SetPath("sms.code_length", codeLength)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsWhitelistedDestinations(destinations []string) *CreateVerifyProfileParams {
	p.SetPath("sms.whitelisted_destinations", destinations)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsDefaultVerificationTimeoutSecs(v int) *CreateVerifyProfileParams {
	p.SetPath("sms.default_verification_timeout_secs", v)
	return p
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params *CreateVerifyProfileParams) (*VerificationProfileResponse, error) {
	resp, err := c.sendRequest("POST", "/verify_profiles", params)
	if err != nil {
		return nil, err
	}

	var result *VerificationProfileResponse
	if !resp.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}
	err = resp.ToJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
