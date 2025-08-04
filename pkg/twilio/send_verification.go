package twilio

import (
	"encoding/json"
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

func (c *Client) SendVerification(serviceSid string, params *sendVerificationParams) (json.RawMessage, error) {
	path := "Services/" + serviceSid + "/Verifications"
	req, err := c.newRequest("POST", path, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return rawJson, nil
}
