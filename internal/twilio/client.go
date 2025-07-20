package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
)

type Client struct {
	apiKeySid    string
	apiKeySecret string
}

func NewClient(apiKeySid, apiKeySecret string) *Client {
	return &Client{apiKeySid: apiKeySid, apiKeySecret: apiKeySecret}
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url, err := neturl.JoinPath("https://verify.twilio.com/v2/", path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiKeySid, c.apiKeySecret)
	return req, nil
}

func (c *Client) processResponse(resp *http.Response, v any) error {
	if resp.StatusCode >= 500 {
		return fmt.Errorf("internal server error")
	} else if resp.StatusCode >= 400 {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Telnyx Error: %s", string(content))
	} // check 300 ?

	err := json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
