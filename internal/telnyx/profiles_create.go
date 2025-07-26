package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type CreateVerifyProfileParams struct {
	*httpx.JsonBody
}

func NewCreateVerifyProfileParams(name string) *CreateVerifyProfileParams {
	p := &CreateVerifyProfileParams{httpx.NewJsonBody()}
	p.Set("name", name)
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
	err := params.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode json, %w", err)
	}
	req, err := c.newRequestWithBody("POST", "/verify_profiles", params)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request with body %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed to send: %w", err)
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
