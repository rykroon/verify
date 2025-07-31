package twilio

import "fmt"

func (c *Client) ListServices() (map[string]any, error) {
	req, err := c.newRequest("GET", "/Services", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response as json: %w", err)
	}
	return result, nil
}

func (c *Client) FetchService(sid string) (map[string]any, error) {
	req, err := c.newRequest("GET", "/Services/"+sid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response as json: %w", err)
	}
	return result, nil
}
