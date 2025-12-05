package telnyx

import (
	"context"
	"os"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

type VerifyCodeParams struct {
	VerificationId string `json:"verification_id"`
	telnyx.VerifyCodePayload
}

func (p *VerifyCodeParams) GetParamPointers() []any {
	return []any{&p.VerificationId, &p.VerifyCodePayload.Code}
}

func VerifyCode(ctx context.Context, params *jsonrpc.Params) (any, error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p VerifyCodeParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	result, err := client.VerifyCode(p.VerificationId, p.VerifyCodePayload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
