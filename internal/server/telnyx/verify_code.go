package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func VerifyCode(ctx context.Context, params jsonrpc.Params) (any, error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p telnyx.VerifyCodeParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, "invalid params", nil)
	}

	resp, err := client.VerifyCode(p.VerificationId, p.VerifyCodePayload)
	if err != nil {
		return nil, err
	}

	//if resp.StatusCode >= 400 ...

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
