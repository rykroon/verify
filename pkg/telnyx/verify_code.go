package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) NewVerifyCodeRequest(verificationId, code string) (*http.Request, error) {
	m := map[string]any{"code": code}
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.NewRequest("POST", path, bytes.NewReader(jsonData))
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
	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
