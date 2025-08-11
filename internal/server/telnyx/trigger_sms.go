package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func TriggerSmsVerification(ctx context.Context, params jsonrpc.Params) (any, *jsonrpc.Error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p telnyx.TriggerSmsParams
	if err := params.UnmarshalTo(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	resp, err := client.TriggerSmsVerification(p)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	return result, nil
}
