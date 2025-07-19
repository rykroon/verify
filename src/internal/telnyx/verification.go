package telnyx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TriggerSmsVerificationParams struct {
	PhoneNumber     string  `json:"phone_number"`
	VerifyProfileId string  `json:"verify_profile_id"`
	CustomCode      *string `json:"custom_code"`
	TimeoutSecs     int     `json:"timeout_secs"`
}

func (c *Client) TriggerSmsVerification(params TriggerSmsVerificationParams) (*ApiResponse[Verification], error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(params)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/verifications/sms", &buf)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("Server Error")
	} else if resp.StatusCode >= 400 {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Telnyx error: %s", string(content))
	} // check 300 ?

	var result ApiResponse[Verification]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("Server Error")
	} else if resp.StatusCode >= 400 {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Telnyx error: %s", string(content))
	} // check 300 ?

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
