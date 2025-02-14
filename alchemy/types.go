package alchemy

// JsonRpcRequest represents a JSON-RPC 2.0 request
type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

// JsonRpcResponse represents a JSON-RPC 2.0 response
type JsonRpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Result  any           `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

// JsonRpcError represents a JSON-RPC 2.0 error
type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// NewJsonRpcRequest creates a new JSON-RPC request
func NewJsonRpcRequest(id int, method string, params ...any) *JsonRpcRequest {
	return &JsonRpcRequest{
		JsonRpc: "2.0",
		Id:      id,
		Method:  method,
		Params:  params,
	}
}
