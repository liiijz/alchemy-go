package alchemy

// AlchemyClient is the main Alchemy API client
// It is responsible for configuration and assembling API namespaces
type AlchemyClient struct {
	httpClient *HTTPClient

	// API namespaces
	Portfolio *PortfolioAPI
	Token     *TokenAPI
}

// NewClient creates a new Alchemy client with the given API key and options
func NewClient(apiKey string, opts ...Option) *AlchemyClient {
	// Create configuration
	config := DefaultConfig(apiKey)
	for _, opt := range opts {
		opt(config)
	}

	// Create HTTP client (request engine)
	httpClient := NewHTTPClient(config)

	// Assemble client with API namespaces
	client := &AlchemyClient{
		httpClient: httpClient,
		Portfolio:  &PortfolioAPI{httpClient: httpClient},
		Token:      &TokenAPI{httpClient: httpClient},
	}

	return client
}

// GetHTTPClient returns the underlying HTTP client
func (c *AlchemyClient) GetHTTPClient() *HTTPClient {
	return c.httpClient
}

// GetConfig returns the client configuration
func (c *AlchemyClient) GetConfig() *Config {
	return c.httpClient.GetConfig()
}
