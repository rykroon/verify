package server

import (
	"github.com/rykroon/verify/internal/jsonrpc"
	"github.com/rykroon/verify/internal/server/telnyx"
)

func GetJsonRpcServer() *jsonrpc.JsonRpcServer {
	server := jsonrpc.NewJsonRpcServer()
	server.Register("telnyx.list_profiles", telnyx.ListProfiles)
	server.Register("telnyx.trigger_sms", telnyx.TriggerSmsVerification)
	return server
}
