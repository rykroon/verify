package twilio

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type CreateServiceForm struct {
	FriendlyName string `url:"FriendlyName" json:"friendly_name"`
	CodeLength   int    `url:"CodeLength" json:"code_length,omitzero"`
}

func (c *Client) CreateService(form CreateServiceForm) (*utils.Content, error) {
	values, err := query.Values(form)
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

func (c *Client) ListServices() (map[string]any, error) {
	req, err := c.NewRequest("GET", "/Services", nil)
	if err != nil {
		return nil, err
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	if !content.IsJson() {
		return nil, fmt.Errorf("expected json but got %s", content.Type)
	}
	var result map[string]any
	err = content.DecodeJsonInto(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}

func (c *Client) FetchService(sid string) (map[string]any, error) {
	req, err := c.NewRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	content, err := utils.SendRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	if !content.IsJson() {
		return nil, fmt.Errorf("expected json but got %s", content.Type)
	}
	var result map[string]any
	err = content.DecodeJsonInto(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}
