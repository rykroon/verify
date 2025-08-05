package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/rykroon/verify/internal/utils"
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

func (c *Client) NewCreateServiceRequest(params *createServiceParams) (*http.Request, error) {
	req, err := c.NewRequest("POST", "Services", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	return req, err
}

func (c *Client) CreateService(params *createServiceParams) (json.RawMessage, error) {
	req, err := c.NewCreateServiceRequest(params)
	if err != nil {
		return nil, err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
