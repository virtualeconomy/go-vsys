# Token Contract V1 Without Split

- [Token Contract V1 Without Split](#token-contract-v1-without-split)
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

## Introduction

Token contract supports defining & managing custom tokens on VSYS blockchain.

A token is a logical entity on the blockchain. It can represent basically everything that can be stored in a database. Be it a fiat currency like USD, financial assets like a share in a company, or even reputation points of an online platform.

A contract can be thought of as a class in OOP with a bunch of methods. An instance needs to be created before using a contract. Every node will maintain the states for a contract instance. The user interacts with the contract instance by publishing transactions. Upon receiving transactions, every node will update the contract instance states accordingly.

_Token Contract V1 Without Split_ is the classic version of token contracts supported by VSYS blockchain. The token unit cannot be updated once specified when registering the contract instance.

> “Unit” is the granularity of splitting a token. It can be thought of as the smallest denomination available. Let’s take real-world money as an example, if the unit is set to 100, it means the smallest denomination is a cent, and 100 cents is a dollar.
>
> With “Unit”, float numbers can be represented in integers so as to avoid the uncertainty comes from float computation. If we set unit == 100, 1.5 tokens are actually stored as 150 on the blockchain.

## Usage with Go SDK

### Registration

The example below shows registering an instance of Token Contract V1 Without Split where the max amount is 10000 & the unit is 100. Fee set as a default value.

```go
// acnt: *vsys.Account

tc, err := RegisterTokCtrtWithoutSplit(
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
*vsys.CtrtId(vsys.Str(CF4MC3UvPvL8jeRW8aqzHVHaQ1c6FTfFoZQ))
```

### From Existing Contract

`tokCtrtId` is the ctrtID of previously registered token.

```go
// ch: *vsys.Chain

tokCtrtId := "CF4MC3UvPvL8jeRW8aqzHVHaQ1c6FTfFoZQ";
tc, err := NewTokCtrtWithoutSplit(tokCtrtId, ch)
if err != nil {
	log.Fatalln(err)
}
```

### Querying

#### Issuer

The address that has the issuing right of the Token contract instance.

```go
// tc: *vsys.TokCtrtWithoutSplit

issuer, err := nc.Issuer()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(issuer)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Maker

The address that made this Token contract instance.

```go
// tc: *vsys.TokCtrtWithoutSplit

maker, err := nc.Maker()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Token ID

The token ID of the token defined in the token contract instance.

Note that theoretically a token contract instance can have multiple kinds of token, it is restricted to 1 kind of token per token contract instance. In other word, the token ID is of the token index `0`.

```go
// tc: *vsys.TokCtrtWithoutSplit
tokId, err := tc.TokId()
if err != nil {
	log.Fatalln(tokId)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWtPBJN8Sa7i5ZHpMBEeyCE8dCgoAxktMfQRqi1rT))
```

#### Unit

The unit of the token defined in this token contract instance.

```go
// tc: *vsys.TokCtrtWithoutSplit

unit, err := tc.Unit()
if err != nil {
    log.Fatal(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(100)
```

#### Token Balance

Query the balance of the token defined in the contract for the given user.

```go
// tc: *vsys.TokCtrtWithoutSplit
// acnt: *vsys.Account

bal, err := tc.GetTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatal(err)
}
fmt.Println(bal)
```

Example output

```
*vsys.Token({Data:vsys.Amount(100) Unit:vsys.Unit(1)})
```

### Actions

#### Supersede

Transfer the issuer role of the contract to a new user.
The maker of the contract has the privilege to take this action.

```go
// by: *vsys.Account
// newIssuer: *vsys.Account
// tc: *vsys.TokCtrtWithoutSplit

resp, err := tc.Supersede(by, newIssuer.Addr.B58Str().Str(), "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(ESYPs25FTcMeviGoZvXYGXJT4FkNUKQSPeosFD3VH9nf) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679297410049000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5UagGFthYmwdQ3jT6rJsfMGgnHCxFw3VjXos64USjdrtHbLCsZSaNk55qoRgSKuK5bB2vH6CMy7DPDdUqnwDMNd6)}]} CtrtId:vsys.Str(CEuEZRW4bLDBtQcrh73g2P9RVE7q9cLR7Du) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1bscu1qPwSQ3dpRTmcaVU6cR8yjTQpcJx7S1jy) Attachment:vsys.Str()})
```

#### Issue

Issue the a certain amount of the token. The issued tokens will belong to the issuer.

Note that only the address with the issuer role can take this action.

```go
// acnt: *vsys.Account
// tc: *vsys.TokCtrtWithoutSplit

resp, err := tc.Issue(by, 100, "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

feeScale: 100,

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(J6yh6tMTbsGFpt9DqTLpHW85vRYRuxtcauFyVjp9QVqQ) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659680117946222000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(iRWrSQzwej9FsFzyGaUNGPcfN2mMij2BjKTuamPQftpUfpy6HDSL7QKRxrd7wGbEsMT1yNTkX8YM8qEvpyHnNSU)}]} CtrtId:vsys.Str(CFEHr5mAEheq89XS3HqXQbVFuNYsQa5Pp3w) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(14JDCrdo1xwsuu) Attachment:vsys.Str(6UZYuvjBHC18dZ)})
```

#### Send

Send a certain amount of the token to another user.

```go
// sender: *vsys.Account
// receiver: *vsys.Account
// tc: *vsys.TokCtrtWithoutSplit

resp, err := tc.Send(sender, string(receiver.Addr.B58Str()), 100, "sending")
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(Er5FVULJAyA4CY73VyStxHHqj8fSctgHYMGWKZoDS2xD) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679399479253000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2gB9farKTqgfTpZHgHYLuFsL7GNKWUBqD8VDriPmqrdiX5Y1FQa1F48SPhvWohqvJ3YV5TFjaCwXjVj2PSzkGSMJ)}]} CtrtId:vsys.Str(CFA1hre9y6aShZ7TSiVEsUVxnsaAuJfZdfd) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(14uNyN6a28tELN5meLbsyCHuaq8V5BX6rDJmU7ATMJaoF5WuQrX) Attachment:vsys.Str(Vch6McStcZsuB3R)})
```

#### Destroy

Destroy a certain amount of the token.

Note that only the address with the issuer role can take this action.

```go
// acnt: *vsys.Accout
// tc: *vsys.TokCtrtWithoutSplit

resp, err := tc.Destroy(acnt, 100, "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(Er5FVULJAyA4CY73VyStxHHqj8fSctgHYMGWKZoDS2xD) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679399479253000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2gB9farKTqgfTpZHgHYLuFsL7GNKWUBqD8VDriPmqrdiX5Y1FQa1F48SPhvWohqvJ3YV5TFjaCwXjVj2PSzkGSMJ)}]} CtrtId:vsys.Str(CFA1hre9y6aShZ7TSiVEsUVxnsaAuJfZdfd) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(14uNyN6a28tELN5meLbsyCHuaq8V5BX6rDJmU7ATMJaoF5WuQrX) Attachment:vsys.Str(Vch6McStcZsuB3R)})
```

#### Transfer

Transfer a certain amount of the token to another account(e.g. user or contract).
`transfer` is the underlying action of `send`, `deposit`, and `withdraw`. It is not recommended to use transfer directly. Use `send`, `deposit`, `withdraw` instead when possible.

```go
// sender: *vsys.Account
// receiver: *vsys.Account
// tc: TokCtrtWithoutSplit

resp, err := tc.Transfer(
	sender, // by 
	string(sender.Addr.B58Str()), // sender
	string(receiver.Addr.B58Str()), // receiver
	100, // amount
	"sending", // attachment
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(AdphQNuPGJgu7UGqsE9NKmgU4tQ5YbDtA6uF8RerYCgV) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679653424343000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2a1uaRNcmiFkEFzkSCgEvx1BbcL6i1BGDbCKk6UR1sJsAGJTfH152rdw1YxaJczRMcCRHHom97a3PoKTqPmua8ou)}]} CtrtId:vsys.Str(CFE6QXayssQKt9Fttk2joyuBb29mwoNzsQs) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(14VJY1ZthR99KumZMvcjjGnwmMggYiBsTzkczYTL1fh62FGgJmJfWBJ1yGDjL7nEHzYr6iGDzhVwK4KCQnjDxub9) Attachment:vsys.Str(5Ndmf6AYog)})
```

#### Deposit

Deposit a certain amount of the token into a token-holding contract instance(e.g. lock contract).

Note that only the token defined in the token-holding contract instance can be deposited into it.

```go
// by: *vsys.Account
// tc: *vsys.TokCtrtWithoutSplit
// lc: *vsys.LockCtrt - can be any contract

resp, err := tc.Deposit(by, lc.CtrtId.B58Str().Str(), 5.0, "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(HvpRzC8zHJJ9AKccaWDQwJZkwE2CR8eXLhsr96zMAKHP) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679869321365000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4mWfhXeu5GKAHk77oBMoFtpQ362HmTqPZb6bELo9YfqPJy4fjGwGXeoVV5peJtz6Y4CeAvRmJjezhgPBHpxBdkW7)}]} CtrtId:vsys.Str(CF9Z91havpNwD6x4sDdpL3gXxYYXi9UY2GB) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(14VJY1ZthR99KumZMvcjjGnwmMggYiBsTzkczYTMmwdFZwMcAPapy3YZynUh3y8q8xtJAUCm1SGNyWL58RwHQn4k) Attachment:vsys.Str()})
```

#### Withdraw

Withdraw a certain amount of the token from a token-holding contract instance(e.g. lock contract).

Note that only the token defined in the token-holding contract instance can be withdrawn from it.

```go
// by: Accout
// tc: TokCtrtWithoutSplit
// lc: LockCtrt

resp, err = tc.Withdraw(by, lc.CtrtId.B58Str().Str(), 5.0, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(s9BuZbs8gBiPUAWY9DE8WzsgXELUSWRPhKcRTrdFdjs) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659679948582979000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(66PGzK5dywwBERiJmkt8TZDxZ3UhZRXhokWH4N3gfZH19RyShmHBvbbsyNzSg2JJU1zEsJHxEDuWVRuym5M2nkQq)}]} CtrtId:vsys.Str(CFBiB2ZPpxNMxqiEY6rjEz5jRgBJ9MBzmuz) FuncIdx:vsys.FuncIdx(6) FuncData:vsys.Str(14WMYfo8tcmRKkYrvmukX5QGJZaaKxkSvhAMAWNfWXHFiGvBAFPzDT4hUHtv6esRx7rnVBumR2V4YpVQvGDhwpWg) Attachment:vsys.Str()})
```
