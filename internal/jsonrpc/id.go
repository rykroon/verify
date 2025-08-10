package jsonrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Id struct {
	raw json.RawMessage
}

func NullId() Id {
	return Id{json.RawMessage(`null`)}
}

func (i Id) String() string {
	return string(i.raw)
}

func (i Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.raw)
}

func (i *Id) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &i.raw); err != nil {
		return err
	}
	if !i.IsValidForRequest() && !i.IsValidForResponse() {
		return errors.New("Invalid jsonrpc id")
	}
	return nil
}

func (i Id) IsNull() bool {
	return string(i.raw) == "null"
}

func (i Id) IsEmpty() bool {
	return len(i.raw) == 0
}

func (i Id) IsString() bool {
	return len(i.raw) != 0 && i.raw[0] == '"'
}

func (i Id) IsInt() bool {
	_, err := strconv.ParseInt(string(i.raw), 10, 64)
	return err == nil
}

func (i Id) IsValidForResponse() bool {
	return i.IsNull() || i.IsString() || i.IsInt()
}

func (i Id) IsValidForRequest() bool {
	return i.IsEmpty() || i.IsString() || i.IsInt()
}
