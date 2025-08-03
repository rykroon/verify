package datastructures

type ReadOnlyMap struct {
	EmbeddedMap
}

func NewReadOnlyMap() ReadOnlyMap {
	return ReadOnlyMap{NewEmbeddedMap()}
}

func (m ReadOnlyMap) Get(key string) (any, bool) {
	val, exists := m.data[key]
	return val, exists
}

func (m ReadOnlyMap) GetBool(key string) (bool, bool) {
	val, exists := m.Get(key)
	if !exists {
		return false, false
	}
	valBool, isInt := val.(bool)
	if !isInt {
		return false, false
	}
	return valBool, true
}

func (m ReadOnlyMap) GetInt(key string) (int, bool) {
	val, exists := m.Get(key)
	if !exists {
		return 0, false
	}
	valInt, isInt := val.(int)
	if !isInt {
		return 0, false
	}
	return valInt, true
}

func (m ReadOnlyMap) GetFloat(key string) (float64, bool) {
	val, exists := m.Get(key)
	if !exists {
		return 0, false
	}
	valFloat, isFloat := val.(float64)
	if !isFloat {
		return 0, false
	}
	return valFloat, true
}

func (m ReadOnlyMap) GetString(key string) (string, bool) {
	val, exists := m.Get(key)
	if !exists {
		return "", false
	}
	valString, isString := val.(string)
	if !isString {
		return "", false
	}
	return valString, true
}
