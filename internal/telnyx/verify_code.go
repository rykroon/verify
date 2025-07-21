package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) VerifyCode(verificationId, code string) (map[string]any, error) {
	params := map[string]any{"code": code}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationId)
	req, err := c.newRequest("POST", path, &buf)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = c.processResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
