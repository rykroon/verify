package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	ds "github.com/rykroon/verify/internal/data_structures"
)

type triggerSmsVerificationParams struct {
	*ds.ParamBuilder
}

func NewTriggerSmsVerificationParams() *triggerSmsVerificationParams {
	return &triggerSmsVerificationParams{ds.NewParamBuilder()}
}

func (p *triggerSmsVerificationParams) SetPhoneNumber(phoneNumber string) {
	p.Set("phone_number", phoneNumber)
}

func (p *triggerSmsVerificationParams) SetVerifyProfileId(verifyProfileId string) {
	p.Set("verify_profile_id", verifyProfileId)
}

func (c *Client) NewTriggerSmsVerificationRequest(params *triggerSmsVerificationParams) (*http.Request, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("POST", "/verifications/sms", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (c *Client) TriggerSmsVerification(params *triggerSmsVerificationParams) (json.RawMessage, error) {
	req, err := c.NewTriggerSmsVerificationRequest(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return rawJson, nil
}
