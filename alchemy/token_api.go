package alchemy

import (
	"encoding/json"
	"fmt"
)

// TokenAPI provides access to Alchemy's Token API
type TokenAPI struct {
	httpClient *HTTPClient
}

// GetTokenBalances retrieves ERC-20 token balances for a specific owner address
// contractAddresses can be:
// - nil or empty: queries all ERC-20 tokens
// - []string{"DEFAULT_TOKENS"}: queries top 100 tokens by 24h volume
// - []string{address1, address2, ...}: queries specific token contracts
func (t *TokenAPI) GetTokenBalances(ownerAddress string, contractAddresses ...string) (*TokenBalancesResponse, error) {
	var params []any

	// First parameter is always the owner address
	params = append(params, ownerAddress)

	// Second parameter is the list of contract addresses or special keyword
	if len(contractAddresses) == 0 {
		// Query all ERC-20 tokens
		params = append(params, "erc20")
	} else if len(contractAddresses) == 1 && contractAddresses[0] == "DEFAULT_TOKENS" {
		// Query default tokens
		params = append(params, "DEFAULT_TOKENS")
	} else {
		// Query specific contracts
		params = append(params, contractAddresses)
	}

	resp, err := t.httpClient.DoRequest("alchemy_getTokenBalances", params...)
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		return nil, fmt.Errorf("empty result from API")
	}

	// Marshal and unmarshal to convert to typed response
	data, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	var result TokenBalancesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return &result, nil
}

// GetMetadata retrieves metadata for a specific token
func (t *TokenAPI) GetMetadata(contractAddress string) (*TokenMetadata, error) {
	resp, err := t.httpClient.DoRequest("alchemy_getTokenMetadata", contractAddress)
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		return nil, fmt.Errorf("empty result from API")
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	var metadata TokenMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return &metadata, nil
}

// GetOwners retrieves all owners for a given token
func (t *TokenAPI) GetOwners(contractAddress string) (map[string]any, error) {
	resp, err := t.httpClient.DoRequest("alchemy_getOwnersForToken", contractAddress)
	if err != nil {
		return nil, err
	}

	result := make(map[string]any)
	if resp.Result != nil {
		data, err := json.Marshal(resp.Result)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
