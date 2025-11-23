package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type CreateVerifyProfileParams struct {
	Name               string `json:"name"`
	WebhookUrl         string `json:"webhook_url,omitzero"`
	WebhookFailoverUrl string `json:"webhook_failover_url,omitzero"`
	Sms                struct {
		MessagingTemplateId            string   `json:"messaging_template_id,omitzero"`
		AppName                        string   `json:"app_name,omitzero"`
		AlphaSender                    string   `json:"alpha_sender,omitzero"`
		CodeLength                     int      `json:"code_length,omitzero"`
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitzero"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitzero"`
	} `json:"sms,omitzero"`
	Call struct {
		MessagingTemplateId            string   `json:"messaging_template_id,omitzero"`
		AppName                        string   `json:"app_name,omitzero"`
		CodeLength                     int      `json:"code_length,omitzero"`
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitzero"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitzero"`
	} `json:"call,omitzero"`
	FlashCall struct {
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitzero"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitzero"`
	} `json:"flashcall,omitzero"`
	Language string `json:"language,omitzero"`
}

func (p *CreateVerifyProfileParams) GetParamPointers() []any {
	return []any{&p.Name, &p.WebhookUrl, &p.WebhookFailoverUrl, &p.Sms, &p.Call, &p.FlashCall, &p.Language}
}

/*
https://developers.telnyx.com/api/verify/create-verify-profile
*/
func (c *Client) CreateVerifyProfile(params CreateVerifyProfileParams) (map[string]any, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	req, err := c.NewRequest("POST", "/verify_profiles", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request %w", err)
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	if !content.IsJson() {
		return nil, fmt.Errorf("expected json but got %s", content.Type)
	}
	var result map[string]any
	err = content.DecodeJsonInto(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}

type ListVerifyProfilesParams struct {
	PageSize   int `json:"page_size" url:"page[size],omitempty"`
	PageNumber int `json:"page_number" url:"page[number],omitempty"`
}

func (p *ListVerifyProfilesParams) GetParamPointers() []any {
	return []any{&p.PageSize, &p.PageNumber}
}

func (c *Client) ListVerifyProfiles(params ListVerifyProfilesParams) (map[string]any, error) {
	req, err := c.NewRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	if !content.IsJson() {
		return nil, fmt.Errorf("expected json but got %s", content.Type)
	}
	var result map[string]any
	err = content.DecodeJsonInto(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}
