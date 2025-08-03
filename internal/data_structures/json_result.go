package datastructures

import "encoding/json"

type RawJson[T any] struct {
	data json.RawMessage
}

func NewResult[T any](data []byte) (*RawJson[T], error) {
	var raw json.RawMessage
	err := json.Unmarshal(data, &raw)
	return &RawJson[T]{raw}, err
}

func (r *RawJson[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &r.data)
}

func (r *RawJson[T]) ToMap() (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal(r.data, &result)
	return result, err
}

func (r *RawJson[T]) ToStruct() (T, error) {
	var result T
	err := json.Unmarshal(r.data, &result)
	return result, err
}
