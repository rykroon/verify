package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type TriggerSmsPayload struct {
	PhoneNumber     string `json:"phone_number"`
	VerifyProfileId string `json:"verify_profile_id"`
	CustomCode      string `json:"custom_code,omitzero"`
	TimeoutSecs     string `json:"timeout_secs,omitzero"`
}

// https://developers.telnyx.com/api/verify/create-verification-sms
func (c *Client) TriggerSmsVerification(payload TriggerSmsPayload) (map[string]any, error) {
	jsonData, err := json.Marshal(payload)
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
