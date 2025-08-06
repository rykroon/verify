package twilio

import (
	"net/http"
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

func (c *Client) NewCreateServiceRequest(params *createServiceParams) (*http.Request, error) {
	req, err := c.NewRequest("POST", "Services", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	return req, err
}
