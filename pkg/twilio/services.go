package twilio

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type CreateServiceParams struct {
	FriendlyName string `url:"FriendlyName"`
}

func (c *Client) CreateService(params CreateServiceParams) (*utils.CachedResponse, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", "Services", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListServices() (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "/Services", nil)
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) FetchService(sid string) (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
