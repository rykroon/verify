package jsonrpc

import (
	"context"
	"encoding/json"
	"net/http"
)

type JsonRpcHandler interface {
	ServerJsonRpc(ctx context.Context, req *JsonRpcRequest) *JsonRpcResponse
}

type HandlerFunc func(ctx context.Context, params JsonRpcParams) (any, *JsonRpcError)

type JsonRpcServer struct {
	methods map[string]HandlerFunc
	parser  JsonRpcRequestParser
}

func NewJsonRpcServer() *JsonRpcServer {
	return &JsonRpcServer{
		methods: make(map[string]HandlerFunc),
		parser:  SimpleParser{},
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

	req, err := s.parser.ParseRequest(r.Body)
	if err != nil {
		s.writeError(w, nil, err)
		return
	}

	if req == nil {
		s.writeError(w, nil, InvalidRequest(nil))
		return
	}

	resp := s.ServeJsonRpc(r.Context(), *req)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		s.writeError(w, req.Id, nil) // check later
	}
}

func (s *JsonRpcServer) ServeJsonRpc(ctx context.Context, req JsonRpcRequest) *JsonRpcResponse {
	handler, exists := s.methods[req.Method]
	if !exists {
		return NewJsonRpcErrorResp(req.Id, MethodNotFound(req.Method))
	}

	if req.Id.IsMissing() {
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

func (s *JsonRpcServer) writeError(w http.ResponseWriter, id any, jsonRpcErr *JsonRpcError) {
	resp := NewJsonRpcErrorResp(id, jsonRpcErr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
