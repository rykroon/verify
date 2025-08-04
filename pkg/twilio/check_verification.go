package twilio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type checkVerificationParams struct {
	url.Values
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
	return &checkVerificationParams{url.Values{}}
}

func (c *Client) CheckVerification(serviceSid string, params *checkVerificationParams) (json.RawMessage, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	req, err := c.newRequest("POST", path, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
