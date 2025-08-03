package telnyx

import (
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

func (c *Client) ListVerifyProfiles(params *listVerifyProfilesParams) (*ListResponse[VerificationProfile], error) {
	req, err := c.newRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	respBody, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if !respBody.IsJson() {
		return nil, fmt.Errorf("expected json response")
	}

	var result *ListResponse[VerificationProfile]
	err = respBody.UnmarshalJson(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
