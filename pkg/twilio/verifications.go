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
func (c *Client) SendVerification(serviceSid string, form SendVerificationForm) (*utils.CachedResponse, error) {
	path := "Services/" + serviceSid + "/Verifications"
	values, err := query.Values(form)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
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

func (c *Client) CheckVerification(serviceSid string, form CheckVerificationForm) (*utils.CachedResponse, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	values, err := query.Values(form)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
