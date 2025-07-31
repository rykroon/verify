package twilio

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type checkVerificationParams struct {
	httpx.FormBody
}

func (p *checkVerificationParams) SetTo(to string) {
	p.Set("To", to)
}

func (p *checkVerificationParams) SetVerificationSid(verificationSid string) {
	p.Set("VerificationSid", verificationSid)
}

func (p *checkVerificationParams) SetCode(code string) {
	p.Set("Code", code)
}

func NewCheckVerificationParams() *checkVerificationParams {
	return &checkVerificationParams{httpx.NewFormBody()}
}

func (c *Client) CheckVerification(serviceSid string, params *checkVerificationParams) (map[string]any, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	req, err := c.newRequest("POST", path, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal json body: %w", err)
	}

	return result, nil
}
