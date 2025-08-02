package datastructures

import (
	"encoding/json"
	"strings"
)

type WriteOnlyMap struct {
	data map[string]any
}

func NewWriteOnlyMap() WriteOnlyMap {
	return WriteOnlyMap{make(map[string]any)}
}

func (m WriteOnlyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func (m WriteOnlyMap) SetString(key, val string) {
	m.data[key] = val
}

func (m WriteOnlyMap) SetBool(key string, val bool) {
	m.data[key] = val
}

func (m WriteOnlyMap) SetInt(key string, val int) {
	m.data[key] = val
}

func (m WriteOnlyMap) SetFloat(key string, val float64) {
	m.data[key] = val
}

func (m WriteOnlyMap) SetStringToPath(path, val string) {
	m.setAnyToPath(path, val)
}

func (m WriteOnlyMap) SetBoolToPath(path string, val bool) {
	m.setAnyToPath(path, val)
}

func (m WriteOnlyMap) SetIntToPath(path string, val int) {
	m.setAnyToPath(path, val)
}

func (m WriteOnlyMap) SetFloatToPath(path string, val float64) {
	m.setAnyToPath(path, val)
}

func (m WriteOnlyMap) setAnyToPath(path string, val any) {
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
