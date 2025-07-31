package twilio

import (
	"fmt"

	"github.com/rykroon/verify/internal/httpx"
)

type createServiceParams struct {
	httpx.FormBody
}

func NewCreateServiceParams() *createServiceParams {
	return &createServiceParams{httpx.NewFormBody()}
}

func (csp *createServiceParams) SetFriendlyName(s string) {
	csp.Set("FriendlyName", s)
}

func (c *Client) CreateService(params *createServiceParams) (map[string]any, error) {
	req, err := c.newRequest("POST", "Services", params)
	if err != nil {
		return nil, err
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}
