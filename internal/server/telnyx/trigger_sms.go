package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func TriggerSmsVerification(ctx context.Context, rawParams json.RawMessage) (any, *jsonrpc.JsonRpcError) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var params telnyx.TriggerSmsParams
	if err := json.Unmarshal(rawParams, &params); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, "invalid params", nil)
	}

	resp, err := client.TriggerSmsVerification(params)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	return result, nil
}
