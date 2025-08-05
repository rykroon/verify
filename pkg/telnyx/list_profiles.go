package telnyx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rykroon/verify/internal/utils"
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

func (c *Client) ListVerifyProfiles(params *listVerifyProfilesParams) (json.RawMessage, error) {
	req, err := c.NewListVerifyProfilesRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode json body as json: %w", err)
	}

	return result, nil
}
