# Lock Contract

- [Lock Contract](#lock-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Token id](#token-id)
      - [Contract balance](#contract-balance)
      - [Lock time](#lock-time)
    - [Actions](#actions)
      - [Lock](#lock)

## Introduction

Lock contract allows users to lock a specific token in the contract for some period of time. This allows users to guarantee they have a certain amount of funds upon lock expiration. This may be helpful in implementing some kinds of staking interactions with users of a VSYS token for instance.

## Usage with Go SDK

### Registration

`tokId` is the token id of the token that deposited into this lock contract.

For testing purpose, you can create a new [token contract]() , then [issue]() some tokens and [deposit]() into the lock contract.

```go
// acnt: *vsys.Account
// tokId: string - like "TWtUothN6spjSiy1amsWWSLX1uyhTELfv4juaHaSL"

// Register a new Lock contract
lc, err := vsys.RegisterLockCtrt(acnt, tokId, "")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(lc.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CF14Edtwb2f6E2ywKtaXzJAVQ6gecjBYaJY))
```

### From Existing Contract

```go
// ch: *vsys.Chain

lc, err := vsys.NewLockCtrt("CFFGTjUwuM41Dk7iVaJg88BrEsPmwQKTmM6", ch)
if err != nil {
    log.Fatalln(err)
}
```

### Querying

#### Maker

The address that made this lock contract instance.

```go
// lc: *vsys.LockCtrt

maker, err := lc.Maker()
if err != nil {
    t.Fatal(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Token id

The token id of the token that deposited into this lock contract.

```go
// lc: *vsys.LockCtrt

tokId, err := lc.TokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWtPBJN8Sa7i5ZHpMBEeyCE8dCgoAxktMfQRqi1rT))
```

#### Contract balance

The token balance within this contract.

Note that the balance is the same no matter the token is locked or not.

```go
// lc: *vsys.LockCtrt
// acnt: *vsys.Account

bal, err := lc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(bal)
fmt.Println(bal.Amount()) // to see results better
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1000000)}
0
```

#### Lock time

The expire timestamp.

```go
// nc: jv.LockCtrt
// acnt: *vsys.Account

lockTime, err := lc.GetCtrtLockTime(testAcnt0.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(lockTime)
```

Example output

```
vsys.VSYSTimestamp(1646984339000000000)
```

### Actions

#### Lock

Lock the token until the expire time. The token can't be withdrawn before the expire time.

```go
// acnt: *vsys.Account
// expireTime: number

later := time.Now().Unix() + 3*int64(avgBlockDelay.Seconds())
resp, err = lc.Lock(testAcnt0, later, "")
if err != nil {
	    t.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(5ymGrJmHDQwVPrfK95hfFomfkKivWh5JQgPvVc35iBCG) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659666103947325000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5KS9MKyr2hWte1vPnMsjrSrZwG6qKMD3Z3NZX5Fm7DPu4GTK1KHawTo1baqEf66shotRjwdkA8GU2eY7MvDAefCM)}]} CtrtId:vsys.Str(CF199YnfvuaqBi58gFRzUD4Btb7Y2GRmxzP) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(14NhyRawcMhvPq) Attachment:vsys.Str()})
```
