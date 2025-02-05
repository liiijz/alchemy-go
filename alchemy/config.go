package alchemy

import (
	"net/http"
	"time"
)

// Config holds the configuration for the Alchemy client
type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Network    string
}

// Option is a function that configures a Config
type Option func(*Config)

// DefaultConfig returns a new Config with default values
func DefaultConfig(apiKey string) *Config {
	return &Config{
		APIKey:  apiKey,
		BaseURL: "https://eth-mainnet.g.alchemy.com/v2",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Network: "eth-mainnet",
	}
}

// WithBaseURL sets a custom base URL
func WithBaseURL(url string) Option {
	return func(c *Config) {
		c.BaseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = client
	}
}

// WithNetwork sets the network
func WithNetwork(network string) Option {
	return func(c *Config) {
		c.Network = network
		// Update BaseURL based on network
		c.BaseURL = "https://" + network + ".g.alchemy.com/v2"
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.HTTPClient.Timeout = timeout
	}
}
