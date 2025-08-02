package datastructures

import "strings"

type NestedMap struct {
	m map[string]any
}

func NewNestedMap() *NestedMap {
	return &NestedMap{make(map[string]any)}
}

func (m *NestedMap) Set(key string, val any) {
	m.m[key] = val
}

func (m *NestedMap) SetPath(path string, val any) {
	keys := strings.Split(path, ".")

	current := m.m
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
