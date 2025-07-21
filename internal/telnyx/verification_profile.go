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

type CreateVerifyProfileSmsParams struct {
	MessagingTemplateId            string   `json:"messaging_template_id.omitempty"`
	AppName                        string   `json:"app_name,omitempty"`
	AlphaSender                    string   `json:"alpha_sender,omitempty"`
	CodeLength                     int      `json:"code_length,omitempty"`
	WhiteListedDestinations        []string `json:"whitelisted_destinations,omitempty"`
	DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
}

type CreateVerifyProfileCallParams struct {
	MessagingTemplateId            string   `json:"messaging_template_id.omitempty"`
	AppName                        string   `json:"app_name,omitempty"`
	AlphaSender                    string   `json:"alpha_sender,omitempty"`
	CodeLength                     int      `json:"code_length,omitempty"`
	WhiteListedDestinations        []string `json:"whitelisted_destinations,omitempty"`
	DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
}

type CreateVerifyProfileFlashCallParams struct {
	WhiteListedDestinations        []string `json:"whitelisted_destinations,omitempty"`
	DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
}

type CreateVerifyProfileParams struct {
	Name               string                              `json:"name"`
	WebhookUrl         string                              `json:"webhook_url,omitempty"`
	WebhookFailoverUrl string                              `json:"webhook_failover_url,omitempty"`
	Sms                *CreateVerifyProfileSmsParams       `json:"sms,omitempty"`
	Call               *CreateVerifyProfileCallParams      `json:"call,omitempty"`
	FlashCall          *CreateVerifyProfileFlashCallParams `json:"flash_call,omitempty"`
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
