package telnyx

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func ListProfiles(ctx context.Context, params jsonrpc.Params) (any, error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p telnyx.ListVerifyProfilesParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	content, err := client.ListVerifyProfiles(p)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(content.Data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
