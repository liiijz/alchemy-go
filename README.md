# Alchemy Go SDK

![Status](https://img.shields.io/badge/status-WIP-yellow)
![Status](https://img.shields.io/badge/status-under%20development-orange)

A Go client library for the [Alchemy API](https://www.alchemy.com/), providing easy access to blockchain data and Web3 infrastructure.

## Features

- Clean, idiomatic Go API
- Modular design with namespaced APIs
- Full JSON-RPC 2.0 support
- Configurable HTTP client with timeout and custom options
- Type-safe response structures
- Comprehensive error handling

## Installation

```bash
go get github.com/liiijz/alchemy-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/liiijz/alchemy-go/alchemy"
)

func main() {
    // Create a new Alchemy client
    client := alchemy.NewClient("your-api-key-here")

    // Get token balances for an address
    address := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
    balances, err := client.Token.GetTokenBalances(address)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d tokens\n", len(balances.TokenBalances))
}
```

## Configuration

### Basic Configuration

```go
client := alchemy.NewClient("your-api-key")
```

### Custom Configuration

```go
client := alchemy.NewClient("your-api-key",
    alchemy.WithNetwork("eth-sepolia"),
    alchemy.WithTimeout(60 * time.Second),
)
```

### Available Options

- `WithNetwork(network string)` - Set the blockchain network
- `WithBaseURL(url string)` - Set a custom base URL
- `WithHTTPClient(client *http.Client)` - Use a custom HTTP client
- `WithTimeout(timeout time.Duration)` - Set request timeout

## API Reference

### Token API

#### GetTokenBalances

Get ERC-20 token balances for a specific address.

```go
// Get all ERC-20 tokens
balances, err := client.Token.GetTokenBalances(address)

// Get default top 100 tokens
balances, err := client.Token.GetTokenBalances(address, "DEFAULT_TOKENS")

// Get specific tokens
balances, err := client.Token.GetTokenBalances(
    address,
    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC
    "0x6b175474e89094c44da98b954eedeac495271d0f", // DAI
)
```

#### GetMetadata

Get metadata for a specific token.

```go
metadata, err := client.Token.GetMetadata(contractAddress)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Name: %s\n", *metadata.Name)
fmt.Printf("Symbol: %s\n", *metadata.Symbol)
fmt.Printf("Decimals: %d\n", *metadata.Decimals)
```

#### GetOwners

Get all owners for a given token.

```go
owners, err := client.Token.GetOwners(contractAddress)
```

## Architecture

The SDK uses a clean architecture with separated concerns:

- **AlchemyClient** - Main client for configuration and API assembly
- **HTTPClient** - HTTP/JSON-RPC request engine
- **TokenAPI** - Token-related API methods
- **PortfolioAPI** - Portfolio-related API methods (coming soon)

### Dependency Injection

```
AlchemyClient
  ├── HTTPClient (request engine)
  ├── TokenAPI
  │     └── HTTPClient (shared)
  └── PortfolioAPI
        └── HTTPClient (shared)
```

All API namespaces share the same HTTP client instance, ensuring consistent configuration and efficient resource usage.

## Error Handling

The SDK provides detailed error information:

```go
balances, err := client.Token.GetTokenBalances(address)
if err != nil {
    if alchemy.IsAlchemyError(err) {
        alchemyErr := err.(*alchemy.AlchemyError)
        fmt.Printf("Alchemy error %d: %s\n", alchemyErr.Code, alchemyErr.Message)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

## Examples

See the example implementations in the repository for more use cases.

## Development

### Prerequisites

- Go 1.23 or higher
- Alchemy API key ([Get one here](https://www.alchemy.com/))

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build ./...
```

## Roadmap

- [ ] WebSocket support
- [ ] NFT API methods
- [ ] Enhanced API methods
- [ ] Webhook support
- [ ] Rate limiting
- [ ] Request retries
- [ ] Comprehensive test coverage

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Alchemy](https://www.alchemy.com/) for providing the API
- The Go community for excellent tooling and libraries

## Resources

- [Alchemy Documentation](https://docs.alchemy.com/)
- [Alchemy API Reference](https://docs.alchemy.com/reference/api-overview)
