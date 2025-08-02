package sinch

import (
	"fmt"

	ds "github.com/rykroon/verify/internal/data_structures"
	"github.com/rykroon/verify/internal/httpx"
)

type reportVerificationParams struct {
	ds.WriteOnlyMap
}

func NewReportVerificationParams() *reportVerificationParams {
	return &reportVerificationParams{ds.NewWriteOnlyMap()}
}

func (p *reportVerificationParams) SetMethod(method string) {
	p.SetString("method", method)
}

func (c *client) ReportVerificationById(id string, params *reportVerificationParams) (map[string]any, error) {
	reqBody, err := httpx.NewJsonBody(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params as json: %w", err)
	}

	req, err := c.newRequest("PUT", fmt.Sprintf("/verifications/id/%s", id), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	respBody, err := c.sendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var result map[string]any
	if err := respBody.UnmarshalJson(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body as json: %w", err)
	}

	return result, nil
}
