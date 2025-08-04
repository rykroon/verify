package twilio

import (
	"encoding/json"
	"strings"

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

func (c *Client) CreateService(params *createServiceParams) (json.RawMessage, error) {
	req, err := c.newRequest("POST", "Services", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
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
