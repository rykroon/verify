package jsonrpc

import (
	"encoding/json"
	"errors"
)

type JsonRpcParams json.RawMessage

func (p JsonRpcParams) IsMissing() bool {
	return len(p) == 0
}

func (p JsonRpcParams) IsByPosition() bool {
	if p.IsMissing() {
		return false
	}
	return p[0] == '['
}

func (p JsonRpcParams) IsByName() bool {
	if p.IsMissing() {
		return false
	}
	return p[0] == '{'
}

func (p JsonRpcParams) IsValid() bool {
	return p.IsMissing() || p.IsByPosition() || p.IsByName()
}

func (p JsonRpcParams) Unmarshal(v any) error {
	if p.IsMissing() {
		return errors.New("params is nil")
	}
	return json.Unmarshal(p, v)
}
