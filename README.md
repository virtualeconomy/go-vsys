# go-vsys
The official Golang SDK for VSYS APIs

[![License](https://img.shields.io/badge/License-BSD_4--Clause-green.svg)](./LICENSE)

> ***Under active development. Contributions are always welcome!***

The official Golang SDK for VSYS APIs. The [old Golang SDK](https://github.com/virtualeconomy/go-v-sdk) is deprecated and will be archived soon.

- [go-vsys](#go-vsys)
  - [Installation](#installation)
  - [Quick Example](#quick-example)

## Installation

Install from Github

```
go get github.com/virtualeconomy/go-vsys/vsys
```

## Quick Example

```go
package main

import (
	"fmt"

	"github.com/virtualeconomy/go-vsys/vsys"
)

const (
	HOST = "http://veldidina.vos.systems:9928"
	SEED = "your seed"
)

func printHeading(s string) {
	fmt.Printf("============= %s ============\n", s)
}

func main() {
	printHeading("Try out NodeAPI")
	// NodeAPI is the wrapper for REST APIs exposed by VSYS network nodes.
	api := vsys.NewNodeAPI(HOST)
	fmt.Println(api)

	// GET /blocks/height
	resp, _ := api.GetCurBlockHeight()
	fmt.Println(resp.Height)

	printHeading("Try out Chain")
	// Chain represents the blockchain.
	ch := vsys.NewChain(api, vsys.TEST_NET)
	fmt.Println(ch)

	// Get chain's height
	height, _ := ch.Height()
	fmt.Println(height)

	printHeading("Try out Account")
	// Wallet represents the wallet where there's a seed phrase.
	wal, _ := vsys.NewWalletFromSeedStr(SEED)
	fmt.Println(wal)

	// Account represents an account in the blockchain network.
	acnt, _ := wal.GetAccount(ch, 0)
	fmt.Println(acnt)

	printHeading("Try out Smart Contract")
	const ctrtId = "CF3cK7TJFfw1AcPk74osKyGeGxee6u5VNXD"
	nc, _ := vsys.NewNFTCtrt(ctrtId, ch)

	// Get the contract's ID
	fmt.Println("Contract id:", nc.CtrtId)
}
```

## Testing

- timeout is needed to avoid default 10min limit for running tests. With -run flag you can supply regexp for tests to run.
```
go test ./vsys/ -v -timeout 0 -run AsWhole
```
