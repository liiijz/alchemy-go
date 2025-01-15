package alchemy

type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type JsonRpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      int           `json:"id"` 
	Result  any           `json:"result"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

type JsonRpcError struct {
}
