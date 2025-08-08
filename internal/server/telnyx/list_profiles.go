package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func ListProfiles(ctx context.Context, rawParams json.RawMessage) (any, *jsonrpc.JsonRpcError) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var params *telnyx.ListVerifyProfilesParams
	if rawParams != nil {
		if err := json.Unmarshal(rawParams, &params); err != nil {
			return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
		}
	}

	resp, err := client.ListVerifyProfiles(params)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil) // internal server error
	}

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	return result, nil
}
