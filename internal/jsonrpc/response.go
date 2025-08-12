package jsonrpc

import (
	"encoding/json"
)

type Response struct {
	JsonRpc string `json:"jsonrpc"`
	Id      Id     `json:"id"`
	Result  any    `json:"result"`
	Error   *Error `json:"error,omitempty"`
}

func (r *Response) MarshalJSON() ([]byte, error) {
	resp := map[string]any{"jsonrpc": r.JsonRpc, "id": r.Id}
	if r.Error == nil {
		resp["result"] = r.Result
	} else {
		resp["error"] = r.Error
	}
	return json.Marshal(resp)
}

func NewJsonRpcSuccessResp(id Id, result any) *Response {
	return &Response{"2.0", id, result, nil}
}

func NewJsonRpcErrorResp(id Id, err Error) *Response {
	return &Response{"2.0", id, nil, &err}
}
