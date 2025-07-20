package twilio

import (
	"net/http"
	"net/url"
	"strings"
)

type CreateServiceParams struct {
	form url.Values
}

func (csp *CreateServiceParams) FriendlyName() string {
	return csp.form.Get("FriendlyName")
}

func (csp *CreateServiceParams) SetFriendlyName(s string) {
	csp.form.Set("FriendlyName", s)
}

func (c *Client) CreateService(params CreateServiceParams) (map[string]any, error) {
	req, err := c.newRequest("POST", "Services", strings.NewReader(params.form.Encode()))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = c.processResponse(resp, result)
	return result, nil
}
