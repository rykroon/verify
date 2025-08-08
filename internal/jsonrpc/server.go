package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	if !json.Valid(bodyBytes) {
		http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	decoder.DisallowUnknownFields()
	var req JsonRpcRequest
	if err := decoder.Decode(&req); err != nil {
		s.writeError(w, req.Id, -32600, "Invalid Request", err.Error())
		return
	}

	if req.JsonRpc != "2.0" {
		s.writeError(w, req.Id, -32600, "Invalid Request", "jsonrpc must be '2.0'")
		return
	}

	if !validParams(req.Params) {
		s.writeError(w, req.Id, -32600, "Invalid Request", "invalid params")
		return
	}

	if req.Method == "" {
		s.writeError(w, req.Id, -32600, "Invalid Request", "missing method")
		return
	}

	if !validId(req.Id) {
		s.writeError(w, req.Id, -32600, "Invalid Request", "invalid id")
		return
	}

	handler, exists := s.methods[req.Method]
	if !exists {
		s.writeError(w, req.Id, -32601, "Method not found", nil)
		return
	}

	if req.IsNotification() {
		fmt.Println("Handle notification")
		w.WriteHeader(http.StatusNoContent) // 204
		//w.Write([]byte(""))
		go handler(r.Context(), req.Params)
		return
	}

	result, rpcErr := handler(r.Context(), req.Params)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	if rpcErr != nil {
		resp := NewJsonRpcErrorResponse(req.Id, *rpcErr)
		encoder.Encode(resp)
	} else {
		resp := NewJsonRpcSuccessResponse(req.Id, result)
		encoder.Encode(resp)
	}
}

func (s *JsonRpcServer) writeError(w http.ResponseWriter, id json.RawMessage, code int, message string, data any) {
	resp := NewJsonRpcErrorResponse(id, *NewJsonRpcError(code, message, data))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
