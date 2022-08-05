# Token Contract V1 With Split

- [Token Contract V1 With Split](#token-contract-v1-with-split)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Issuer](#issuer)
      - [Maker](#maker)
      - [Token ID](#token-id)
      - [Unit](#unit)
      - [Token Balance](#token-balance)
    - [Actions](#actions)
      - [Supersede](#supersede)
      - [Issue](#issue)
      - [Send](#send)
      - [Destroy](#destroy)
      - [Transfer](#transfer)
      - [Deposit](#deposit)
      - [Withdraw](#withdraw)
      - [Split](#split)

## Introduction

_Token Contract V1 with Split_ is the twin case for _[Token Contract V1 Without Split](./tok_ctrt_no_split.md)_.
The token unit can be updated at any time after the contract instance is registered.

! Usage of Token Contract V1 with Split is the same as Token Contract V1 Without Split, but former has additional `Split` method.
For other methods please refer to _[Token Contract V1 Without Split](./tok_ctrt_no_split.md)_ and note that it has subtle difference of different function indices and `CtrtMeta`.

## Usage with Go SDK

### Registration

The example below shows registering an instance of Token Contract V1 With Split where the max amount is 100 & the unit is 100.

```go
// acnt: *vsys.Account

tc, err := RegisterTokCtrtWithSplit(
    acnt, // by
    1000, // max
    1, // unit
    "", // tokDescription
    "", // ctrtDescription
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(tc.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CEu9aFoVwdApYBAPFy4hTYc2NUJRzoL5VYc))
```

### From Existing Contract

`tokCtrtId` is the ctrtID of previously registered token.

```go
// ch: *vsys.Chain

tokCtrtId := "CEu9aFoVwdApYBAPFy4hTYc2NUJRzoL5VYc";
tc, err := NewTokCtrtWithoutSplit(tokCtrtId, ch)
if err != nil {
    log.Fatalln(err)
}
```

### Actions

#### Split

Update the unit of the token.

The address with the issuer & maker role can take this action.

```go
// acnt: *vsys.Account
// tc: *vsys.TokCtrtWithSplit

resp, err := tc.Split(acnt, 12, "")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(GZXkp7DWNyaraTLWzawTyQQkvUjq7ayRnzvx3swrZhoL) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659682684649092000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4dP7C4RkgKs5MrGUh7eL9Yhqve31b43xxvpRUDDZ7oBJHMmYFYqRL3pbFqwfQumw1FUUHfrDN8aesV3NiCczuofV)}]} CtrtId:vsys.Str(CF8TNioMm3xTto2qaaRJopvtgJjJdZWTrrs) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(14JDCrdo1xwstP) Attachment:vsys.Str()})
```
