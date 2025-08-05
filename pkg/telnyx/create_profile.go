package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

type createVerifyProfileParams struct {
	utils.ParamBuilder
}

func NewCreateVerifyProfileParams() *createVerifyProfileParams {
	return &createVerifyProfileParams{*utils.NewParamBuilder()}
}

func (p *createVerifyProfileParams) SetName(name string) {
	p.Set("name", name)
}

func (p *createVerifyProfileParams) SetSmsMessagingTemplateId(templateId string) {
	p.SetPath("sms.messaging_template_id", templateId)
}

func (p *createVerifyProfileParams) SetSmsAppName(appName string) {
	p.SetPath("sms.app_name", appName)
}

func (p *createVerifyProfileParams) SetSmsCodeLength(codeLength string) {
	p.SetPath("sms.code_length", codeLength)
}

func (p *createVerifyProfileParams) SetSmsWhitelistedDestinations(destinations []string) {
	p.SetPath("sms.whitelisted_destinations", destinations)
}

func (p *createVerifyProfileParams) SetSmsDefaultVerificationTimeoutSecs(timeoutSecs int) {
	p.SetPath("sms.default_verification_timeout_secs", timeoutSecs)
}

func (c *Client) NewCreateVerifyProfileRequest(params *createVerifyProfileParams) (*http.Request, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	req, err := c.NewRequest("POST", "/verify_profiles", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request %w", err)
	}
	return req, nil
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params *createVerifyProfileParams) (json.RawMessage, error) {
	req, err := c.NewCreateVerifyProfileRequest(params)
	if err != nil {
		return nil, err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
