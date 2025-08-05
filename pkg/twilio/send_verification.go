package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/rykroon/verify/internal/utils"
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

func (c *Client) SendVerification(serviceSid string, params *sendVerificationParams) (json.RawMessage, error) {
	req, err := c.NewSendVerificationRequest(serviceSid, params)
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}
	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return rawJson, nil
}
