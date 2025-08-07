package sinch

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type Identity struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

type StartVerificationParams struct {
	Identity Identity `json:"identity"`
	Method   string   `json:"method"`
}

// https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-start/
func (c *Client) StartVerification(params StartVerificationParams) (*utils.CachedResponse, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params to json: %w", err)
	}
	req, err := c.NewRequest("POST", "/verifications", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ReportVerificationParams struct {
	Code   string `json:"code"`
	Method string `json:"method"`
}

// https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-report/#tag/Verifications-report/operation/ReportVerificationById
func (c *Client) ReportVerificationById(verificationId string, params ReportVerificationParams) (*utils.CachedResponse, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}

	req, err := c.NewRequest("PUT", fmt.Sprintf("/verifications/id/%s", verificationId), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetVerificationById(verificationId string) (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "verifications/id/"+verificationId, nil)
	if err != nil {
		return nil, err
	}

	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
