package telnyx

import (
	"context"
	"os"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

type TriggerSmsParams struct {
	telnyx.TriggerSmsPayload
}

func (p *TriggerSmsParams) GetParamPointers() []any {
	return []any{&p.VerifyProfileId, &p.PhoneNumber}
}

func TriggerSmsVerification(ctx context.Context, params *jsonrpc.Params) (any, error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p TriggerSmsParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	result, err := client.TriggerSmsVerification(p.TriggerSmsPayload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
