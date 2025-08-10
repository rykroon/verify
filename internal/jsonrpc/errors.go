package jsonrpc

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("jsonrpc.JsonRpcError(Code=%d, Message=%s, Data=%v)", e.Code, e.Message, e.Data)
}

func NewJsonRpcError(code int, msg string, data any) *Error {
	return &Error{code, msg, data}
}

func ParseError(data any) *Error {
	return NewJsonRpcError(-32700, "Parse Error", data)
}

func InvalidRequest(data any) *Error {
	return NewJsonRpcError(-32600, "Invalid Request", data)
}

func MethodNotFound(data any) *Error {
	return NewJsonRpcError(-32601, "Method Not Found", data)
}

func InvalidParams(data any) *Error {
	return NewJsonRpcError(-32602, "Invalid Params", data)
}
