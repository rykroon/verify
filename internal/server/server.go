package server

import (
	"context"
	"fmt"
	"time"

	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/internal/server/telnyx"
)

func GetJsonRpcServer() *jsonrpc.JsonRpcServer {
	server := jsonrpc.NewJsonRpcServer()
	server.Register("telnyx.list_profiles", telnyx.ListProfiles)
	server.Register("telnyx.trigger_sms", telnyx.TriggerSmsVerification)
	server.Register("echo", Echo)
	server.Register("sleep", Sleep)
	return server
}

func Echo(ctx context.Context, params jsonrpc.Params) (any, *jsonrpc.Error) {
	type Params struct {
		Text string `json:"text"`
	}
	var p Params
	if err := params.Unmarshal(&p); err != nil {
		return nil, jsonrpc.NewJsonRpcError(-32602, "Invalid params", err.Error())
	}

	return p.Text, nil
}

func Sleep(ctx context.Context, params jsonrpc.Params) (any, *jsonrpc.Error) {
	fmt.Println("Before Sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("After Sleep")
	return true, nil
}
