package jsonrpc

import (
	"encoding/json"
)

type RawRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Id      json.RawMessage `json:"id"`
}

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params,omitzero"`
	Id      Id     `json:"id,omitzero"`
}

// func (r *Request) UnmarshalJSON(data []byte) error {
// 	decoder := json.NewDecoder(bytes.NewReader(data))
// 	decoder.DisallowUnknownFields()
// 	var raw RawRequest
// 	if err := decoder.Decode(&raw); err != nil {
// 		return fmt.Errorf("failed to decode jsonrpc request: %w", err)
// 	}
// 	if raw.Jsonrpc != "2.0" {
// 		return InvalidRequest("jsonrpc must be 2.0")
// 	}
// 	r.Jsonrpc = raw.Jsonrpc

// 	if raw.Method == "" {
// 		return InvalidRequest("must provide a method")
// 	}
// 	r.Method = raw.Method

// 	if err := json.Unmarshal(raw.Id, &r.Id); err != nil {
// 		return fmt.Errorf("failed to decode id: %w", err)
// 	}
// 	if !r.Id.IsValidForRequest() {
// 		return InvalidRequest("invalid id")
// 	}

// 	if err := json.Unmarshal(raw.Params, &r.Params); err != nil {
// 		return fmt.Errorf("failed to decode params: %w", err)
// 	}

// 	return nil
// }

func (r Request) IsNotification() bool {
	return r.Id.IsEmpty()
}
