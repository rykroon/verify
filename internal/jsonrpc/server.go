package jsonrpc

import (
	"context"
	"encoding/json"
	"net/http"
)

type HandlerFunc func(ctx context.Context, params json.RawMessage) (any, *JsonRpcError)

type JsonRpcServer struct {
	methods map[string]HandlerFunc
}

func NewJsonRpcServer() *JsonRpcServer {
	return &JsonRpcServer{
		methods: make(map[string]HandlerFunc),
	}
}

func (s *JsonRpcServer) Register(method string, handler HandlerFunc) {
	s.methods[method] = handler
}

func (s *JsonRpcServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var req JsonRpcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, req.Id, -32700, "Parse error", err)
		return
	}

	if req.Id == nil {
		//w.WriteHeader(http.StatusNoContent) // 204
		//go processNotification(req)
	}

	if req.JsonRpc != "2.0" {
		s.writeError(w, req.Id, -32600, "Invalid Request", "jsonrpc must be '2.0'")
		return
	}

	handler, ok := s.methods[req.Method]
	if !ok {
		s.writeError(w, req.Id, -32601, "Method not found", nil)
		return
	}

	result, rpcErr := handler(r.Context(), req.Params)

	var resp *JsonRpcResponse
	if rpcErr != nil {
		resp = NewJsonRpcErrorResponse(req.Id, rpcErr)
	} else {
		resp = NewJsonRpcSuccessResponse(req.Id, result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *JsonRpcServer) writeError(w http.ResponseWriter, id *json.RawMessage, code int, message string, data any) {
	resp := NewJsonRpcErrorResponse(id, &JsonRpcError{
		Code:    code,
		Message: message,
		Data:    data,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
