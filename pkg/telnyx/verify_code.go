package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type VerifyCodeParams struct {
	Code string `json:"code"`
}

func (c *Client) VerifyCode(verificationId string, params VerifyCodeParams) (*utils.CachedResponse, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize json %w", err)
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.NewRequest("POST", path, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
