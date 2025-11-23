package twilio

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type SendVerificationParams struct {
	ServiceSid string `json:"service_sid"`
	SendVerificationForm
}

type SendVerificationForm struct {
	To      string `url:"To" json:"to"`
	Channel string `url:"Channel" json:"channel"`
}

// https://www.twilio.com/docs/verify/api/verification#start-new-verification
func (c *Client) SendVerification(serviceSid string, form SendVerificationForm) (map[string]any, error) {
	path := "Services/" + serviceSid + "/Verifications"
	values, err := query.Values(form)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
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

type CheckVerificationParams struct {
	ServiceSid string `json:"service_sid"`
	CheckVerificationForm
}

type CheckVerificationForm struct {
	To              string `url:"To" json:"to"`
	VerificationSid string `url:"VerificationSid" json:"verification_sid"`
	Code            string `url:"Code" json:"code"`
}

func (c *Client) CheckVerification(serviceSid string, form CheckVerificationForm) (map[string]any, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	values, err := query.Values(form)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
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
