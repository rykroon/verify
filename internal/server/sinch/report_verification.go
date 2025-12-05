package sinch

import (
	"context"
	"os"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/pkg/sinch"
)

type ReportVerificationParams struct {
	Id string `json:"id"`
	sinch.ReportVerificationPayload
}

func ReportVerification(ctx context.Context, params *jsonrpc.Params) (any, error) {
	client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	var p ReportVerificationParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	result, err := client.ReportVerificationById(p.Id, p.ReportVerificationPayload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
