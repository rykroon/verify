package telnyx

import (
	"context"
	"os"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/pkg/telnyx"
)

func ListProfiles(ctx context.Context, params *jsonrpc.Params) (any, error) {
	client := telnyx.NewClient(nil, os.Getenv("TELNYX_API_KEY"))

	var p telnyx.ListVerifyProfilesParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	result, err := client.ListVerifyProfiles(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}
