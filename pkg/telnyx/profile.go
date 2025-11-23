package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type UpdateVerifyProfileParams struct {
	VerifyProfileId string `json:"verify_profile_id"`
	UpdateVerifyProfilePayload
}

func (p *UpdateVerifyProfileParams) GetParamPointers() []any {
	return []any{&p.VerifyProfileId, &p.UpdateVerifyProfilePayload}
}

type UpdateVerifyProfilePayload struct {
	Name string `json:"name,omitzero"`
	Sms  struct {
		MessagingTemplateId            string   `json:"messaging_template_id,omitzero"`
		AppName                        string   `json:"app_name,omitzero"`
		CodeLength                     int      `json:"code_length,omitzero"`
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitzero"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitzero"`
	} `json:"sms,omitzero"`
}

/*
https://developers.telnyx.com/api/verify/update-verify-profile
*/
func (c *Client) UpdateVerifyProfile(verifyProfileId string, params UpdateVerifyProfilePayload) (map[string]any, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	req, err := c.NewRequest("PATCH", "/verify_profiles/"+verifyProfileId, bytes.NewReader(jsonData))
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

type RetrieveVerifyProfileParams struct {
	VerifyProfileId string `json:"verify_profile_id"`
}

func (p *RetrieveVerifyProfileParams) GetParamPointers() []any {
	return []any{&p.VerifyProfileId}
}

func (c *Client) RetrieveVerifyProfile(verifyProfileId string) (map[string]any, error) {
	req, err := c.NewRequest("GET", "verify_profiles/"+verifyProfileId, nil)
	if err != nil {
		return nil, err
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

func (c *Client) DeleteVerifyProfile(verifyProfileId string) (map[string]any, error) {
	req, err := c.NewRequest("DELETE", "verify_profiles/"+verifyProfileId, nil)
	if err != nil {
		return nil, err
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
