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
	if p.IsEmpty() || p.IsObject() || p.IsArray() {
		return nil
	}
	return errors.New("Invalid params")
}

func (p Params) IsEmpty() bool {
	return len(p) == 0
}

func (p Params) IsArray() bool {
	return len(p) != 0 && p[0] == '['
}

func (p Params) IsObject() bool {
	return len(p) != 0 && p[0] == '{'
}

func (p Params) UnmarshalTo(v any) error {
	if p.IsEmpty() {
		return errors.New("params is nil")
	}
	return json.Unmarshal(p, v)
}
