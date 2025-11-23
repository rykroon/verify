package twilio

import (
	"context"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/twilio"
)

func CheckVerification(ctx context.Context, params jsonrpc.Params) (any, error) {
	var p twilio.CheckVerificationParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	client := twilio.NewClient(
		os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"), nil,
	)

	result, err := client.CheckVerification(p.ServiceSid, p.CheckVerificationForm)
	if err != nil {
		return nil, err
	}

	return result, nil
}
