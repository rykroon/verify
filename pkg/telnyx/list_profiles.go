package telnyx

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type listVerifyProfilesParams struct {
	url.Values
}

func NewListVerifyProfilesParams() *listVerifyProfilesParams {
	return &listVerifyProfilesParams{url.Values{}}
}

func (p *listVerifyProfilesParams) SetPageSize(pageSize int) {
	p.Set("page[size]", strconv.Itoa(pageSize))
}

func (p *listVerifyProfilesParams) SetPageNumber(pageNumber int) {
	p.Set("page[number]", strconv.Itoa(pageNumber))
}

func (c *Client) ListVerifyProfiles(params *listVerifyProfilesParams) (json.RawMessage, error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
