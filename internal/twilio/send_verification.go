package twilio

import (
	"fmt"

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

func (c *Client) SendVerification(serviceSid string, params *sendVerificationParams) (map[string]any, error) {
	path := "Services/" + serviceSid + "/Verifications"
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
