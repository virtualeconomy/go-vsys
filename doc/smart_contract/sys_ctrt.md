# System Contract

- [System Contract](#system-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [Actions](#actions)
      - [Send](#send)
      - [Deposit](#deposit)
      - [Withdraw](#withdraw)
      - [Transfer](#transfer)

## Introduction

The System Contract on V Systems is quite unique in that it is directly included in the protocol and not registered by users. Since Contract variables and VSYS tokens use different databases, it is normally not possible for them to interact. However, the System Contract handles the mixing of these two databases, and allows users to do things such as deposit and withdraw VSYS token from contracts.

## Usage with Go SDK

### Registration

```go
// acnt: *vsys.Account
// ch: *vsys.Chain

// initial a new system contract on testnet
mainnetCtrt := vsys.NewSysCtrtForMainnet(acnt.Chain)
// initial a new system contract on testnet
testnetCtrt := vsys.NewSysCtrtForTestnet(acnt.Chain)

fmt.Println(mainnetCtrt.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CCL1QGBqPAaFjYiA8NMGVhzkd3nJkGeKYBq))
```

### Actions

#### Send

Send VSYS tokens to another user.

```go
// sc: *vsys.SysCtrt
// acnt0: *vsys.Account
// acnt1: *vsys.Account
// amount: float64

resp, err = sc.Send(
  acnt0, // by
  acnt1.Addr.B58Str().Str(), // receiver
  10, // amount
  "". // attachment
)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(GbkCZzjLjp57VpacQJiLrVkJfCAB6gTtkc2psY7o7t9i) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659931522952066000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(43BuiGTeX3WcHH5VxPe4RaNFSpjPHzwB8iSphuEzef4ASzFKwXagtoZy1ESeRXsnsQ9RanGsyH6Xro6C7KBDHgMC)}]} CtrtId:vsys.Str(CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(14uNyN6a28tELN5meLbsyCHuaq8V5BX6rDJmU7ATMJaoF73Gfz7) Attachment:vsys.Str()})
```

#### Deposit

Deposit VSYS tokens to a token-holding contract instance(e.g. lock contract).

Note that only the token defined in the token-holding contract instance can be deposited into it.

```go
// sc: *vsys.SysCtrt
// acnt: *vsys.Account
// lc: *vsys.LockCtrt
// amount: float64

lcId := lc.CtrtId.B58Str().Str();

resp, err = sc.Deposit(
  acnt, // by
  lcId, // ctrtId
  10, // amount
  "", // attachment
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(DhyyHBcUtUjv2oK9y1sqgCuca5WemN85eCcCsWkJy68J) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659931919720595000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2hwfFsP5yhCriFpVm6YW7N1FKmmYGXe9ketHSSwYfdFh6NRaS4Q8E4TWFW55u2F5tuFRpj3dao8SGfnNwFp8WFn4)}]} CtrtId:vsys.Str(CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(14VJY1ZthR99KumZMvcjjGnwmMggYiBsTzkczYTMmwdPNFkpULXaUg11ynCnJ7PAqewTixfGjTC6XwkJu1ucT1rf) Attachment:vsys.Str()})*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(DhyyHBcUtUjv2oK9y1sqgCuca5WemN85eCcCsWkJy68J) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659931919720595000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2hwfFsP5yhCriFpVm6YW7N1FKmmYGXe9ketHSSwYfdFh6NRaS4Q8E4TWFW55u2F5tuFRpj3dao8SGfnNwFp8WFn4)}]} CtrtId:vsys.Str(CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(14VJY1ZthR99KumZMvcjjGnwmMggYiBsTzkczYTMmwdPNFkpULXaUg11ynCnJ7PAqewTixfGjTC6XwkJu1ucT1rf) Attachment:vsys.Str()})
```

#### Withdraw

Withdraw VSYS tokens from a token-holding contract instance(e.g. lock contract).

Note that only the one who deposits the token can withdraw.

```go
// acnt0: *vsys.Account
// sc: *vsys.SysCtrt
// lc: *vsys.LockCtrt
// amount: float64

const lcId = lc.ctrtId.data;

lcId := lc.CtrtId.B58Str().Str();

resp, err = sc.Withdraw(
  acnt, // by
  lcId, // ctrtId
  10, // amount
  "", // attachment
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(77DECrd5HAJu5fNNisM3fSG6xK6i5CUHBnDSLkEsrchR) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659931920016001000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(38pBkU2PnXCDLP3rVzAMtARYBqBCK3dho57URFwWrdibvCEgBfSfnjA15gcfhDTEeeUFZFgsoV7JSn5CKXzyLc7e)}]} CtrtId:vsys.Str(CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(14WMYfjwYyNiNDuXJCfYqBWCzreXBLEjt4a8mrixiYxU8qT6ofpWGMuMqQGRmireKtWF2yz9aTqpA8UwUJxWFGfh) Attachment:vsys.Str()})
```

#### Transfer

Transfer the VSYS token to another account(e.g. user or contract).
`transfer` is the underlying action of `send`, `deposit`, and `withdraw`. It is not recommended to use transfer directly. Use `send`, `deposit`, `withdraw` instead when possible.

```go
// acnt0: *vsys.Account
// acnt1: *vsys.Account
// nc: *vsys.SysCtrt
// amount: float64

resp, err = sc.Transfer(
    acnt0, // by
    acnt0.Addr.B58Str().Str(),
    acnt1.Addr.B58Str().Str(), // ctrtId
    10, // amount
    "", // attachment
)
if err != nil {
  log.Fatalln(err)
}
fmt.Println(resp)
```

Example output
```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(CMPfS49owQmWAt2hRwKWK7dC9JRdxfMoiQGex4UuMmhi) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659932081561444000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5jdWvUsehLxoUGA1PK61ZBWaMjUJygg9bxohRG1gXM85cww4WFFfRwUqWiZ8j37cLhf3XLqBkr3ffz7xL3SAqM3F)}]} CtrtId:vsys.Str(CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(14VJY1ZthR99KumZMvcjjGnwmMggYiBsTzkczYTL1fh62FGgJmJfWBJ1yGDjL7nEHzYr6iGDzhVwK4KCQnkkLAij) Attachment:vsys.Str()})
```
