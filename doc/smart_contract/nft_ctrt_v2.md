# NFT Contract V2

- [NFT Contract V2](#nft-contract-v2)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Issuer](#issuer)
      - [Maker](#maker)
      - [Regulator](#regulator)
      - [Unit](#unit)
      - [Is user in the list](#is-user-in-the-list)
      - [Is contract in the list](#is-contract-in-the-list)
    - [Actions](#actions)
      - [Issue](#issue)
      - [Send](#send)
      - [Transfer](#transfer)
      - [Deposit](#deposit)
      - [Withdraw](#withdraw)
      - [Supersede](#supersede)
      - [Add/remove a user from the list](#addremove-a-user-from-the-list)
      - [Add/remove a contract from the list](#addremove-a-contract-from-the-list)

## Introduction

NFT contract V2 adds additional whitelist/blacklist regulation feature upon [NFT contract V1](./nft_ctrt.md).

For the whitelist flavor, only users & contracts included in the list can interact with the NFT contract instance.

For the blacklist flavor, only users & contracts excluded from the list can interact with the NFT contract instance.

## Usage with Go SDK

Examples of NFT contract V2 with whitelist are shown below. The usage of the blacklist one is very similar.

### Registration

```go
// acnt: *vsys.Account

// Register a new NFT contract
nc, err := RegisterNFTCtrtV2Whitelist(acnt, "ctrtDescription")
if err != nil {
   log.Fatal(err)
}
fmt.Println(nc.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CF5xGQwkj1bkiRFuRzF53LmM9Wk6ghDU8jk))
```

### From Existing Contract

```go
// acnt: *vsys.Account
// ch: *vsys.Chain

ncId := "CF46ab6o5HTfyLwgyBhwrhkEmLxbaLkSJ8a";
nc, err = vsys.NewNFTCtrtV2Whitelist(ncId, ch);
```

### Querying

#### Issuer

The address that has the issuing right of the NFT contract instance.

```go
// nc: *vsys.NFTCtrtV2Whitelist

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

The address that made this NFT contract instance.

```go
// nc: *vsys.NFTCtrtV2Whitelist

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

#### Regulator

The address that serves as the regulator of the NFT contract instance.

```go
// nc: *vsys.NFTCtrtV2Whitelist

regulator, err := nc.Regulator()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(regulator)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Unit

The unit of tokens defined in this NFT contract instance.

As the unit is obviously fixed to 1 for NFTs, the support of querying unit of NFT is for the compatibility with other token-defining contracts.

```go
// nc: *vsys.NFTCtrtV2Whitelist
unit, err := nc.Unit()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Is user in the list

Check if the user is in the whitelist/blacklist

```go
// acnt: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

inList, err := nc.IsUserInList(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(inList)
```

Example output

```
false
```

#### Is contract in the list

Check if the user is in the whitelist/blacklist

```go
// nc: *vsys.NFTCtrtV2Whitelist
arbitraryCtrtId = 'CEsGmTPZMvPkkG7g5gyqgRcXRVc2ZcVXz9J';

inList, err := nc.IsCtrtInList(arbitraryCtrtId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(inList)
```

Example output

```
false
```

### Actions

#### Issue

Define a new NFT and issue it. Only the issuer of the contract instance can take this action. The issued NFT will belong to the issuer.

```go
// acnt: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err := nc.Issue(acnt, "description", "attachment")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7aSmkT4b1Akwqb7xs966VEcyDq3B6UhD39xemQDPBJVH) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659667325515528000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(22SmncziLCKq5ZrqewA6LJU2sPARnmYVuJyD8uJ9ESKnP25DWP7rAYz8y7woaU541V1HcSAL1AuUWvPJ2dSRXRLA)}]} CtrtId:vsys.Str(CEwLraxTgd5kqq6xujwJVdvh9LepnCtgtGT) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(12exFLNpw2q1rqUCQ1TSD) Attachment:vsys.Str(6UZYuvjBHC18dZ)})
```

#### Send

Send an NFT to another user.

```go
// acnt0: *vsys.Account
// acnt1: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err := nc.Send(
    acnt0, // by
    string(acnt1.Addr.B58Str()), // receiver
    0, // tokIdx
    "sending nft", // attachment
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(2jzWXrenDrGkTzq3r7ccBBXWaaYBjF2695bYAKU3Ha7F) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659667459188403000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4gfVxdpM3Av9NxdMGJA787iSPuaxvK5RE6nkNbhVeMZ5AhsAXqitSL84rFZDzYdh8Nvygobe3bc4Tznn4diwyTYR)}]} CtrtId:vsys.Str(CEu6VZgiQLi3MKkZYcveM2vPD7Li5ZUYCPW) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1bbXGbm97k4F4a3sXETMi6qjJQttC27hakqt4pMBr9zT9) Attachment:vsys.Str(Vch6McStcZsuB3R)})
```

#### Transfer

Transfer the ownership of an NFT to another account(e.g. user or contract).
`transfer` is the underlying action of `send`, `deposit`, and `withdraw`. It is not recommended to use transfer directly. Use `send`, `deposit`, `withdraw` instead when possible.

```go
// acnt0: *vsys.Account
// acnt1: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err := nc.Transfer(
	acnt0, // by
	string(acnt0.Addr.B58Str()), // from
	string(acnt1.Addr.B58Str()), // to
	0, // tokIdx
	"transfering nft", // attachment
)
if err != nil {
   log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(ChWSqrYfa86dGYDhknHEb2Qn2qyHVRiaPtegyoNSpSq1) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659667870536415000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2tt1GrMe9PmHVgS68zBhXTFGtsthaHX4ZVznzQfWFctVQispm12xXyFLonWkNtJEfmqmw6ZfnpEMVP1uTtJ8EUsP)}]} CtrtId:vsys.Str(CFBQN4YxGsietdAFQcp698bnFSYMVwLMbyP) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1Xv7sJ3F8TibBYnc4tpnpwzQhsnMMwd3rT7ceEY8s9pLBtJJXtX1bpTyKvBZ4WM2goXMSRxzE2QwtSZEfZ) Attachment:vsys.Str(Vch6McStcZsuB3R)})
```

#### Deposit

Deposit an NFT to a token-holding contract instance(e.g. lock contract).

Note that only the token defined in the token-holding contract instance can be deposited into it.

```go
// by: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist
// lc: *vsys.LockCtrt or any other Ctrt

resp, err := nc.Deposit(
    by, //by
    string(lc.CtrtId.B58Str()), // ctrtId
    0, // tokIdx
    "", // ctrtId
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(BLB3dkTWiQNqawTY5gpnE1iscfQCt8oVnLWVU17J666T) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659667984406840000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4bAX9i5FHY7bdiSwFhXretkj4Fcgs4311KQ91xACWB2oQdwhHKErMXA9ASvM5RJzYhR5717zz8saAeKkPr7pdbte)}]} CtrtId:vsys.Str(CF2YVRextJW5n6L3jNt1MZBAqnjS2qz8et5) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(1Xv7sJ3F8TibBYnc4tpnpwzQhsnMMwd3rT7ceEYQVfp7V6KDAKWrhpa1fchK2kok75vkxaaQ4s2Tkte7wM) Attachment:vsys.Str()})
```

#### Withdraw

Withdraw an NFT from a token-holding contract instance(e.g. lock contract).

Note that only the one who deposits the token can withdraw.

```go
// by: *vsysAccount
// lc: *vsys.LockCtrt
// nc: *vsys.NFTCtrt

resp, err := nc.Withdraw(
    by, //by
    string(lc.CtrtId.B58Str()), // ctrtId
    0, // tokIdx
    "", // ctrtId
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7mfZdHnw18rF4jQxEjqSZB8uE9Btn2Z2vjWZExi8R89k) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659668133916825000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3AV9NCNcbFJKvNjaDE2u5JmuhK6829xN9hPDSLpqzC7gvE7GtgHp7tr1i89YeEHj1a5RU6dGruDdK4jesRwW7UpS)}]} CtrtId:vsys.Str(CEwyYxh2bvEqPcpbbYWuUXA5anc64hxtbe1) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(1Y5SeNkP6dmx6pzBaKDpb9RpzrNuxBxqkApPGf8QvxDUwYDX1z93SaPcAYsPy1rUWj2aXvtZfMaQZx9Jp7) Attachment:vsys.Str()})
```

#### Supersede

Transfer the issuer & regulator role of the contract to a new user.

Note that only the contract maker has the privilege to take this action.

```go
// by: *vsys.Account
// newIssuer: *vsys.Account
// newRegulator: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err := nc.Supersede(
	by,
	newIssuer.Addr.B58Str().Str(),
	newRegulator.Addr.B58Str().Str(),
	"")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(EakdyUdoxZ4uCZZRY12pULNUnCoR1PCabBmPG3UkKsSM) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659669688440816000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(DyMYGrm3zgJVPvJRWb59GuAHBudahou8oBhFitkKNpN2CpGyZp2wYPLpHbtp1XaQA3VeLTGaFm1beFBcAQ3cDtR)}]} CtrtId:vsys.Str(CF6bKccG5zAcmwE6ZPmWn3FN23AtR7dBEq7) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1iSiatNyb1DDH9BaTyxhnTrAyermNcUxPmyevchWCMt14uT4NN8T6rL24sBegdGdpgmRhZVE1du) Attachment:vsys.Str()})
```

#### Add/remove a user from the list

Add/remove a user from the whitelist/blacklist.

Note the regulator has the privilege to take this action.

```go
// by: *vsys.Account
// newUser: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err := nc.UpdateListUser(
	by,
	newUser.Addr.B58Str().Str(),
	true,
	"",
)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(4TGHNyGWU6vDmQzj163yCWXiTDQYjYvydHxzUxTmoDiL) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659669824735864000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(39dUeak5jH6GcutYVv79L72pKW5iPpLh4bF3hGEW5Thu186YEgbkwvRhQr2AvEKAztwqqFQcY6W2e8StCc85MuCt)}]} CtrtId:vsys.Str(CFDeuUNXceiMsMarEv6RgayrKMMhR5s8E4K) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1QLRyQ2bdPeXxwK1NwyjTTKNjvUy4Ft5CsmXqwBZi) Attachment:vsys.Str()})
```

#### Add/remove a contract from the list

Add/remove a contract from the whitelist/blacklist.

Note the regulator has the privilege to take this action.

```go
// acnt0: *vsys.Account
// acnt1: *vsys.Account
// nc: *vsys.NFTCtrtV2Whitelist

resp, err = nc.UpdateListCtrt(
	by, // by
	ctrtId, // ctrtId
	false, // value
	"", // attachment
)
if err != nil {
	t.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(6bZJvj2wtNxtCWLJBMRgNd8XPfjzAoNaQcdYPDR5ob1N) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659669964098182000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4nKD2eyMzgxbARr3ipY4Y9fyhGns7nyQCTneTFHqbP7waPNWEai94JmFqpNxJd15ge5guhvBKSEmUcG67bfrFGLc)}]} CtrtId:vsys.Str(CFBVhETZHJt83qCV2edGDa8EMaeoJnGkcpd) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1QWyS19bRamf9yjJbDQDrENYE2gyzKqxSchVRsW87) Attachment:vsys.Str()})
```
