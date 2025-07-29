package twilio

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type CheckVerificationParams struct {
	serviceSid string
	httpx.FormBody
}

func (p *CheckVerificationParams) SetTo(to string) {
	p.Set("to", to)
}

func (p *CheckVerificationParams) SetVerificationSid(verificationSid string) {
	p.Set("verificationSid", verificationSid)
}

func (p *CheckVerificationParams) SetCode(code string) {
	p.Set("code", code)
}

func NewCheckVerificationParams(serviceSid string) *CheckVerificationParams {
	return &CheckVerificationParams{serviceSid, httpx.NewFormBody()}
}

func (c *Client) CheckVerification(params *CheckVerificationParams) (map[string]any, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", params.serviceSid)
	req, err := c.newRequest("POST", path, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	respBody, err := resp.ReadBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal json body: %w", err)
	}

	return result, nil
}
