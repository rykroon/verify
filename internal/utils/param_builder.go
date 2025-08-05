package utils

import (
	"encoding/json"
	"strings"
)

type ParamBuilder struct {
	data map[string]any
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{make(map[string]any)}
}

func (p *ParamBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.data)
}

func (p *ParamBuilder) Set(key string, val any) {
	p.data[key] = val
}

func (m *ParamBuilder) SetPath(path string, val any) {
	keys := strings.Split(path, ".")

	current := m.data
	// If the key doesn't exist or isn't a map, create a new map
	for i, key := range keys {
		if i == len(keys)-1 { // last key
			current[key] = val
			return
		}
		if next, exists := current[key]; exists {
			if nextMap, isMap := next.(map[string]any); isMap {
				current = nextMap
			} else {
				// Overwrite non-map with a new map
				newMap := make(map[string]any)
				current[key] = newMap
				current = newMap
			}
		} else {
			// create new map
			newMap := make(map[string]any)
			current[key] = newMap
			current = newMap
		}
	}
}
