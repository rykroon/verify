package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRpcHandler interface {
	ServerJsonRpc(ctx context.Context, req *Request) *Response
}

type HandlerFunc func(ctx context.Context, params Params) (any, *Error)

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

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var req Request

	if err := decoder.Decode(&req); err != nil {
		s.writeError(w, NullId(), ParseError(err.Error()))
		return
	}

	fmt.Println(req.Jsonrpc, req.Method, req.Id, req.Params)

	if !req.Id.IsValidForRequest() {
		s.writeError(w, NullId(), InvalidRequest("invalid id"))
		return
	}

	// validate Request
	if req.Jsonrpc != "2.0" {
		s.writeError(w, req.Id, InvalidRequest("jsonrpc must be 2.0"))
		return
	}

	resp := s.ServeJsonRpc(r.Context(), req)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		s.writeError(w, req.Id, nil) // check later
	}
}

func (s *JsonRpcServer) ServeJsonRpc(ctx context.Context, req Request) *Response {
	handler, exists := s.methods[req.Method]
	if !exists {
		return NewJsonRpcErrorResp(req.Id, MethodNotFound(req.Method))
	}

	if req.IsNotification() {
		go handler(ctx, req.Params)
		return nil
	}

	result, rpcErr := handler(ctx, req.Params)
	if rpcErr != nil {
		return NewJsonRpcSuccessResp(req.Id, result)

	} else {
		return NewJsonRpcErrorResp(req.Id, rpcErr)
	}
}

func (s *JsonRpcServer) writeError(w http.ResponseWriter, id Id, jsonRpcErr *Error) {
	resp := NewJsonRpcErrorResp(id, jsonRpcErr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
