package server

import (
	"context"
	"fmt"
	"time"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/internal/server/telnyx"
	"github.com/rykroon/verify/internal/server/twilio"
)

func GetJsonRpcServer() jsonrpc.JsonRpcServer {
	server := jsonrpc.NewServer()
	server.Register("telnyx.list_profiles", telnyx.ListProfiles)
	server.Register("telnyx.trigger_sms", telnyx.TriggerSmsVerification)
	server.Register("telnyx.verify_code", telnyx.VerifyCode)
	server.Register("twilio.send_verification", twilio.SendVerification)
	server.Register("twilio.check_verification", twilio.CheckVerification)
	server.Register("echo", Echo)
	server.Register("sleep", Sleep)
	return server
}

func Echo(ctx context.Context, params *jsonrpc.Params) (any, error) {
	type Params struct {
		Text string `json:"text"`
	}
	var p Params
	if err := params.DecodeInto(&p); err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorCodeInvalidParams, err.Error(), nil)
	}

	return p.Text, nil
}

func Sleep(ctx context.Context, params *jsonrpc.Params) (any, error) {
	fmt.Println("Before Sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("After Sleep")
	return true, nil
}
