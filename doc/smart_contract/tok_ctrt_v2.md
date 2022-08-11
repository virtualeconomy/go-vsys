# Token Contract V2 Without Split

- [Token Contract V2 Without Split](#token-contract-v2-without-split)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [User membership](#user-membership)
      - [Contract membership](#contract-membership)
      - [Others](#others)
    - [Actions](#actions)
      - [Add/remove a user from the list](#addremove-a-user-from-the-list)
      - [Add/remove a contract from the list](#addremove-a-contract-from-the-list)
      - [Supersede](#supersede)
      - [Others](#others-1)

## Introduction

_Token Contract V2 Without Split_ adds additional whitelist/blacklist regulation feature upon _[Token Contract V1 Without Split](./tok_ctrt_no_split.md)_

For the whitelist flavor, only users & contracts included in the list can interact with the contract instance.

For the blacklist flavor, only users & contracts excluded from the list can interact with the contract instance.

## Usage with Go SDK

### Registration

The example below shows registering an instance of Token Contract V2 Without Split with whitelist where the max amount is 100 & the unit is 100.

The usage of the blacklist one is very similiar.

```go
// acnt: *vsys.Account

// can use TokCtrtV2Blacklist instead of Whitelist
tc, err := vsys.RegisterTokCtrtWithoutSplitV2Whitelist(
	acnt, // by
	1000, // max
	1, // unit
	"", // tokDescription
	"", // ctrtDescription
)
if err != nil {
    log.Fatal(err)
}
fmt.Printlnt(tc.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CFAs41T54TeSe2hSf87bc67xAJNmGZabYNq))
```

### From Existing Contract

```go
// ch: *vsys.Chain

tcId := "CEu9aFoVwdApYBAPFy4hTYc2NUJRzoL5VYc" 
tc, err := vsys.NewTokCtrtV2Whitelist(tcId, ch) // ctrtId, chain
```

### Querying

#### User membership

Get the status of whether the user address is in the list. Returns boolean value.

```go
// acnt: *vsys.Account
// tc: *vsys.TokCtrtV2Whitelist

inList, err := tc.IsUserInList(newUser.Addr.B58Str().Str())
if err != nil {
	log.Fatal(err)
}
fmt.Println(inList)
```

Example ouput

```
true
```

#### Contract membership

Get the status of whether the contract id is in the list. Returns boolean value

```go
// tc: *vsys.TokCtrtV2Whitelist

// CtrtId we are interested in
ctrtId := "CF5Zkj2Ycx72WrBnjrcNHvJRVwsbNX1tjgT";

inList, err := tc.IsCtrtInList(ctrtId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(inList)
```

Example ouput

```
false
```

#### Others

Usage of other queries are same as of _[Token Contract V1 Without Split](./tok_ctrt_no_split.md)_. However, function indices are different.

* Issuer
* Maker
* Token ID
* Unit
* Token Balance

### Actions

#### Add/remove a user from the list

Add/remove a user from the whitelist/blacklist.

Note the regulator has the privilege to take this action.

```go
// by: *vsys.Account
// newUser: *vsys.Account
// tc: TokCtrtV2Whitelist

resp, err = tc.UpdateListUser(
    by,
    newUser.Addr.B58Str().Str(),
    true, // val
    "attachment",
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FxzUaLMvLrtKZVKXB9aVu6k5e2CoY8Vf1vK938vx6n1D) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659941617534065000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2VCvQMW6VUnd2zvXp8wDAC3qeNUDrFi9Wx4qojwLytfmbGEx4frZjYJDkQgSv9nRSJTKAbTTG55MYvA3YKxZdDBR)}]} CtrtId:vsys.Str(CFDrmGhxW7dLbKXvK9awRJBuXJK1eVoZCpz) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1QLRyTYc1KT9fwmwFJn79wKxhUDUvAmYGwK5P35Wo) Attachment:vsys.Str()})
```

#### Add/remove a contract from the list

Add/remove a contract from the whitelist/blacklist.

Note the regulator has the privilege to take this action.

```go
// acnt0: *vsys.Account
// tc: TokCtrtV2Whitelist

// arbitrary CtrtId
ctrtId := "CF5Zkj2Ycx72WrBnjrcNHvJRVwsbNX1tjgT";

resp, err := tc.UpdateListCtrt(
    by,
    ctrtId,
    true, // val, false to remove ctrt from list
    "",
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(8p2mT325Vj19ZHEqDfRWCzPQE5cdDnTw3eqXTXhyaUYJ) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659941670269542000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(ophhKnqqybuUrkD8dCviU1ozSW1tDd2HWoQnAxDmxFgF2RLJBbXYERYtcwYu9tpKgW6ERSKvgX3wBGS8W31vJQb)}]} CtrtId:vsys.Str(CFEvMGPfr86NHd5DCo5Y34pGRqf1mvhjj5d) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1QWyS4okkmKgCYjSsQSwojNHLzK5S94HhrHkRnGTS) Attachment:vsys.Str()})
```

#### Supersede

Transfer the issuer role of the contract to a new user.

The maker of the contract has the privilege to take this action.

```go
// by: *vsys.Account
// newIssuer: *vsys.Account
// newRegulator: *vsys.Account
// tc: TokCtrtV2Whitelist

resp, err := tc.Supersede(by, newIssuer.Addr.B58Str().Str(), newRegulator.Addr.B58Str().Str(), "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(B9Em9FpdrYGU1uHgYC8SHtBvschFmt4cXtwoQYesQ3HC) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659941893704269000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2Bx2fzvAD52TMr3cmfw9X2M7GcnQoFUyKTVSyGZi9BfZ9Si5E6q13resrBXz1MJtFdJ5XTkpbEXzdk7Vcxcja2cF)}]} CtrtId:vsys.Str(CEwa19GzL1DDKTLZVcXWHefMhAvHwi9WkBe) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1iSiatNyb1DDH9BaTyxhnTrAyermNcUxPmyevchWCMt14uT4NN8T6rL24sBegdGdpgmRhZVE1du) Attachment:vsys.Str()})
```

#### Others

Usage of other actions is the same as of _[Token Contract V1 Without Split](./tok_ctrt_no_split.md)_. However, function indices are different.

* Issue
* Send
* Destroy
* Transfer
* Deposit
* Withdraw

