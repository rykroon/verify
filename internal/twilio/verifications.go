package twilio

import (
	"net/http"
	"net/url"
	"strings"
)

type SendVerificationParams struct {
	serviceSid string
	form       url.Values
}

func (svp *SendVerificationParams) ServiceSid() string {
	return svp.serviceSid
}

func (svp *SendVerificationParams) SetServiceSid(s string) {
	svp.serviceSid = s
}

func (svp *SendVerificationParams) To() string {
	return svp.form.Get("To")
}

func (svp *SendVerificationParams) SetTo(s string) {
	svp.form.Set("To", s)
}

func (svp *SendVerificationParams) Channel() string {
	return svp.form.Get("Channel")
}

func (svp *SendVerificationParams) SetChannel(s string) {
	svp.form.Set("Channel", s)
}

func (c *Client) SendVerification(params SendVerificationParams) (map[string]any, error) {
	path := "Services/" + params.serviceSid + "/Verifications"
	req, err := c.newRequest("POST", path, strings.NewReader(params.form.Encode()))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	err = c.processResponse(resp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
