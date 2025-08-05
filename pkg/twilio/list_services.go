package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

func (c *Client) NewListServicesRequest() (*http.Request, error) {
	req, err := c.NewRequest("GET", "/Services", nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) ListServices() (json.RawMessage, error) {
	req, err := c.NewListServicesRequest()
	if err != nil {
		return nil, err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}

func (c *Client) NewFetchServiceRequest(sid string) (*http.Request, error) {
	req, err := c.NewRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil

}

func (c *Client) FetchService(sid string) (json.RawMessage, error) {
	req, err := c.NewFetchServiceRequest(sid)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
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
