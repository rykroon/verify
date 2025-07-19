package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateVerifyProfileParams struct {
	Name               string `json:""`
	WebhookUrl         string `json:"webhook_url,omitempty"`
	WebhookFailoverUrl string `json:"webhook_failover_url,omitempty"`
	Sms                *struct {
		MessagingTemplateId            string   `json:"messaging_template_id.omitempty"`
		AppName                        string   `json:"app_name,omitempty"`
		AlphaSender                    string   `json:"alpha_sender,omitempty"`
		CodeLength                     int      `json:"code_length,omitempty"`
		WhiteListedDestinations        []string `json:"whitelisted_destinations"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
	} `json:"sms,omitempty"`
	Call *struct {
		MessagingTemplateId            string   `json:"messaging_template_id.omitempty"`
		AppName                        string   `json:"app_name,omitempty"`
		AlphaSender                    string   `json:"alpha_sender,omitempty"`
		CodeLength                     int      `json:"code_length,omitempty"`
		WhiteListedDestinations        []string `json:"whitelisted_destinations"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
	} `json:"call,omitempty"`
	FlashCall *struct {
		WhiteListedDestinations        []string `json:"whitelisted_destinations"`
		DefaultVerificationTimeoutSecs int      `json:"default_verification_timeout_secs,omitempty"`
	} `json:"flash_call,omitempty"`
}

func (c *Client) CreateVerifyProfile(params CreateVerifyProfileParams) (*ApiResponse[VerificationProfile], error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verify_profiles", &buf)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("Server Error")
	} else if resp.StatusCode >= 400 {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Telnyx error: %s", string(content))
	} // check 300 ?

	var result ApiResponse[VerificationProfile]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
