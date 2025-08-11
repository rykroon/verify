package jsonrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Id json.RawMessage

func NullId() Id {
	return Id("null")
}

func (i Id) String() string {
	return string(i)
}

func (id Id) MarshalJSON() ([]byte, error) {
	// overrides default behavior of base64 encoded bytes.
	return id, nil
}

func (i *Id) UnmarshalJSON(data []byte) error {
	*i = append((*i)[:0], data...)
	if i.IsEmpty() || i.IsNull() || i.IsString() || i.IsInt() {
		return nil
	}
	return errors.New("Invalid jsonrpc id")
}

func (i Id) IsNull() bool {
	return string(i) == "null"
}

func (i Id) IsEmpty() bool {
	return len(i) == 0
}

func (i Id) IsString() bool {
	return len(i) != 0 && i[0] == '"'
}

func (i Id) IsInt() bool {
	_, err := strconv.ParseInt(string(i), 10, 64)
	return err == nil
}

func (i Id) IsValidForResponse() bool {
	return i.IsNull() || i.IsString() || i.IsInt()
}

func (i Id) IsValidForRequest() bool {
	return i.IsEmpty() || i.IsString() || i.IsInt()
}
