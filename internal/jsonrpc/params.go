package jsonrpc

import (
	"encoding/json"
	"errors"
)

type Params struct {
	raw json.RawMessage
}

func (p Params) String() string {
	return string(p.raw)
}

func NewPositionalParams(args ...any) (Params, error) {
	var p Params
	data, err := json.Marshal(args)
	if err != nil {
		return p, err
	}
	p.raw = data
	return p, nil
}

func (p *Params) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &p.raw); err != nil {
		return err
	}
	if p.IsZero() || p.IsByName() || p.IsByPosition() {
		return nil
	}
	return errors.New("Invalid params")
}

func (p Params) IsZero() bool {
	return len(p.raw) == 0
}

func (p Params) IsByPosition() bool {
	if p.IsZero() {
		return false
	}
	return p.raw[0] == '['
}

func (p Params) IsByName() bool {
	if p.IsZero() {
		return false
	}
	return p.raw[0] == '{'
}

func (p Params) Unmarshal(v any) error {
	if p.IsZero() {
		return errors.New("params is nil")
	}
	return json.Unmarshal(p.raw, v)
}
