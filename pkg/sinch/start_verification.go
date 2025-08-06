package sinch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/utils"
)

type startVerificationParams struct {
	*utils.ParamBuilder
}

func NewStartVerificationParams() *startVerificationParams {
	return &startVerificationParams{utils.NewParamBuilder()}
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

func (c *client) NewStartVerificationRequest(params *startVerificationParams) (*http.Request, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode params to json: %w", err)
	}
	req, err := c.NewRequest("POST", "/verifications", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	return req, nil
}
