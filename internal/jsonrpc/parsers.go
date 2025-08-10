package jsonrpc

import (
	"encoding/json"
	"fmt"
	"io"
)

type JsonRpcRequestParser interface {
	// parse request is meant for parsing AND validating the request.
	ParseRequest(io.Reader) (*JsonRpcRequest, *JsonRpcError)
}

type SimpleParser struct{}

func (p SimpleParser) ParseRequest(reader io.Reader) (*JsonRpcRequest, *JsonRpcError) {
	var req *JsonRpcRequest
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		return nil, ParseError(err.Error())
	}
	if req == nil {
		return nil, InvalidRequest(nil)
	}
	if req.Jsonrpc != "2.0" {
		return nil, InvalidRequest(`"jsonrpc" must be 2.0`)
	}

	if !req.Params.IsValid() {
		return nil, InvalidRequest("invalid params type")
	}

	if !req.Id.IsValid() {
		return nil, InvalidParams("invalid id")
	}
	return req, nil
}

type MapParser struct{}

func (p MapParser) ParseRequest(reader io.Reader) (*JsonRpcRequest, *JsonRpcError) {
	var raw map[string]json.RawMessage
	if err := json.NewDecoder(reader).Decode(&raw); err != nil {
		return nil, ParseError(err.Error())
	}

	var version string
	if err := json.Unmarshal(raw["jsonrpc"], &version); err != nil || version != "2.0" {
		return nil, InvalidRequest("jsonrpc must be '2.0'")
	}
	delete(raw, "jsonrpc")

	var method string
	if err := json.Unmarshal(raw["method"], &method); err != nil {
		return nil, InvalidRequest("missing method")
	}
	delete(raw, "method")

	id := JsonRpcId{raw["id"]}
	if !id.IsValid() {
		return nil, InvalidRequest("invalid id")
	}

	params := JsonRpcParams(raw["params"])
	if !params.IsValid() {
		return nil, InvalidRequest("invalid params")
	}
	delete(raw, "params")

	if len(raw) > 0 {
		keys := make([]string, 0, len(raw))
		for k := range raw {
			keys = append(keys, k)
		}
		return nil, InvalidRequest(fmt.Sprintf("unexpected fields: %v", keys))
	}

	return &JsonRpcRequest{"2.0", method, params, id}, nil
}

type StreamParser struct{}

func (p StreamParser) ParseRequest(reader io.Reader) (*JsonRpcRequest, *JsonRpcError) {
	return nil, nil // do later
}
