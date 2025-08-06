package sinch

import (
	"net/http"
)

func (c *client) NewGetVerificationByIdRequest(id string) (*http.Request, error) {
	req, err := c.NewRequest("GET", "verifications/id/"+id, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
