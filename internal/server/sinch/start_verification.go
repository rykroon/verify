package sinch

import (
	"context"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/sinch"
)

func StartVerification(ctx context.Context, params jsonrpc.Params) (any, error) {
	client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	var p sinch.StartVerificationPayload
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	result, err := client.StartVerification(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}
