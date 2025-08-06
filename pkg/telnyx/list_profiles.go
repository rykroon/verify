package telnyx

import (
	"fmt"
	"net/http"
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

func (c *Client) NewListVerifyProfilesRequest(params *listVerifyProfilesParams) (*http.Request, error) {
	req, err := c.NewRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}
	return req, nil
}
