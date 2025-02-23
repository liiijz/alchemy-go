package alchemy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

// HTTPClient is the core HTTP/JSON-RPC request engine
type HTTPClient struct {
	config *Config
	reqID  atomic.Int64
}

// NewHTTPClient creates a new HTTP client with the given configuration
func NewHTTPClient(config *Config) *HTTPClient {
	return &HTTPClient{
		config: config,
	}
}

// DoRequest performs a JSON-RPC request
func (h *HTTPClient) DoRequest(method string, params ...any) (*JsonRpcResponse, error) {
	reqID := int(h.reqID.Add(1))
	req := NewJsonRpcRequest(reqID, method, params...)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/%s", h.config.BaseURL, h.config.APIKey)
	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := h.config.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var jsonResp JsonRpcResponse
	if err := json.Unmarshal(respBody, &jsonResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if jsonResp.Error != nil {
		return nil, &AlchemyError{
			Code:    jsonResp.Error.Code,
			Message: jsonResp.Error.Message,
			Data:    jsonResp.Error.Data,
		}
	}

	return &jsonResp, nil
}

// GetConfig returns the client configuration
func (h *HTTPClient) GetConfig() *Config {
	return h.config
}
