package twilio

import (
	"encoding/json"
	"net/url"
	"strings"
)

type createServiceParams struct {
	url.Values
}

func NewCreateServiceParams() *createServiceParams {
	return &createServiceParams{url.Values{}}
}

func (csp *createServiceParams) SetFriendlyName(s string) {
	csp.Set("FriendlyName", s)
}

func (c *Client) CreateService(params *createServiceParams) (json.RawMessage, error) {
	req, err := c.newRequest("POST", "Services", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
