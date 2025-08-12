package jsonrpc

import (
	"encoding/json"
	"errors"
)

type Params json.RawMessage

func (p Params) String() string {
	return string(p)
}

func (p *Params) UnmarshalJSON(data []byte) error {
	*p = append((*p)[:0], data...)
	if p.IsEmpty() || p.IsNamed() || p.IsPositional() {
		return nil
	}
	return errors.New("not a valid params type")
}

func (p Params) IsEmpty() bool {
	return len(p) == 0
}

func (p Params) IsPositional() bool {
	return len(p) != 0 && p[0] == '['
}

func (p Params) IsNamed() bool {
	return len(p) != 0 && p[0] == '{'
}

func (p Params) UnmarshalTo(v any) error {
	if p.IsEmpty() {
		return errors.New("params is nil")
	}
	if positional, ok := v.(Positional); ok && p.IsPositional() {
		pointers := positional.GetParamPointers()
		return json.Unmarshal(p, &pointers)
	}
	return json.Unmarshal(p, v)
}

type Positional interface {
	GetParamPointers() []any
}
