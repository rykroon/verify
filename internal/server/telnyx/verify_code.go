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

func VerifyCode(ctx context.Context, params json.RawMessage) (any, *jsonrpc.JsonRpcError) {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))

	verificationId := ""
	code := ""
	req, err := client.NewVerifyCodeRequest(verificationId, code)

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
