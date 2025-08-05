package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rykroon/verify/internal/utils"
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

func (c *Client) CheckVerification(serviceSid string, params *checkVerificationParams) (json.RawMessage, error) {
	req, err := c.NewCheckVerificationRequest(serviceSid, params)
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
