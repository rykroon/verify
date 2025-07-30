package httpx

import (
	"encoding/json"
	"encoding/xml"
)

type MarshalFunc func(any) ([]byte, error)

func NewBodyUsingMarshalFunc(fn MarshalFunc, v any, contentType string) (*Body, error) {
	data, err := fn(v)
	if err != nil {
		return nil, err
	}
	return NewBody(data, contentType), nil
}

func NewJsonBody(v any) (*Body, error) {
	return NewBodyUsingMarshalFunc(json.Marshal, v, "application/json")
}

func NewXmlBody(v any) (*Body, error) {
	return NewBodyUsingMarshalFunc(xml.Marshal, v, "application/xml")
}
