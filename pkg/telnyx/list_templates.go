package telnyx

import (
	"fmt"
	"net/http"
)

func (c *Client) NewListMessageTemplatesRequest() (*http.Request, error) {
	req, err := c.NewRequest("GET", "/verify_profiles/templates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}
