package telnyx

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/rykroon/verify/internal/utils"
)

type ListVerifyProfilesParams struct {
	PageSize   int `json:"page_size,omitempty" url:"page[size],omitempty"`
	PageNumber int `json:"page_number,omitempty" url:"page[number],omitempty"`
}

func (c *Client) ListVerifyProfiles(params *ListVerifyProfilesParams) (*utils.CachedResponse, error) {
	req, err := c.NewRequest("GET", "verify_profiles", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with params: %w", err)
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	resp, err := utils.DoAndReadAll(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
