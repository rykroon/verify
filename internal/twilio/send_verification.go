package twilio

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type sendVerificationParams struct {
	serviceSid string
	httpx.FormBody
}

func NewSendVerificationParams(serviceSid, to, channel string) *sendVerificationParams {
	p := &sendVerificationParams{serviceSid, httpx.FormBody{}}
	p.Set("to", to)
	p.Set("channel", channel)
	return p
}

func (c *Client) SendVerification(params *sendVerificationParams) (map[string]any, error) {
	path := "Services/" + params.serviceSid + "/Verifications"
	req, err := c.newRequest("POST", path, params)
	if err != nil {
		return nil, err
	}
	respBody, err := c.do(req)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}
