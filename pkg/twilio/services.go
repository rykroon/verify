package twilio

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type CreateServiceParams struct {
	FriendlyName string `url:"FriendlyName" json:"friendly_name"`
	CodeLength   int    `url:"CodeLength" json:"code_length,omitzero"`
}

func (c *Client) CreateService(params CreateServiceParams) (*utils.Content, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", "Services", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return content, err
}

func (c *Client) ListServices() (*utils.Content, error) {
	req, err := c.NewRequest("GET", "/Services", nil)
	if err != nil {
		return nil, err
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *Client) FetchService(sid string) (*utils.Content, error) {
	req, err := c.NewRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return content, nil
}
