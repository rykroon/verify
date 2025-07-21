package telnyx

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VerificationProfileResponse struct {
	Data struct {
		Id                 string `json:"id"`
		Name               string `json:"name"`
		WebhookUrl         string `json:"webhook_url"`
		WebhookFailoverUrl string `json:"webhook_failover_url"`
		RecordType         string `json:"record_type"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
		Sms                struct {
			MessagingTemplateId            string   `json:"messaging_template_id"`
			AppName                        string   `json:"app_name"`
			AlphaSender                    string   `json:"alpha_sender"`
			CodeLength                     string   `json:"code_length"`
			WhitelistedDestinations        []string `json:"white_listed_destinations"`
			DefaultTimeoutVerificationSecs int      `json:"default_timeout_verification_secs"`
		} `json:"sms"`
		Call struct {
			MessagingTemplateId            string   `json:"messaging_template_id"`
			AppName                        string   `json:"app_name"`
			CodeLength                     string   `json:"code_length"`
			WhitelistedDestinations        []string `json:"white_listed_destinations"`
			DefaultTimeoutVerificationSecs int      `json:"default_timeout_verification_secs"`
		} `json:"call"`
		FlashCall struct {
			DefaultTimeoutVerificationSecs int `json:"default_timeout_verification_secs"`
		} `json:"flash_call"`
		Language string `json:"string"`
	} `json:"data"`
}

type CreateVerifyProfileParams struct {
	m map[string]any
}

func NewCreateVerifyProfileParams(name string) *CreateVerifyProfileParams {
	m := map[string]any{"name": name}
	p := &CreateVerifyProfileParams{m: m}
	return p
}

func (p *CreateVerifyProfileParams) setSms(k string, v any) {
	val, exists := p.m["sms"]
	if !exists {
		p.m["sms"] = map[string]any{k: v}
		return
	}
	sms, isMap := val.(map[string]any)
	if !isMap {
		panic("should never happen")
	}
	sms[k] = v

}

func (p *CreateVerifyProfileParams) SetSmsMessagingTemplateId(templateId string) {
	p.setSms("messaging_template_id", templateId)
}

func (p *CreateVerifyProfileParams) SetSmsAppName(appName string) {
	p.setSms("app_name", appName)
}

func (p *CreateVerifyProfileParams) SetSmsCodeLength(codeLength string) {
	p.setSms("code_length", codeLength)
}

func (p *CreateVerifyProfileParams) SetSmsWhitelistedDestinations(destinations []string) {
	p.setSms("whitelisted_destinations", destinations)
}

func (p *CreateVerifyProfileParams) SetSmsDefaultVerificationTimeoutSecs(v int) {
	p.setSms("default_verification_timeout_secs", v)
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params *CreateVerifyProfileParams) (*VerificationProfileResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verify_profiles", &buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result *VerificationProfileResponse
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (*VerificationProfileResponse, error) {
	path := "verify_profiles/" + verifyProfileId
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var result *VerificationProfileResponse
	err = c.processResponse(resp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
