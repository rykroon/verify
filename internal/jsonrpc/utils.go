package jsonrpc

import (
	"encoding/json"
)

func isNull(raw json.RawMessage) bool {
	return string(raw) == "null"
}

func isMissing(raw json.RawMessage) bool {
	return len(raw) == 0
}

func isEmptyString(raw json.RawMessage) bool {
	return string(raw) == `""`
}

func validId(id json.RawMessage) bool {
	if isMissing(id) {
		// notification
		return true
	}

	if isNull(id) {
		// do not allow null ids.
		return false
	}

	// Try string
	var s string
	if err := json.Unmarshal(id, &s); err == nil {
		return true
	}

	// Try int
	var i int
	if err := json.Unmarshal(id, &i); err == nil {
		return true
	}

	return false
}

func validParams(params json.RawMessage) bool {
	if isMissing(params) {
		return true
	}

	//check map
	var m map[string]any
	if err := json.Unmarshal(params, &m); err == nil {
		return true
	}

	var a []any
	if err := json.Unmarshal(params, &a); err == nil {
		return true
	}

	return false
}
