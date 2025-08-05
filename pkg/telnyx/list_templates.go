package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) NewListMessageTemplatesRequest() (*http.Request, error) {
	req, err := c.NewRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (c *Client) ListMessageTemplates() (json.RawMessage, error) {
	req, err := c.NewListMessageTemplatesRequest()
	if err != nil {
		return nil, err
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
