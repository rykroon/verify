package twilio

import (
	"context"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/twilio"
)

func SendVerification(ctx context.Context, params jsonrpc.Params) (any, *jsonrpc.Error) {
	var p twilio.SendVerificationParams
	if err := params.UnmarshalTo(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	_, err := client.SendVerification(p.ServiceSid, p.SendVerificationForm)
	if err != nil {
		return nil, nil // come back to this
	}

	// return result

	return nil, nil
}
