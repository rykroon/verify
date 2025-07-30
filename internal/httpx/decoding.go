package httpx

import (
	"encoding/json"
	"encoding/xml"
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

// xml

func (b *Body) IsXml() bool {
	return b.contentType == "application/xml"
}

func (b *Body) UnmarshalXml(v any) error {
	return b.UnmarshalUsingFunc(v, xml.Unmarshal)
}
