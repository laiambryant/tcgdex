# TCGDex Go SDK

A Go SDK for the TCGDex API.

[![Go Reference](https://pkg.go.dev/badge/github.com/laiambryant/tcgdex.svg)](https://pkg.go.dev/github.com/laiambryant/tcgdex)
[![Go Report Card](https://goreportcard.com/badge/github.com/laiambryant/tcgdex)](https://goreportcard.com/report/github.com/laiambryant/tcgdex)
[![GitHub license](https://img.shields.io/github/license/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/issues)
[![GitHub stars](https://img.shields.io/github/stars/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/stargazers)
[![Coverage Status](https://coveralls.io/repos/github/laiambryant/tcgdex/badge.svg)](https://coveralls.io/github/laiambryant/tcgdex)

## Installation

```bash
go get github.com/laiambryant/tcgdex
```

## Quick Start

Create an SDK instance with [`tcgdex.New`](tcgdex.go) and use the endpoints:

```go
sdk := tcgdex.New()
card, err := sdk.Card.Get(context.Background(), "swsh1-1")
set, err := sdk.Set.List(context.Background(), nil)
serie, err := sdk.Serie.Get(context.Background(), "swsh")
```

## Configuration

Pass options to [`tcgdex.New`](tcgdex.go):

- `WithBaseURL(url)` - Override API base URL
- `WithUserAgent(ua)` - Set custom User-Agent
- `WithHTTPClient(client)` - Provide custom HTTP client
- `WithCache(ttl)` - Enable response caching

## API

### SDK

- [`TCGDex`](tcgdex.go) - Main SDK type with Card, Set, and Serie endpoints

### Client

- [`client.Client`](client/client.go) - HTTP client for API requests
- [`client.Option`](client/client_options.go) - Configuration options

### Endpoints

- [`endpoint.Endpoint`](endpoint/endpoint.go) - Generic endpoint with Get and List methods
- [`endpoint.DecodeError`](endpoint/errors.go) - JSON decoding error

### Query

- [`query.Query`](query/query.go) - Builder for filter and pagination parameters

### Models

- [`models.Card`](models/card.go) - Card details
- [`models.CardResume`](models/card_resume.go) - Card summary
- [`models.Set`](models/set.go) - Set details
- [`models.SetResume`](models/set_resume.go) - Set summary
- [`models.Serie`](models/serie.go) - Serie details
- [`models.SerieResume`](models/serie_resume.go) - Serie summary

### Enums

- [`enums.Language`](enums/enums.go) - Language codes
- [`enums.Quality`](enums/enums.go) - Image quality levels
- [`enums.Extension`](enums/enums.go) - Image formats
