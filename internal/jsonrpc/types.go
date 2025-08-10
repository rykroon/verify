package jsonrpc

import "encoding/json"

type JsonRpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  JsonRpcParams `json:"params,omitzero"`
	Id      JsonRpcId     `json:"id,omitzero"`
}

type JsonRpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      any           `json:"id"`
	Result  any           `json:"result"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

func (r *JsonRpcResponse) MarshalJSON() ([]byte, error) {
	resp := map[string]any{"jsonrpc": r.JsonRpc, "id": r.Id}
	if r.Error == nil {
		resp["result"] = r.Result
	} else {
		resp["error"] = r.Error
	}
	return json.Marshal(resp)
}

func NewJsonRpcSuccessResp(id any, result any) *JsonRpcResponse {
	return &JsonRpcResponse{"2.0", id, result, nil}
}

func NewJsonRpcErrorResp(id any, err *JsonRpcError) *JsonRpcResponse {
	return &JsonRpcResponse{"2.0", id, nil, err}
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewJsonRpcError(code int, msg string, data any) *JsonRpcError {
	return &JsonRpcError{code, msg, data}
}

func ParseError(data any) *JsonRpcError {
	return NewJsonRpcError(-32700, "Parse Error", data)
}

func InvalidRequest(data any) *JsonRpcError {
	return NewJsonRpcError(-32600, "Invalid Request", data)
}

func MethodNotFound(data any) *JsonRpcError {
	return NewJsonRpcError(-32601, "Method Not Found", data)
}

func InvalidParams(data any) *JsonRpcError {
	return NewJsonRpcError(-32602, "Invalid Params", data)
}
