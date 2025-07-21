package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Verification struct {
	Data struct {
		Id              string  `json:"id"`
		Type            string  `json:"type"`
		RecordType      string  `json:"record_type"`
		PhoneNumber     string  `json:"phone_number"`
		VerifyProfileId string  `json:"verify_profile_id"`
		CustomCode      *string `json:"custom_code"`
		TimeoutSecs     int     `json:"timeout_secs"`
		Status          string  `json:"status"`
		CreatedAt       string  `json:"created_at"`
		UpdatedAt       string  `json:"updated_at"`
	} `json:"data"`
}

type TriggerSmsVerificationParams struct {
	m map[string]any
}

func (p *TriggerSmsVerificationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.m)
}

func (p *TriggerSmsVerificationParams) SetCustomCode(customCode string) {
	p.m["custom_code"] = customCode
}

func (p *TriggerSmsVerificationParams) SetTimeoutSecs(timeoutSecs string) {
	p.m["timeout_secs"] = timeoutSecs
}

func NewTriggerSmsVerificationParams(phoneNumber, verifyProfileId string) *TriggerSmsVerificationParams {
	m := map[string]any{"phone_number": phoneNumber, "verify_profile_id": verifyProfileId}
	p := &TriggerSmsVerificationParams{m: m}
	return p
}

func (c *Client) TriggerSmsVerification(params *TriggerSmsVerificationParams) (*Verification, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verifications/sms", &buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result *Verification
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) VerifyCode(verificationId, code string) (map[string]any, error) {
	params := map[string]any{"code": code}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.newRequest("POST", path, &buf)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
