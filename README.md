[![Go](https://github.com/Henelik/qtrade-api-go/actions/workflows/go.yml/badge.svg)](https://github.com/Henelik/qtrade-api-go/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/Henelik/qtrade-api-go/branch/master/graph/badge.svg?token=WE6RKWXNH2)](https://codecov.io/gh/Henelik/qtrade-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Henelik/qtrade-api-go)](https://goreportcard.com/report/github.com/Henelik/qtrade-api-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# qTrade API Go <a href="https://qtrade.io"><img src="https://qtrade.io/images/logo.png" alt="qTrade" width="32"/></a>

This is an unofficial Go client for the [qTrade.io](https://qtrade.io) crypto exchange API.

The client provides helpful methods, data structures, and enums to make the experience of using the qTrade API in Go as seamless as possible.

## Features

* All documented API methods are implemented
* Automatic HMAC signature generation
* Automatic API error checking and parsing
* Enumerated data types for Markets, Currencies, and Order Types
* Automatic rate limit waiting
* Configurable retries

## Documentation

Instantiating a client and making a request is easy:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Henelik/qtrade-api-go/qtrade/v1"
)

func main() {
	config := qtrade.Configuration{
		HMACKeypair: "1:111111111111111111111111111",
		Endpoint:    "https://api.qtrade.io",
		Timeout:     time.Second * 60,
	}

	client, err := qtrade.NewClient(config)
	if err != nil {
		panic(err)
	}

	balances, err := client.GetBalances(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	
	for _, balance := range balances {
		fmt.Printf("Balance for %s: %v", balance.Currency, balance.Balance)
    }
}
```

Please refer to the [official documentation](https://qtrade-exchange.github.io/qtrade-docs) for more information.

## Planned Features
* Helper functions (e.g. cancel all orders on a given market)
* Improve client documentation
