package twilio

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type SendVerificationParams struct {
	serviceSid string
	*httpx.FormBody
}

func (svp *SendVerificationParams) ServiceSid() string {
	return svp.serviceSid
}

func (svp *SendVerificationParams) SetServiceSid(s string) {
	svp.serviceSid = s
}

func (svp *SendVerificationParams) SetTo(s string) {
	svp.Set("To", s)
}

func (svp *SendVerificationParams) SetChannel(s string) {
	svp.Set("Channel", s)
}

func (c *Client) SendVerification(params SendVerificationParams) (map[string]any, error) {
	path := "Services/" + params.serviceSid + "/Verifications"
	req, err := c.newRequest("POST", path, params)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var result map[string]any
	if err := respBody.ToJson(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}
