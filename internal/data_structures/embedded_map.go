package datastructures

import (
	"encoding/json"
	"strings"
)

type EmbeddedMap struct {
	data map[string]any
}

func NewEmbeddedMap() EmbeddedMap {
	return EmbeddedMap{make(map[string]any)}
}

func (m EmbeddedMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func (m EmbeddedMap) Has(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m EmbeddedMap) HasPath(path string) bool {
	keys := strings.Split(path, ".")

	current := m.data
	for idx, key := range keys {
		if idx == len(keys)-1 {
			_, exists := current[key]
			return exists
		}
		if next, exists := current[key]; exists {
			if nextMap, isMap := next.(map[string]any); isMap {
				current = nextMap
				continue
			}
		}
		break
	}
	return false
}
