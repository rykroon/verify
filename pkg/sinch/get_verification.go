package sinch

import (
	"encoding/json"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

func (c *client) GetVerificationById(id string) (json.RawMessage, error) {
	req, err := c.NewRequest("GET", "verifications/id/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	var raw json.RawMessage
	if err := json.Unmarshal(resp.Body, &raw); err != nil {
		return nil, err
	}

	return raw, nil
}
