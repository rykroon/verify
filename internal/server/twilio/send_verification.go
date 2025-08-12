package twilio

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/pkg/twilio"
)

func SendVerification(ctx context.Context, params jsonrpc.Params) (any, error) {
	var p twilio.SendVerificationParams
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.InvalidParams(err.Error())
	}

	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.SendVerification(p.ServiceSid, p.SendVerificationForm)
	if err != nil {
		return nil, err
	}

	//if resp.StatusCode >= 400 ...

	// return result
	var result map[string]any
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
