package telnyx

import (
	"bytes"
	"encoding/json"

	"github.com/rykroon/verify/internal/utils"
)

type TriggerSmsParams struct {
	PhoneNumber     string `json:"phone_number"`
	VerifyProfileId string `json:"verify_profile_id"`
	CustomCode      string `json:"custom_code,omitzero"`
	TimeoutSecs     string `json:"timeout_secs,omitzero"`
}

func (p *TriggerSmsParams) GetParamPointers() []any {
	return []any{&p.PhoneNumber, &p.VerifyProfileId}
}

// https://developers.telnyx.com/api/verify/create-verification-sms
func (c *Client) TriggerSmsVerification(params TriggerSmsParams) (*utils.Content, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", "/verifications/sms", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return content, nil
}
