package jsonrpc

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params,omitzero"`
	Id      Id     `json:"id,omitzero"`
}

func (r Request) IsNotification() bool {
	return r.Id.IsEmpty()
}

func (r Request) IdForResponse() Id {
	if r.Id.IsEmpty() {
		return NullId()
	}
	return r.Id
}
