# go-vsys
The official Golang SDK for VSYS APIs

[![License](https://img.shields.io/badge/License-BSD_4--Clause-green.svg)](./LICENSE)

> ***Under active development. Contributions are always welcome!***

The official Golang SDK for VSYS APIs. The [old Golang SDK](https://github.com/virtualeconomy/go-v-sdk) is deprecated and will be archived soon.

- [go-vsys](#go-vsys)
	- [Installation](#installation)
	- [Quick Example](#quick-example)
	- [Docs](#docs)
		- [Account & Wallet](#account--wallet)
		- [Chain & API](#chain--api)
		- [Smart Contracts](#smart-contracts)
	- [Testing](#testing)
	- [Contributing](#contributing)

## Installation

Install from Github using `go get`:

```bash
go get github.com/virtualeconomy/go-vsys
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

## Docs
### Account & Wallet
- [Account](./doc/account.md)
- [Wallet](./doc/wallet.md)
### Chain & API
- [Chain](./doc/chain.md)
- [Api](./doc/api.md)

### Smart Contracts
- [NFT Contract V1](./doc/smart_contract/nft_ctrt.md)
- [NFT Contract V2](./doc/smart_contract/nft_ctrt_v2.md)
- [Token Contract V1 without split](./doc/smart_contract/tok_ctrt_no_split.md)
- [Token Contract V1 with split](./doc/smart_contract/tok_ctrt_split.md)
- [Token Contract V2 without split](./doc/smart_contract/tok_ctrt_no_split_v2.md)
- [Atomic Swap Contract](./doc/smart_contract/atomic_swap_ctrt.md)
- [Payment Channel Contract](./doc/smart_contract/pay_chan_ctrt.md)
- [Lock Contract](./doc/smart_contract/lock_ctrt.md)
- [System Contract](./doc/smart_contract/sys_ctrt.md)
- [V Escrow Contract](./doc/smart_contract/v_escrow_ctrt.md)
- [V Option Contract](./doc/smart_contract/v_option_ctrt.md)
- [V Stable Swap Contract](./doc/smart_contract/v_stable_swap_ctrt.md)
- [V Swap Contract](./doc/smart_contract/v_swap_ctrt.md)

## Testing

Functional tests are functions that simulate the behavior of a normal user to interact with `py_vsys`(e.g. register a smart contract & call fucntions of it).

> NOTE that the test environement defined as global variables in [helper_test.go](./vsys/helper_test.go) has to be configured through environemnt variables before the test cases can be executed.

Then go to the root of the project and run.
```bash
go test ./vsys/ -v -timeout 0 -run AsWhole
```
> timeout is needed to avoid default 10min limit for running tests. With -run flag you can supply regexp for tests to run.

The above command will run functions that tests contracts as a whole i.e. without setting up new resources per each function test. If you want to test individual contracts or functions:
```bash
go test ./vsys/ -v -timeout 0 -run NFTCtrt_Supersede
```
> Test function names are made so that you can run all functions of certain contract by supplying its name to `-run` argument. Then can add function name as suffix to test specific function.

More information on how to run tests can be found in [testing package](https://pkg.go.dev/testing) documentation.

## Contributing

**Contributions are always welcome!**

See [the development documentation](./doc/dev.md) for more details and please adhere to conventions mentioned in it.
