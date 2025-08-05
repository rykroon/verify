package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) NewVerifyCodeRequest(verificationId, code string) (*http.Request, error) {
	m := map[string]any{"code": code}
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.newRequest("POST", path, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (c *Client) VerifyCode(verificationId, code string) (json.RawMessage, error) {
	req, err := c.NewVerifyCodeRequest(verificationId, code)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
