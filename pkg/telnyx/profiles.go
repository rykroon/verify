package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type CreateVerifyProfileParams struct {
	Name string `json:"name"`
	Sms  *struct {
		MessagingTemplateId            string   `json:"messaging_template_id,omitempty"`
		AppName                        string   `json:"app_name,omitempty"`
		CodeLength                     int      `json:"code_length,omitempty"`
		WhitelistedDestinations        []string `json:"whitelisted_destinations,omitempty"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
	} `json:"sms,omitempty"`
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
	PageSize   int `json:"page_size,omitempty" url:"page[size],omitempty"`
	PageNumber int `json:"page_number,omitempty" url:"page[number],omitempty"`
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
