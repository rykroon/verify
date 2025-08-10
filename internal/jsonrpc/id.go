package jsonrpc

import (
	"encoding/json"
	"strconv"
)

type JsonRpcId struct {
	json.RawMessage
}

// func (i JsonRpcId) UnmarshalJSON(data []byte) error {
// 	err := json.Unmarshal(data, &i)
// 	if err != nil {
// 		return err
// 	}
// 	if len(i) == 0 {
// 		return nil
// 	}
// 	if i[0] == '"' {
// 		return nil
// 	}
// 	_, err = strconv.ParseInt(string(i), 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	return errors.New("Invalid jsonrpc id")
// }

func (i JsonRpcId) IsMissing() bool {
	return len(i.RawMessage) == 0
}

func (i JsonRpcId) IsString() bool {
	if i.IsMissing() {
		return false
	}
	return i.RawMessage[0] == '"'
}

func (i JsonRpcId) IsInt() bool {
	if i.IsMissing() {
		return false
	}
	_, err := strconv.ParseInt(string(i.RawMessage), 10, 64)
	return err == nil
}

func (i JsonRpcId) IsValid() bool {
	return i.IsMissing() || i.IsString() || i.IsInt()
}
