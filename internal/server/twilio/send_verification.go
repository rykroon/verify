package twilio

import (
	"context"
	"os"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/pkg/twilio"
)

type SendVerificationParams struct {
	ServiceSid string `json:"service_sid"`
	twilio.SendVerificationForm
}

func SendVerification(ctx context.Context, params *jsonrpc.Params) (any, error) {
	var p SendVerificationParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	client := twilio.NewClient(
		os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"), nil,
	)

	result, err := client.SendVerification(p.ServiceSid, p.SendVerificationForm)
	if err != nil {
		return nil, err
	}

	return result, nil
}
