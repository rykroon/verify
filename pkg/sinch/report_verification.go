package sinch

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/rykroon/verify/internal/utils"
)

type reportVerificationParams struct {
	*utils.ParamBuilder
}

func NewReportVerificationParams() *reportVerificationParams {
	return &reportVerificationParams{utils.NewParamBuilder()}
}

func (p *reportVerificationParams) SetCode(code string) {
	p.Set("code", code)
}

func (p *reportVerificationParams) SetMethod(method string) {
	p.Set("method", method)
}

func (c *client) ReportVerificationById(id string, params *reportVerificationParams) (json.RawMessage, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}

	req, err := c.newRequest("PUT", fmt.Sprintf("/verifications/id/%s", id), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := c.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}
