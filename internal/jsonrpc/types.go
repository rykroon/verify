package jsonrpc

func NewJsonRpcSuccessResponse(id any, result any) map[string]any {
	return map[string]any{"jsonrpc": "2.0", "id": id, "result": result}
}

func NewJsonRpcErrorResponse(id any, err JsonRpcError) map[string]any {
	return map[string]any{"jsonrpc": "2.0", "id": id, "error": err}
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewJsonRpcError(code int, msg string, data any) *JsonRpcError {
	return &JsonRpcError{code, msg, data}
}
