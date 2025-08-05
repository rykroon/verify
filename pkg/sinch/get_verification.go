package sinch

import "encoding/json"

func (c *client) GetVerificationById(id string) (json.RawMessage, error) {
	req, err := c.newRequest("GET", "verifications/id/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var raw json.RawMessage
	if err := json.Unmarshal(resp.Body, &raw); err != nil {
		return nil, err
	}

	return raw, nil
}
