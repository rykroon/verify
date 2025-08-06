package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/telnyx"
)

func ListProfiles(ctx context.Context, params json.RawMessage) (any, *jsonrpc.JsonRpcError) {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	payload := telnyx.NewListVerifyProfilesParams()
	req, err := client.NewListVerifyProfilesRequest(payload)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}
	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
	if err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, jsonrpc.NewJsonRpcError(0, err.Error(), nil)
	}

	return result, nil
}
