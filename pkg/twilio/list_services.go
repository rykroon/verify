package twilio

import (
	"fmt"
	"net/http"
)

func (c *Client) NewListServicesRequest() (*http.Request, error) {
	req, err := c.NewRequest("GET", "/Services", nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) NewFetchServiceRequest(sid string) (*http.Request, error) {
	req, err := c.NewRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil

}
