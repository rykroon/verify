package telnyx

import (
	"net/http"

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
	p.Set("sms.messaging_template_id", templateId)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsAppName(appName string) *CreateVerifyProfileParams {
	p.Set("sms.app_name", appName)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsCodeLength(codeLength string) *CreateVerifyProfileParams {
	p.Set("sms.code_length", codeLength)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsWhitelistedDestinations(destinations []string) *CreateVerifyProfileParams {
	p.Set("sms.whitelisted_destinations", destinations)
	return p
}

func (p *CreateVerifyProfileParams) SetSmsDefaultVerificationTimeoutSecs(v int) *CreateVerifyProfileParams {
	p.Set("sms.default_verification_timeout_secs", v)
	return p
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params *CreateVerifyProfileParams) (*VerificationProfileResponse, error) {
	body, err := params.ToReader()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verify_profiles", body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result *VerificationProfileResponse
	err = httpx.HandleResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
