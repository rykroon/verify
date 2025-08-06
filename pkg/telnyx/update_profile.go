package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type UpdateVerifyProfileParams struct {
	VerifyProfileId string `json:"verify_profile_id"`
	*UpdateVerifyProfileBodyParams
}

func (p *UpdateVerifyProfileParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.UpdateVerifyProfileBodyParams)
}

type UpdateVerifyProfileBodyParams struct {
	Name string `json:"name,omitempty"`
	Sms  *struct {
		MessagingTemplateId            string   `json:"messaging_template_id,omitempty"`
		AppName                        string   `json:"app_name,omitempty"`
		CodeLength                     int      `json:"code_length,omitempty"`
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitempty"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
	} `json:"sms,omitempty"`
}

func NewUpdateVerifyProfileParams() *UpdateVerifyProfileParams {
	return &UpdateVerifyProfileParams{"", &UpdateVerifyProfileBodyParams{}}
}

/*
https://developers.telnyx.com/api/verify/update-verify-profile
*/
func (c *Client) UpdateVerifyProfile(params *UpdateVerifyProfileParams) (*utils.CachedResponse, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}
	fmt.Println(string(jsonData))
	req, err := c.NewRequest("PATCH", "/verify_profiles/"+params.VerifyProfileId, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
