package datastructures

import (
	"strings"
)

type WriteOnlyMap struct {
	EmbeddedMap
}

func NewWriteOnlyMap() WriteOnlyMap {
	return WriteOnlyMap{NewEmbeddedMap()}
}

func (m WriteOnlyMap) Set(key string, val any) {
	m.data[key] = val
}

func (m WriteOnlyMap) Del(key string) {
	delete(m.data, key)
}

func (m WriteOnlyMap) SetPath(path string, val any) {
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

func (m WriteOnlyMap) DelPath(path string) {
	keys := strings.Split(path, ".")

	current := m.data
	for idx, key := range keys {
		if idx == len(keys)-1 {
			delete(current, key)
			return
		}

		if next, exists := current[key]; exists {
			if nextMap, isMap := next.(map[string]any); isMap {
				current = nextMap
				continue
			}
		}
		break
	}
}
