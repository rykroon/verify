package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	var raw map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		s.writeError(w, json.RawMessage("null"), -32700, "Parse Error", err.Error())
		return
	}

	var version string
	if err := json.Unmarshal(raw["jsonrpc"], &version); err != nil || version != "2.0" {
		s.writeError(w, json.RawMessage("null"), -32600, "Invalid Request", "jsonrpc must be '2.0'")
		return
	}
	delete(raw, "jsonrpc")

	var method string
	if err := json.Unmarshal(raw["method"], &method); err != nil {
		s.writeError(w, nil, -32600, "Invalid Request", "missing method")
		return
	}
	delete(raw, "method")

	var id any
	if rawId, exists := raw["id"]; exists {
		decoder := json.NewDecoder(bytes.NewReader(rawId))
		decoder.UseNumber()
		if err := decoder.Decode(&id); err != nil {
			s.writeError(w, nil, -32600, "Invalid Request", "invalid id")
			return
		}
		switch v := id.(type) {
		case string:
		case json.Number:
			if _, err := v.Int64(); err != nil {
				s.writeError(w, nil, -32600, "Invalid Request", "invalid id type float")
				return
			}
		default:
			s.writeError(w, nil, -32600, "Invalid Request", "invalid id type")
			return
		}
		delete(raw, "id")
	}

	params := raw["params"]
	delete(raw, "params")

	if len(raw) > 0 {
		keys := make([]string, 0, len(raw))
		for k := range raw {
			keys = append(keys, k)
		}
		s.writeError(w, nil, -32600, "Invalid Request", fmt.Sprintf("unexpected fields: %v", keys))
		return
	}

	handler, exists := s.methods[method]
	if !exists {
		s.writeError(w, id, -32601, "Method not found", nil)
		return
	}

	if id == nil {
		fmt.Println("Handle notification")
		w.WriteHeader(http.StatusNoContent) // 204
		w.Write([]byte(""))
		go handler(r.Context(), params)
		return
	}

	result, rpcErr := handler(r.Context(), params)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	var resp map[string]any
	if rpcErr != nil {
		resp = NewJsonRpcErrorResponse(id, *rpcErr)

	} else {
		resp = NewJsonRpcSuccessResponse(id, result)
	}
	encoder.Encode(resp)
}

func (s *JsonRpcServer) writeError(w http.ResponseWriter, id any, code int, message string, data any) {
	resp := NewJsonRpcErrorResponse(id, *NewJsonRpcError(code, message, data))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
