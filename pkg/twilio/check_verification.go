package twilio

import (
	"fmt"
	"net/http"
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

func (c *Client) NewCheckVerificationRequest(serviceSid string, params *checkVerificationParams) (*http.Request, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	req, err := c.NewRequest("POST", path, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	return req, nil
}
