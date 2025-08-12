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
func (c *Client) CreateVerifyProfile(params CreateVerifyProfileParams) (*utils.CachedResponse, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	req, err := c.NewRequest("POST", "/verify_profiles", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListVerifyProfilesParams struct {
	PageSize   int `json:"page_size" url:"page[size],omitzero"`
	PageNumber int `json:"page_number" url:"page[number],omitzero"`
}

func (p *ListVerifyProfilesParams) GetParamPointers() []any {
	return []any{&p.PageSize, &p.PageNumber}
}

func (c *Client) ListVerifyProfiles(params *ListVerifyProfilesParams) (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
