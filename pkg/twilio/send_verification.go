package twilio

import (
	"net/http"
	"net/url"
	"strings"
)

type sendVerificationParams struct {
	url.Values
}

func NewSendVerificationParams(to, channel string) *sendVerificationParams {
	p := &sendVerificationParams{url.Values{}}
	p.Set("To", to)
	p.Set("Channel", channel)
	return p
}

func (c *Client) NewSendVerificationRequest(serviceSid string, params *sendVerificationParams) (*http.Request, error) {
	path := "Services/" + serviceSid + "/Verifications"
	req, err := c.NewRequest("POST", path, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	return req, nil
}
