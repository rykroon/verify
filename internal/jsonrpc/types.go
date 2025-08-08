package jsonrpc

import "encoding/json"

type JsonRpcRequest struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	Id      json.RawMessage `json:"id,omitempty"`
}

func (r JsonRpcRequest) IsNotification() bool {
	return len(r.Id) == 0
}

type JsonRpcResponse struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      json.RawMessage `json:"id"`
}

type JsonRpcSuccessResponse struct {
	JsonRpcResponse
	Result any `json:"result"`
}

type JsonRpcErrorResponse struct {
	JsonRpcResponse
	Error JsonRpcError `json:"error"`
}

func NewJsonRpcSuccessResponse(id json.RawMessage, result any) JsonRpcSuccessResponse {
	return JsonRpcSuccessResponse{JsonRpcResponse{"2.0", id}, result}
}

func NewJsonRpcErrorResponse(id json.RawMessage, err JsonRpcError) JsonRpcErrorResponse {
	if len(id) == 0 {
		id = json.RawMessage([]byte("null"))
	}
	return JsonRpcErrorResponse{JsonRpcResponse{"2.0", id}, err}
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewJsonRpcError(code int, msg string, data any) *JsonRpcError {
	return &JsonRpcError{code, msg, data}
}
