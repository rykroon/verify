package telnyx

import (
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) ListMessageTemplates() (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
