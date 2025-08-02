package sinch

import (
	"fmt"

	ds "github.com/rykroon/verify/internal/data_structures"
	"github.com/rykroon/verify/internal/httpx"
)

type startVerificationParams struct {
	ds.WriteOnlyMap
}

func NewStartVerificationParams() *startVerificationParams {
	return &startVerificationParams{ds.NewWriteOnlyMap()}
}

func (p startVerificationParams) SetIdentityType(identityType string) {
	p.SetStringToPath("identity.type", identityType)
}

func (p startVerificationParams) SetIdentityEndpoint(identityEndpoint string) {
	p.SetStringToPath("identity.endpoint", identityEndpoint)
}

func (p startVerificationParams) SetMethod(method string) {
	p.SetString("method", method)
}

func (c *client) StartVerification(params *startVerificationParams) (map[string]any, error) {
	body, err := httpx.NewJsonBody(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create new json body: %w", err)
	}
	req, err := c.newRequest("POST", "/verifications", body)
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
