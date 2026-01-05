# TCGDex Go SDK

A Go SDK for the TCGDex API.

[![Go Reference](https://pkg.go.dev/badge/github.com/laiambryant/tcgdex.svg)](https://pkg.go.dev/github.com/laiambryant/tcgdex)
[![Go Report Card](https://goreportcard.com/badge/github.com/laiambryant/tcgdex)](https://goreportcard.com/report/github.com/laiambryant/tcgdex)
[![GitHub license](https://img.shields.io/github/license/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/issues)
[![GitHub stars](https://img.shields.io/github/stars/laiambryant/tcgdex.svg)](https://github.com/laiambryant/tcgdex/stargazers)
[![Coverage Status](https://coveralls.io/repos/github/laiambryant/tcgdex/badge.svg?branch=main)](https://coveralls.io/github/laiambryant/tcgdex?branch=main)

## Installation

```bash
go get github.com/laiambryant/tcgdex
```

## Usage

```go
package main

import (
"context"
"fmt"
"github.com/laiambryant/tcgdex/tcgdex"
)

func main() {
  sdk := tcgdex.New()
  card, err := sdk.Card.Get(context.Background(), "swsh1-1")
  if err != nil {
    panic(err)
  }
  fmt.Println(card.Name)
}
```
