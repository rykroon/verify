package sinch

import (
	"bytes"
	"encoding/json"
	"fmt"

	ds "github.com/rykroon/verify/internal/data_structures"
)

type startVerificationParams struct {
	*ds.ParamBuilder
}

func NewStartVerificationParams() *startVerificationParams {
	return &startVerificationParams{ds.NewParamBuilder()}
}

func (p startVerificationParams) SetIdentityType(identityType string) {
	p.SetPath("identity.type", identityType)
}

func (p startVerificationParams) SetIdentityEndpoint(identityEndpoint string) {
	p.SetPath("identity.endpoint", identityEndpoint)
}

func (p startVerificationParams) SetMethod(method string) {
	p.Set("method", method)
}

func (c *client) StartVerification(params *startVerificationParams) (json.RawMessage, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params to json: %w", err)
	}
	req, err := c.newRequest("POST", "/verifications", bytes.NewReader(data))
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
