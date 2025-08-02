package httpx

import (
	"encoding/json"
)

type UnmarshalFunc func([]byte, any) error

func (b *Body) UnmarshalUsingFunc(v any, fn UnmarshalFunc) error {
	return fn(b.data, v)
}

// json

func (b *Body) IsJson() bool {
	return b.contentType == "application/json"
}

func (b *Body) UnmarshalJson(v any) error {
	return b.UnmarshalUsingFunc(v, json.Unmarshal)
}
