package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type VerifyCodeParams struct {
	VerificationId string `json:"verification_id"`
	VerifyCodePayload
}

func (p *VerifyCodeParams) GetParamPointers() []any {
	return []any{&p.VerificationId, &p.VerifyCodePayload.Code}
}

type VerifyCodePayload struct {
	Code   string `json:"code"`
	Status string `json:"status,omitzero"`
}

// https://developers.telnyx.com/api/verify/verify-verification-code-by-id
func (c *Client) VerifyCode(verificationId string, params VerifyCodePayload) (*utils.CachedResponse, error) {
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
