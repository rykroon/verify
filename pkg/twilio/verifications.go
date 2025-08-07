package twilio

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type SendVerificationParams struct {
	To      string `url:"To"`
	Channel string `url:"Channel"`
}

func (c *Client) SendVerification(serviceSid string, params *SendVerificationParams) (*utils.CachedResponse, error) {
	path := "Services/" + serviceSid + "/Verifications"
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CheckVerificationParams struct {
	To              string `url:"To"`
	VerificationSid string `url:"VerificationSid"`
	Code            string `url:"Code"`
}

func (c *Client) CheckVerification(serviceSid string, params *CheckVerificationParams) (*utils.CachedResponse, error) {
	path := fmt.Sprintf("Services/%s/VerificationCheck", serviceSid)
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
