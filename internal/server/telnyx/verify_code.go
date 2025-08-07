package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

type VerifyCodeParams struct {
	VerificationId string `json:"verification_id"`
	telnyx.VerifyCodeParams
}

func VerifyCode(ctx context.Context, rawParams json.RawMessage) (any, *jsonrpc.JsonRpcError) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var params VerifyCodeParams
	if err := json.Unmarshal(rawParams, &params); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, "invalid params", nil)
	}

	resp, err := client.VerifyCode(params.VerificationId, params.VerifyCodeParams)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	return result, nil
}
