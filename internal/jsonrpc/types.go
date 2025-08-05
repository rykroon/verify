package jsonrpc

import "encoding/json"

type JsonRpcRequest struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	Id      json.RawMessage `json:"id,omitempty"`
}

type JsonRpcResponse struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  any             `json:"result,omitempty"`
	Error   *JsonRpcError   `json:"error,omitempty"`
	Id      json.RawMessage `json:"id"`
}

func NewJsonRpcSuccessResponse(id json.RawMessage, result any) *JsonRpcResponse {
	return &JsonRpcResponse{"2.0", result, nil, id}
}

func NewJsonRpcErrorResponse(id json.RawMessage, err *JsonRpcError) *JsonRpcResponse {
	return &JsonRpcResponse{"2.0", nil, err, id}
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
