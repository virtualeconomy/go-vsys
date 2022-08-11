# V Stable Swap Contract

- [V Stable Swap Contract](#v-stable-swap-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Base Token ID](#base-token-id)
      - [Target Token ID](#target-token-id)
      - [Base Token Contract](#base-token-contract)
      - [Target Token Contract](#target-token-contract)
      - [Base Token Unit](#base-token-unit)
      - [Target Token Unit](#target-token-unit)
      - [Max Order Limit Per User](#max-order-limit-per-user)
      - [Unit of Price of Base Token](#unit-of-price-of-base-token)
      - [Unit of Price of Target Token](#unit-of-price-of-target-token)
      - [Base Token Balance](#base-token-balance)
      - [Target Token Balance](#target-token-balance)
      - [User Orders](#user-orders)
      - [Order Owner](#order-owner)
      - [Base Token Fee](#base-token-fee)
      - [Target Token Fee](#target-token-fee)
      - [Base Token Minimum Trading Amount](#base-token-minimum-trading-amount)
      - [Base Token Maximum Trading Amount](#base-token-maximum-trading-amount)
      - [Target Token Minimum Trading Amount](#target-token-minimum-trading-amount)
      - [Target Token Maximum Trading Amount](#target-token-maximum-trading-amount)
      - [Base Token Price](#base-token-price)
      - [Target Token Price](#target-token-price)
      - [Base Token Locked Amount](#base-token-locked-amount)
      - [Target Token Locked Amount](#target-token-locked-amount)
      - [Order Status](#order-status)
    - [Actions](#actions)
      - [Supersede](#supersede)
      - [Set Order](#set-order)
      - [Update Order](#update-order)
      - [Deposit to Order](#deposit-to-order)
      - [Withdraw from Order](#withdraw-from-order)
      - [Close Order](#close-order)
      - [Swap Base Tokens to Target Tokens](#swap-base-tokens-to-target-tokens)
      - [Swap Target Tokens to Base Tokens](#swap-target-tokens-to-base-tokens)

## Introduction

The V Stable Swap contract supports creating a liquidity pool of 2 kinds of tokens for exchange on VSYS. The fee is fixed.

The order created in the contract acts like a liquidity pool for two kinds of tokens(i.e. the base token & the target token). Traders are free to take either side of the trade(i.e. base to target or target to base).

The V Stable Swap contract can accept any type of token in the VSYS blockchain, including option tokens created through the V Option Contract.

## Usage with Go SDK

### Registration

Register a contract instance.

```go
// acnt: *vsys.Account
// baseTokId: string E.g. "TWssXmoLvyB3ssAaJiKk5d7ambFHBxcmr9sMRtPLa"
// targetTokId: string E.g. "TWtoBbmn5UgQd9KgtbWkBY96hiUJWzeTTggGrb8ba"

vss, err := RegisterVStableSwapCtrt(by, baseTokId, targetTokId, 5, 1, 1, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(vss.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CF1LQZ5U2S1WiXHbVdY8CwKjhqC1kF8GZwt))
```

### From Existing Contract

Get an object for an existing contract instance.

```go
// ch: *vsys.Chain

vssId := "CF1LQZ5U2S1WiXHbVdY8CwKjhqC1kF8GZwt"
vss, err := NewVStableSwapCtrt(vssId, ch);
```

### Querying

#### Maker

The address that made this contract instance.

```go
// vss: *vsys.VStableSwapCtrt

maker, err := vss.Maker()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Base Token ID

The token ID of the base token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

tokId, err := vss.BaseTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Target Token ID

The token ID of the target token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

tokId, err := vss.TargetTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Base Token Contract

The token contract object of the base token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

tc, err := vss.BaseTokCtrt()
if err != nil {
	log.Fatal(err)
}
fmt.Println(tc)
fmt.Printf("%T", tc)
```

Example output

```
*vsys.Ctrt({CtrtId:*vsys.CtrtId(vsys.Str(CF3f1AQd3C46RBnxJiAxDY5NByzTWZstNNU)) Chain:*vsys.Chain({NodeAPI:*vsys.NodeAPI(http://veldidina.vos.systems:9928) ChainID:vsys.ChainID(T)})})
*vsys.TokCtrtWithoutSplit
```

#### Target Token Contract

The token contract object of the target token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

tc, err := vss.BaseTokCtrt()
if err != nil {
	log.Fatal(err)
}
fmt.Println(tc)
fmt.Printf("%T", tc)
```

Example output

```
*vsys.Ctrt({CtrtId:*vsys.CtrtId(vsys.Str(CF3f1AQd3C46RBnxJiAxDY5NByzTWZstNNU)) Chain:*vsys.Chain({NodeAPI:*vsys.NodeAPI(http://veldidina.vos.systems:9928) ChainID:vsys.ChainID(T)})})
*vsys.TokCtrtWithoutSplit
```

#### Base Token Unit

The unit of the base token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

unit, err := vss.BaseTokUnit()
if err != nil {
    log.Fatal(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Target Token Unit

The unit of the target token in the contract instance.

```go
// vss: *vsys.VStableSwapCtrt

unit, err := vss.TargetTokUnit()
if err != nil {
    log.Fatal(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Max Order Limit Per User

The maximum number of orders each user can create.

```go
// vss: *vsys.VStableSwapCtrt

num, err := vss.MaxOrderPerUser()
if err != nil {
	log.Fatal(err)
}
fmt.Println(num)
```

Example output

```
5
```

#### Unit of Price of Base Token

The unit of price of base token(i.e. how many target tokens are required to get one base token).

```go
// vss: *vsys.VStableSwapCtrt

unit, err := vss.BasePriceUnit()
if err != nil {
    log.Fatal(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Unit of Price of Target Token

The unit of price of target token(i.e. how many base tokens are required to get one target token).

```go
// vss: *vsys.VStableSwapCtrt

unit, err := vss.TargetPriceUnit()
if err != nil {
    log.Fatal(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Base Token Balance

Get the base token balance of the given user.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account

bal, err := vss.GetBaseTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatal(err)
}
fmt.Println(bal)
```

Example output

```
*vsys.Token({Data:vsys.Amount(100) Unit:vsys.Unit(100)})
```

#### Target Token Balance

Get the target token balance of the given user.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account

bal, err := vss.GetTargetTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatal(err)
}
fmt.Println(bal)
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(100)})
```

#### User Orders

Get the number of orders the user has made.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account

orders, err := vss.GetUserOrders(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatal(err)
}
fmt.Println(orders)
```

Example output

```
0
```

#### Order Owner

Get the owner of the order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

owner, err := vss.GetOrderOwner(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(owner)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Base Token Fee

Get the fee for trading base token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

fee, err := vss.GetFeeBase(order_id)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(fee)
```

Example output

```
*vsys.Token({Data:vsys.Amount(100) Unit:vsys.Unit(100)})
```

#### Target Token Fee

Get the fee for trading target token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

fee, err := vss.GetFeeTarget(order_id)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(fee)
```

Example output

```
*vsys.Token({Data:vsys.Amount(100) Unit:vsys.Unit(100)})
```

#### Base Token Minimum Trading Amount

Get the minimum trading amount for base token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetMinBase(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(1) Unit:vsys.Unit(100)})
```

#### Base Token Maximum Trading Amount

Get the maximum trading amount for base token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetMaxBase(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(1) Unit:vsys.Unit(100)})
```

#### Target Token Minimum Trading Amount

Get the minimum trading amount for target token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetMinTarget(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(1) Unit:vsys.Unit(100)})
```

#### Target Token Maximum Trading Amount

Get the maximum trading amount for target token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetMaxTarget(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(1) Unit:vsys.Unit(100)})
```

#### Base Token Price

Get the price for base token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

price, err := vss.GetPriceBase(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(price)
```

Example output

```
*vsys.Token({Data:vsys.Amount(2) Unit:vsys.Unit(1)})
```

#### Target Token Price

Get the price for target token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

price, err := vss.GetPriceBase(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(price)
```

Example output

```
*vsys.Token({Data:vsys.Amount(2) Unit:vsys.Unit(1)})
```

#### Base Token Locked Amount

Get the locked amount of base token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetBaseTokLocked(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(500) Unit:vsys.Unit(1)})
```

#### Target Token Locked Amount

Get the locked amount of target token in the given order.

```go
// vss: *vsys.VStableSwapCtrt
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

amnt, err := vss.GetTargetTokLocked(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(500) Unit:vsys.Unit(1)})
```

#### Order Status

Get the status of the given order(i.e. if the order is active).

```go
// vss: *vsys.VStableSwapCtrt
// ordeId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

status, err := vss.GetTargetTokLocked(orderId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(status)
```

Example output

```
true
```

### Actions

#### Supersede

Transfer the contract right to another account.

Only the maker of the contract has the right to take this action.

```go
// vss: *vsys.VStableSwapCtrt
// by: *vsys.Account
// newIssuer: *vsys.Account

resp, err := vss.Supersede(by, string(newIssuer.Addr.B58Str()), "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FRb38stt5pywh16asMACMnpFb71QjcWXeLPDqjncWdXH) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660030825058242000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2eAAFG2WCdS7pPdYvmTQe8T29MDHp3uiUuV33oDQDy6j5aT41bBz89veqLRKNSaLa2egUM6vvpaz6zp2Eaz18fAu)}]} CtrtId:vsys.Str(CEuFn7XPv55bCb1EV73X6SiGU93pAM8yP9P) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1bscu1qPwSQ3dpRTmcaVU6cR8yjTQpcJx7S1jy) Attachment:vsys.Str()})
```

#### Set Order

Create an order and deposit initial amounts into the order.

The transaction ID returned by this action serves as the ordeId.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account

resp, err := vss.SetOrder(
	acnt, // by
	1, // feeBase
	1, // feeTarget
	0, // minBase
	100, // maxBase
	0, // minTarget
	100, // MaxTarget
	1, // priceBase
	1, // priceTarget
	500, // baseDeposit
	500, // targetDeposit
	"", // attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(pedAwwDVQqBLaFvkDy3WCRpAUmAgCh2RbnpXJyjhjqh) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031036099708000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5miSNu1hDsM5JQLXrZA2wYmSwmLpkPmccWjUQvZJpbtkotczAJHyeaTb8RhwNi6h9hUVL6UT6q2gdzck59jrnaEU)}]} CtrtId:vsys.Str(CEtVaVhkmUqNNSk9GH1Lypq4ZX8FqNvmvsC) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(17vgyw5jxgmT6KmwVEdUesvR7JtSoqjauiyXntaieHKQV4mouqmj1SFAJWQyDK6p5csTsPqvEGWfRKF4DLPPstBQrJmq8gT9kNV6GMEsv1mj4ajHALhbdTFdagDTy) Attachment:vsys.Str()})
```

#### Update Order

Update the order settings(e.g. fee, price)

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

resp, err = vss.UpdateOrder(acnt,
	orderId, // by
	1, // feeBase
	1, // feeTarget
	0, // minBase
	100, // maxBase
	0, // minTarget
	100, // maxTarget
	1, // priceBase
	1, // priceTarget
	"", // attachment
)
if err != nil {
t.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(2dtecUqWjA24nerTrb5nPEaGJKwVEoRm5XDoagG2DbQz) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031143749583000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2xxnhTKRFxxjjBFV867pQoYmsLZ1s6ay1TS4QThHz1T7bL85sJzqYSMwxpcuuVLYuF6mSEWqJiAV8Pxmy9VTzPuZ)}]} CtrtId:vsys.Str(CEsTnHZRyBFmgX3eaM6Z2FnfNYhazj3nwgd) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1G3pzCLWeWNwNzkiZfnK4D2uyfkLKHT4mdh1hmmCoyCWs6zAJMYhVcrBwTkVr6bPqp3vwhvhnUz2qtjHsa75w99BADDNwmpZXYCYiZjR3xPUxNMeB2Rrqw29HL8b6tqThmTKTLjMvzZU3VEAkoPe) Attachment:vsys.Str()})
```

#### Deposit to Order

Deposit more tokens into the order.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account
// orderId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

resp, err := vss.OrderDeposit(
	acnt, // by
	orderId, // orderId
	200, // baseDeposit
	100, // targetDeposit
	"", // attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(6gEhyBj6tZkVR4dKaYH9hVG8DUBjeqZrCx2scpaNbmmM) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031215117193000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(vcwdfKnHBkUH39tnkgkoL4GPFYDWkZh7E9j6gFZUjPcaMwqZux4L7VMtssrn2wFWaieoZz7DUUmVoh6Go6fqaZX)}]} CtrtId:vsys.Str(CF49x4yBfB6T77fuSVwTo4hiTHtPa1urp3F) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1FELDwJmxQhELp8Y17b8TTDL9cDvvRLjLDESmQj3GFVKcRBRyuLiqguAxZ5hp5eAaosVTDXGtB) Attachment:vsys.Str()})
```

#### Withdraw from Order

Withdraw some tokens from the order.

```go
// vss: *vsys.VStableSwapCtrt
// acnt: *vsys.Account
// ordeId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"


resp, err := vss.OrderWithdraw(
    acnt, // by
    orderId, // orderId
    200, // baseWithdraw
    100, // targetWithdraw
    "", // attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(4FvPvT37VyDJbc6M2tKtiiqwWgcR9VVLpVPqVWY1UyRL) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031223385588000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5GB74NGcNwJhQMcE7VRdHJjCE2qRDHGyy2iFjesTLJfKqaKwb2yTyxn7znaDew2MDF6MCF75NzET9Peq6Kt7YuCc)}]} CtrtId:vsys.Str(CF49x4yBfB6T77fuSVwTo4hiTHtPa1urp3F) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(1FELDwJmxQhELp8Y17b8TTDL9cDvvRLjLDESmQj3GFVKcRBRyuLiqguAxZ5hp5eAaosVTDXGtB) Attachment:vsys.Str()})
```

#### Close Order

Close the given order.

```go
// vss: *vsys.VStableSwapCtrt
// acnt0: *vsys.Account
// ordeId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

resp, err := vss.CloseOrder(acnt, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7ujZbpxsq8fUHM2BS8jdNxsbwLFydd94B2TBqoD1tuwR) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031379336085000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2kekFomW8iai7j7gJ51CvoWaye7pFQ4G3B5vMuiPXzf2V4GQDivoutCzKrVjq6XoQ2rieduiNQ24ZmcwS5WkDsN9)}]} CtrtId:vsys.Str(CEshXnnRQg5vjKDxPWQq9bdTtLTCi91UEHo) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(1TeCHbDb8sPu8TdNrki9zwvZfN8iviQ7qGEY5FqVK1YnuRnw9) Attachment:vsys.Str()})
```

#### Swap Base Tokens to Target Tokens

Trade base tokens for the target tokens.

```go
// vss: *vsys.VStableSwapCtrt
// acnt1: *vsys.Account
// ordeId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

deadline := time.Now().Unix() + 1500
resp, err := vss.SwapBaseToTarget(
	acnt, // by
	orderId, // orderId
	10, // amount
	1, //swapFee
	1, // price
	deadline, //deadline
	"", // attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(4Dy5TmaGghAZkHgXaRYSYb9cfTHm9MEVqaJcZRpfWEh3) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031440673839000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(56uHQu64dmsR7zivTanczHwF242jJ6GzmBYCyDessqgUivuqdJDRk6qUC8zCJ5XijXzqajPU3ypfHRyRWrzFiCd)}]} CtrtId:vsys.Str(CF25VvgQM4BdJ3CYZymWYw6rYpidED9ma9u) FuncIdx:vsys.FuncIdx(6) FuncData:vsys.Str(15KQH1YZCeNURVzK52Y7fsmQNuck4TFPEEWCj5jHc9fJkchEWyciy32GgCm7fuV8GqnBQDCJ1odnvJnGikUQviyDFvXzPoXfsBd) Attachment:vsys.Str()})
```

#### Swap Target Tokens to Base Tokens

Trade target tokens for the base tokens.

```go
// vss: *vsys.VStableSwapCtrt
// acnt1: *vsys.Account
// ordeId: string E.g. "JChwB1yFyFMUjSLCruuTDHVPWHWqvYvQBkFkinnmRmvY"

deadline := time.Now().Unix() + 1500
resp, err = vss.SwapTargetToBase(
	acnt, // by
	orderId, // orderId
	10, // amount
	1, // swapFee
	1, // price
	deadline, //deadline
	"", //attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FTPi94J2ha2V3M6vv7deQz6cf5HuTVAYRqLL5ZHqdGGf) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660031449068525000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(643C5ietcHoSAayCFbvZJytLxj8KEfF4czuP8Bu9k4C8AU5RSfVC77471mNxD1vk6ERQch7SKAL6JxiqGszpyZn8)}]} CtrtId:vsys.Str(CF25VvgQM4BdJ3CYZymWYw6rYpidED9ma9u) FuncIdx:vsys.FuncIdx(7) FuncData:vsys.Str(15KQH1YZCeNURVzK52Y7fsmQNuck4TFPEEWCj5jHc9fJkchEWyciy32GgCm7fuV8GqnBQDCJ1odnvJnGikUQviyDFvXzPoXfsBd) Attachment:vsys.Str()})
```
