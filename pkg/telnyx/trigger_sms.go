package telnyx

import (
	"bytes"
	"encoding/json"

	"github.com/rykroon/verify/internal/utils"
)

type TriggerSmsParams struct {
	PhoneNumber     string `json:"phone_number"`
	VerifyProfileId string `json:"verify_profile_id"`
}

func (c *Client) TriggerSmsVerification(params TriggerSmsParams) (*utils.CachedResponse, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", "/verifications/sms", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
