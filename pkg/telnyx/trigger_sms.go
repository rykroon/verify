package telnyx

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

type triggerSmsVerificationParams struct {
	*utils.ParamBuilder
}

func NewTriggerSmsVerificationParams() *triggerSmsVerificationParams {
	return &triggerSmsVerificationParams{utils.NewParamBuilder()}
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
	req, err := c.NewRequest("POST", "/verifications/sms", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	return req, nil
}
