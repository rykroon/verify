package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateVerifyProfileParams struct {
	m map[string]any
}

func NewCreateVerifyProfileParams(name string) *CreateVerifyProfileParams {
	p := &CreateVerifyProfileParams{make(map[string]any)}
	p.m["name"] = name
	return p
}

func (p *CreateVerifyProfileParams) GetSetMap(key string) map[string]any {
	var result map[string]any
	maybeMap, exists := p.m[key]
	if exists {
		result, isMap := maybeMap.(map[string]any)
		if !isMap {
			// if not a map, overwrite
			result = make(map[string]any)
			p.m[key] = result
		}
	} else {
		// create a new map
		result := make(map[string]any)
		p.m[key] = result
	}

	return result
}

func (p *CreateVerifyProfileParams) SetSmsMessagingTemplateId(templateId string) *CreateVerifyProfileParams {
	sms := p.GetSetMap("sms")
	sms["messaging_template_id"] = templateId
	return p
}

func (p *CreateVerifyProfileParams) SetSmsAppName(appName string) *CreateVerifyProfileParams {
	sms := p.GetSetMap("sms")
	sms["app_name"] = appName
	return p
}

func (p *CreateVerifyProfileParams) SetSmsCodeLength(codeLength string) *CreateVerifyProfileParams {
	sms := p.GetSetMap("sms")
	sms["code_length"] = codeLength
	return p
}

func (p *CreateVerifyProfileParams) SetSmsWhitelistedDestinations(destinations []string) *CreateVerifyProfileParams {
	sms := p.GetSetMap("sms")
	sms["whitelisted_destinations"] = destinations
	return p
}

func (p *CreateVerifyProfileParams) SetSmsDefaultVerificationTimeoutSecs(v int) *CreateVerifyProfileParams {
	sms := p.GetSetMap("sms")
	sms["default_verification_timeout_secs"] = v
	return p
}

func (c *Client) NewCreateVerifyProfileRequest(params *CreateVerifyProfileParams) (*http.Request, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	req, err := c.newRequest("POST", "/verify_profiles", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request %w", err)
	}
	return req, nil
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params *CreateVerifyProfileParams) (json.RawMessage, error) {
	req, err := c.NewCreateVerifyProfileRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
