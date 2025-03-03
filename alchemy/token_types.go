package alchemy

// TokenBalancesResponse represents the response from alchemy_getTokenBalances
type TokenBalancesResponse struct {
	Address      string         `json:"address"`
	TokenBalances []TokenBalance `json:"tokenBalances"`
}

// TokenBalance represents a single token balance
type TokenBalance struct {
	ContractAddress string  `json:"contractAddress"`
	TokenBalance    *string `json:"tokenBalance"` // hex string, can be null
	Error           *string `json:"error,omitempty"`
}

// TokenMetadata represents token metadata information
type TokenMetadata struct {
	Name     *string `json:"name,omitempty"`
	Symbol   *string `json:"symbol,omitempty"`
	Decimals *int    `json:"decimals,omitempty"`
	Logo     *string `json:"logo,omitempty"`
}
