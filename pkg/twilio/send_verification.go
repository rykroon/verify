package twilio

import (
	"encoding/json"
	"strings"

	"github.com/rykroon/verify/internal/httpx"
)

type sendVerificationParams struct {
	httpx.FormBody
}

func NewSendVerificationParams(to, channel string) *sendVerificationParams {
	p := &sendVerificationParams{httpx.NewFormBody()}
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
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return rawJson, nil
}
