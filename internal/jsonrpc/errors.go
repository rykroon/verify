package jsonrpc

import "fmt"

const (
	ErrorCodeParseError     = -32700
	ErrorCodeInvalidRequest = -32600
	ErrorCodeMethodNotFound = -32601
	ErrorCodeInvalidParams  = -32602
	ErrorCodeInternalError  = -32603
)

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
	return NewJsonRpcError(ErrorCodeParseError, "Parse Error", data)
}

func InvalidRequest(data any) *Error {
	return NewJsonRpcError(ErrorCodeInvalidRequest, "Invalid Request", data)
}

func MethodNotFound(data any) *Error {
	return NewJsonRpcError(ErrorCodeMethodNotFound, "Method Not Found", data)
}

func InvalidParams(data any) *Error {
	return NewJsonRpcError(ErrorCodeInvalidParams, "Invalid Params", data)
}

func InternalError(data any) *Error {
	return NewJsonRpcError(ErrorCodeInternalError, "InternalError", data)
}
