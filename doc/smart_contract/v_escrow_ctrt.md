# V Escrow Contract

- [V Escrow Contract](#v-escrow-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Judge](#judge)
      - [Token id](#token-id)
      - [Duration](#duration)
      - [Judge Duration](#judge-duration)
      - [Contract balance](#contract-balance)
      - [Order payer](#order-payer)
      - [Order recipient](#order-recipient)
      - [Order amount](#order-amount)
      - [Order recipient deposit](#order-recipient-deposit)
      - [Order judge deposit](#order-judge-deposit)
      - [Order fee](#order-fee)
      - [Order recipient amount](#order-recipient-amount)
      - [Order refund](#order-refund)
      - [Order recipient refund](#order-recipient-refund)
      - [Order expiration time](#order-expiration-time)
      - [Order status](#order-status)
      - [Order recipient deposit status](#order-recipient-deposit-status)
      - [Order judge deposit status](#order-judge-deposit-status)
      - [Order submit status](#order-submit-status)
      - [Order judge status](#order-judge-status)
      - [Order recipient locked amount](#order-recipient-locked-amount)
      - [Order judge locked amount](#order-judge-locked-amount)
    - [Actions](#actions)
      - [Supersede](#supersede)
      - [Create](#create)
      - [Recipient_deposit](#recipient_deposit)
      - [Judge deposit](#judge-deposit)
      - [Payer cancel](#payer-cancel)
      - [Recipient cancel](#recipient-cancel)
      - [Judge cancel](#judge-cancel)
      - [Submit work](#submit-work)
      - [Approve work](#approve-work)
      - [Apply to judge](#apply-to-judge)
      - [Do judge](#do-judge)
      - [Submit penalty](#submit-penalty)
      - [Payer refund](#payer-refund)
      - [Recipient refund](#recipient-refund)
      - [Collect](#collect)

## Introduction

[Escrow contract](https://en.wikipedia.org/wiki/Escrow) is a contractual arrangement in which a third party (the stakeholder or escrow agent) receives and disburses money or property for the primary transacting parties, with the disbursement dependent on conditions agreed to by the transacting parties.

The V Escrow contract allows two parties to do transactions with one another if they have mutual trust in a third party. It is expected that the third party will be a large trusted entity that receives fees for facilitating transactions between parties.

## Usage with Go SDK

### Registration

`tokId` is the token id of the token that deposited into this V Escrow contract.

Note that the caller is the judge of the escrow contract.

For testing purpose, you can create a new [token contract]() , then [issue]() some tokens and [deposit]() into the escrow contract.

```go
// acnt: *vsys.Account
// tokId: *vsys.TokenId
// dur: int64
// judgeDur: int64

// Register a new V Escrow contract
vec, err := RegisterVEscrowCtrt(
	acnt,
	string(tokId.B58Str()),
	dur,
	judgeDur,
	"description",
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(vec.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CF8Db346pRd2gxu8wZyXznPtdQpx88hUihd))
```

### From Existing Contract

`vecId` is the escrow contract's id.

```go
// ch: *vsys.Chain
// vecId: string

vecId := "CEzgEwke6qw4im78x22aNgnqKe3dVxfeciD";
vec, err := vsys.NewVEscrowCtrt(
  vecId, // ctrtId
  ch, // chain
)
```

### Querying

#### Maker

The address that made this v escrow contract instance.

```go
// vec: *vsys.VEscrowCtrt

maker, err := vec.Maker()
if err != nil {
    log.Fatal(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Judge

The judge of the contract.

```go
// vec: *vsys.VEscrowCtrt

judge, err := vec.Maker()
if err != nil {
    log.Fatal(err)
}
fmt.Println(judge)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Token id

The token_id of the contract. Caches result if it hasn't been retrieved earlier.

```go
// vec: *vsys.VEscrowCtrt

TokId, err := vec.TokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(TokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Duration

The duration where the recipient can take actions in the contract.

```go
// vec: *vsys.VEscrowCtrt

duration, err := vec.Duration()
if err != nil {
    log.Fatal(err)
}
fmt.Println(duration)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Judge Duration

The duration where the judge can take actions in the contract.

```go
// vec: *vsys.VEscrowCtrt

judgeDuration, err := vec.JudgeDuration()
if err != nil {
log.Fatal(err)
}
fmt.Println(judgeDuration)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Contract balance

The balance of the token within this contract belonging to the user address.

```go
// vec: *vsys.VEscrowCtrt
// acnt: *vsys.Account

bal, err := vec.GetCtrtBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatal(err)
}
fmt.Println(bal)
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1)})
```

#### Order payer

The payer of the order.

```go
// vec: *vsys.VEscrowCtrt
// orderId: string - TransactionID of escrow order

payerAddr, err := vec.GetOrderPayer(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(payerAddr)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Order recipient

The recipient of the order.

```go
// vec: *vsys.VEscrowCtrt
// orderId: string - TransactionID of buy order

rcptAddr, err := vec.GetOrderRecipient(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(rcptAddr)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Order amount

The amount of the order.

```go
// vec: *vsys.VEscrowCtrt
// orderId: string - TransactionID of escrow order

orderAmount, err := vec.GetOrderAmount(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(orderAmount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order recipient deposit

The amount the recipient should deposit in the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

rcptDepositAmount, err := vec.GetOrderRecipientDeposit(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(rcptDepositAmount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order judge deposit

The amount the judge should deposit in the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

judgeDepositAmount, err := vec.GetOrderJudgeDeposit(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(judgeDepositAmount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order fee

The fee of the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

fee, err := vec.GetOrderFee(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(fee)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order recipient amount

The amount the recipient will receive from the order if the order goes smoothly(i.e. work is submitted & approved).

The recipient amount = order amount - order fee.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

rcptAmnt, err := vec.GetOrderRecipientAmount(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(rcptAmnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order refund

The refund amount of the order.

The refund amount means how much the payer will receive if the refund occurs.

It is defined when the order is created.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

refund, err := vec.GetOrderRefund(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(refund)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order recipient refund

The recipient refund amount of the order.

The recipient refund amount means how much the recipient will receive if the refund occurs.

The recipient refund amount = The total deposit(order amount + judge deposit + recipient deposit) - payer refund

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

orderRcptRefund, err := vec.GetOrderRecipientRefund(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(orderRcptRefund)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Order expiration time

The expiration time of the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

expireAt, err := vec.GetOrderExpirationTime(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(expireAt)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Order status

The status of the order. (check if the order is still active)

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

orderStatus, err := vec.GetOrderStatus(orderId)
if err != nil {
	log.Fatal(err)
}
fmt.Println(orderStatus)
```

Example output

```
true
```

#### Order recipient deposit status

The recipient deposit status of the order.

The order recipient deposit status means if the recipient has deposited into the order.

true - recipient deposited

false - recipient didn't deposit

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

orderRcptDepStatus, err := vec.GetOrderRecipientDepositStatus(orderId)
if err != nil {
	log.Fatal(err)
}
fmt.Println(orderRcptDepStatus)
```

Example output

```
true
```

#### Order judge deposit status

The judge deposit status of the order.

The order judge deposit status means if the judge has deposited into the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

orderJudgeDepStatus, err := vec.GetOrderJudgeDepositStatus(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(orderJudgeDepStatus)
```

Example output

```
true
```

#### Order submit status

The submit status of the order.

true - submitted

false - not submitted

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

status, err := vec.GetOrderSubmitStatus(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(status)
```

Example output

```
true
```

#### Order judge status

The judge status of the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

status, err := vec.GetOrderJudgeStatus(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(status)

```

Example output

```
true
```

#### Order recipient locked amount

The amount from the recipient that is locked(deposited) in the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

rcptLockedAmount, err := vec.GetOrderRecipientLockedAmount(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(rcptLockedAmount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(20) Unit:vsys.Unit(1)})
```

#### Order judge locked amount

The amount from the judge that is locked(deposited) in the order.

```go
// vec: *vsys.VEscrow
// orderId: string - TransactionID of escrow order

judgeLockedAmount, err := vec.GetOrderRecipientLockedAmount(orderId)
if err != nil {
    log.Fatal(err)
}
fmt.Println(judgeLockedAmount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

### Actions

#### Supersede

Transfer the judge right of the contract to another account.

```go
// acnt: *vsys.Account
// newJudge: *vsys.Account

resp, err := vec.Supersede(acnt, string(newJudge.Addr.B58Str()), "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(9WcXFScJc9N93kgHvKxwkNGctscYzZ4VSSXqxt3dpuNA) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659687567872034000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3iY2hGSX2okFbunkfMBSwm8TX6DGKGjbnekr2BxnPs8Wkx6zVd17pPJzW4SLuKdk4B1dgac1LxkFP9UCLRjj3GQ9)}]} CtrtId:vsys.Str(CF2wxmcTdSJwcW6ukQK6JmTGyUVz1ZC7PR4) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1bscu1qPwSQ3dpRTmcaVU6cR8yjTQpcJx7S1jy) Attachment:vsys.Str()})
```

#### Create

Create an escrow order and called by the payer.

Note that this transaction id of this action is the **order ID**.

```go
// acnt: *vsys.Account
// recipient: string
// amount: float64
// rcptDepositAmount: float64
// judgeDepositAmount: float64
// orderFee: float64
// refundAmount : float64
// expireTime : Unix timestamp in int64

resp, err := vec.Create(
	payer,
	string(recipient.Addr.B58Str()),
	10, // amount
	2, // rcptDepositAmount
	3, // judgeDepositAmount
	4, // orderFee
	5,  // refundAmount
	expireAt, // expireTime
	"",
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7FkcZ4jyKEw2Qyff3vAj3FWdxPpWpUANwtPPqU2Q4bLG) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659687674576928000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(126zAcrxfSbFDrL8peTr6kbqktu9pzPjxjPLqUSfV8QbCCng8PLu1FHEF52FNGShC68vVmgSz7NtRgZNjkoTJoAy)}]} CtrtId:vsys.Str(CF3WCADjeXNq953mNXVp4aBQJnAfku4qGoB) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(12VHeVmvFYHUtSP17LR8Jj2LZGaVmW9DkQCQgPzraC4LcDoci2NCGp2WGF7QE5iAkdnLhdoaVFfmnwFb9NdjC9amu5dcgrDg21bABggnWEFRyrKdh) Attachment:vsys.Str()})
```

#### Recipient_deposit

Deposit tokens the recipient deposited into the contract into the order.

Note that it is called by the recipient.

```go
// acnt: *vsys.Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.RecipientDeposit(acnt, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FzVfQeQ5eW6Z9CVx18xZgE7U9qZzawt2w94tMmC7nNWW) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688436278735000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(J4kte26zsT8US4KWk1JPnDBhaRTQQyaxUQJQfT4Gtryo) Addr:vsys.Str(AU33yXy9pa3jsvEvQE3mFdxsvVSYmWUi2CZ) Signature:vsys.Str(232mULjdhFY8j4yqF7qV7ZRqDUJh7kRn6QVtoHVQrJzX3S6yfQtFkfAX78rnCzAHsudT4ADP2pqbZoYDB667gvdF)}]} CtrtId:vsys.Str(CEsMXhackNrxR1mcR5gPtDSQjM65YjvymeV) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1TeCHbGgBwL8mKiDWj1wSD7hgnpLJQk8kFqHGb5E5izkbXqkS) Attachment:vsys.Str()})
```

#### Judge deposit

Deposit tokens the judge deposited into the contract into the order.

Note that it is called by the judge.

```go
// acnt: *vsys.Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.JudgeDeposit(acnt, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(4ujbxjc4UXJyG1dD99vQ9fJVCKp3q3VYiP8KwNcLY9a2) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688496240223000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3XxcS7VB3GfkKCfrMTWWZeWwJ6dFu1qJ195atCRGAAJXnTzGm5VyXViqVv3gtiMQuFbDXTcwsSFsa5EVWAtdYD3p)}]} CtrtId:vsys.Str(CFADWhbQr1iJsM1Be8ELh9tJEGCZvGWFpQr) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1TeCHenGAon9PbL5vry5zKM1BoCwfrZzwxbqFyWkVAz5GqtQJ) Attachment:vsys.Str()})
```

#### Payer cancel

Cancel the order by the payer.

Note that it is called by the payer.

```go
// payer: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.PayerCancel(payer, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(4wkdrATvJB19Q8PwzxtL7uLQd4Z6pT4E2nwW7rdonLKc) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688603045262000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(3mLi7FM26Xn53eGtLoNZCkSFa4LEYCv79xZxcyxMEa5b6rxCqCCdxvwnyiWR9qU52A9M2LqmXchN5RJSp2KCW8na)}]} CtrtId:vsys.Str(CEseqCu1op3EpU8122Tmwh8F7c6KPMNNsJZ) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(1TeCHjfeysQptW4PorppSSmRuaKC2oPnXtvMiSgUN9FkLve5v) Attachment:vsys.Str()})
```

#### Recipient cancel

Cancel the order by the recipient.

Note that it is called by the recipient.

```go
// recipient: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.RecipientCancel(recipient, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(3MTLYE9cZBbrnuMgvXrAMdA15eXiSKu6zo9m2To6nyDx) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688668093766000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(J4kte26zsT8US4KWk1JPnDBhaRTQQyaxUQJQfT4Gtryo) Addr:vsys.Str(AU33yXy9pa3jsvEvQE3mFdxsvVSYmWUi2CZ) Signature:vsys.Str(4DUG5jATpTGcrr3aV9HEQvQQsQxA2DojW4PAeQGo34AaN5t74zaSYA6KLAcD6cRupHZk2oR7MctkRNhd3u7U5iav)}]} CtrtId:vsys.Str(CEtcqkTvn3FJQG5Yi4hDcN5u2cmKXQsrZqA) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(1TeCHgUCvZKiXkLmzb4io8rb3AvFCN85fFGRMP4SXD5StCLQP) Attachment:vsys.Str()})
```

#### Judge cancel

Cancel the order by the judge.

Note that it is called by the judge.

```go
// judge: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.JudgeCancel(judge, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(2WQbRvJpWkVJRRHMW2MLcQFVRSMAa5XNAspn4cz9ZeZC) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688751363015000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2SXynyWR1CXdhw4pP7v1fEXPRXazDpc3MmyAz7oZ1QHmmPakwEQDpxCG4pdpzVKpZJ3siWHvauSozyxFnxHPvBSM)}]} CtrtId:vsys.Str(CEtKzoTpWoXoimnpBAmFHn21XDzamiiuzXw) FuncIdx:vsys.FuncIdx(6) FuncData:vsys.Str(1TeCHm6RCwBM6XeYTCEWrSwcAS6aZsN9GwLV6dLyLjH69DdAg) Attachment:vsys.Str()})
```

#### Submit work

Submit the work by the recipient.

Note that it is called by the recipient.

```go
// recipient: *vsys.Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.SubmitWork(recipient, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(GsRwjyaFcQXNFgt1rWVeE7ZBkPHGrR63BB9wHRMZdVS8) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688864627222000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(J4kte26zsT8US4KWk1JPnDBhaRTQQyaxUQJQfT4Gtryo) Addr:vsys.Str(AU33yXy9pa3jsvEvQE3mFdxsvVSYmWUi2CZ) Signature:vsys.Str(3cNpxUCbUjHrrQKRHopJsbQUunsReDihQP7FoGwckoH75WL9GoQhbrso4Dr5yQ7xzqpawcdFv5MKDcFqfwZYwipw)}]} CtrtId:vsys.Str(CF8KZV5iD7YsDEztSDnNrkpc9uZU5Zg9wJp) FuncIdx:vsys.FuncIdx(7) FuncData:vsys.Str(1TeCHdACMmiF4BRbVijkXNQi37gGVDuTPfbpHp8NDxevF4Fh6) Attachment:vsys.Str()})
```

#### Approve work

Approve the work and agrees the amounts are paid by the payer.

Note that it is called by the payer.

```go
// payer: *vsys.Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.ApproveWork(payer, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(G2bp7q9a4XTL9qtYN5kCyu9WX1VZnTbw6uhLYL3TRBoN) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688935036079000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(yj1YPHVH1qrXYB536etpqrvAK9CuV7kaRhag7e2W3rp4wxWhUT83aZ6weGBtrG7U3w7h2N7xSJWu3wpsLuSczFL)}]} CtrtId:vsys.Str(CF6Gbn7b15uv91QDZMv9KpnLntq2HGxGDEb) FuncIdx:vsys.FuncIdx(8) FuncData:vsys.Str(1TeCHkVKmr8zQzM9u8cDm9CgACEfgvprGQwB4kkyY6GjoEuR5) Attachment:vsys.Str()})
```

#### Apply to judge

Apply for the help from judge by the payer.

Note that it is called by the payer.

```go
// payer: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.ApplyToJudge(payer, orderId, "")
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(GuyQ9w3etdxZM2QTzKAT9CiqwUyqXC61nJJA4nvWTfPY) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659688999270122000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(2sNksPm3ZaqmgRpTtcKpSHNpnjyaGCNX3cn5V4SF4xTKrVpyEtkCci2Yr1dgHjnZnrd9CatG41XfRKSjRE5xNRy4)}]} CtrtId:vsys.Str(CF3EvXc5PXkfvBFBC1WdDsucoQqt9HVUMbc) FuncIdx:vsys.FuncIdx(9) FuncData:vsys.Str(1TeCHcRHp1E6yMpCr4UNQunQbmyXRF3E31KG1cgGfRL5Y7ooz) Attachment:vsys.Str()})
```

#### Do judge

Judge the work and decides on how much the payer & recipient will receive.

Note that it is called by the judge.

```go
// judge: Account
// orderId: string - transaction ID of order(create) transaction

resp, err = vec.DoJudge(
	judge, // by
	orderId, // orderId
	45, // toPayer
	2, // toJudge
	"", // attachment
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(HTRjFhgDegryyr1QY3aaeRTGgnzZs1Z3KzyQp6qfrrWz) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659689104915714000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3CKFA5dPJQuUghAh8ZEWrK7V2rdGXaHkjYwZGy61jWhMRATRBWKqzAS2u5C7sj5sz5LfUP2kFnxfitCakm1GkXCn)}]} CtrtId:vsys.Str(CF39zaovG4cSfrE7xhs8Jur2vQZbsgS58as) FuncIdx:vsys.FuncIdx(10) FuncData:vsys.Str(1FELDyKMuTXwf5CDfRvtFBhwhfXQzmY1mYDUVYZtEzeZsddFjuZ4xgn2cbnbDVZiDtrxrcD4fE) Attachment:vsys.Str()})
```

#### Submit penalty

Submit penalty by the payer for the case where the recipient does not submit work before the expiration time. The payer will obtain the recipient deposit amount and the payer amount(fee deducted).

The judge will still get the fee.

Note that it is called by the payer.

```go
// payer: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.SubmitPenalty(payer, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(5X8r6ga2Wofs36YJwcrbX6HateoTRrC5FbbNpYddhc9G) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659689165127824000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(pTe8uXXAeHjxZ9MHcPg5Ltwgxh4CmHzPaRwtpNc8etrYFrCEivHxHrf75zvArwvXkAPzAD2a2PagPMqNKoBeVZr)}]} CtrtId:vsys.Str(CF8h8C5iTic4ViygP5UbyaqtwKxdRqK73gt) FuncIdx:vsys.FuncIdx(11) FuncData:vsys.Str(1TeCHcQ6CCBsCZBSCfNyudHZSVC9Tx9j7DrijXfsrp4EejZPe) Attachment:vsys.Str()})
```

#### Payer refund

Make the refund action by the payer when the judge does not judge the work in time after the apply_to_judge function is invoked.

The judge loses his deposit amount and the payer receives the refund amount.

The recipient receives the rest.

Note that it is called by the payer.

```go
// payer: Account
// orderId: string - transaction ID of order(create) transaction

resp, err = vec.PayerRefund(payer, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(8pFfm4iqpobzLiss8xvPo7kNp6YiCjKZhUqMgcoNLxJv) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659689267327989000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(67KRZbCRU8FcL9W8sjuyRCx3s6tVzZwV5i6q63jdDEfqwcNXRfSrabbLeKmBz144ZTi8Ko1oEAWNyDQTAvTj2QLp)}]} CtrtId:vsys.Str(CFFbveUDEB73fat9sypxGn7FDk1j7UQDyfz) FuncIdx:vsys.FuncIdx(12) FuncData:vsys.Str(1TeCHd1EyagzNAYQY7ejschcB86dqnRcrCoxUroR72uNe9cEc) Attachment:vsys.Str()})
```

#### Recipient refund

Make the refund action by the recipient when the judge does not judge the work in time after the apply_to_judge function is invoked.

The judge loses his deposit amount and the payer receives the refund amount.

The recipient receives the rest.

Note that it is called by the recipient.

```go
// recipient: Account
// orderId: string - transaction ID of order(create) transaction

resp, err = vec.RecipientRefund(recipient, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(HAoPFbdPhkF5F4huiqDMyPL5Pqyk4jfKsr6qt8Ah7YPj) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659689370455060000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(J4kte26zsT8US4KWk1JPnDBhaRTQQyaxUQJQfT4Gtryo) Addr:vsys.Str(AU33yXy9pa3jsvEvQE3mFdxsvVSYmWUi2CZ) Signature:vsys.Str(4ahCj9Fe1VLPWdUwyfUrqb2yhkRV6NsBVX67GQfVWa4zfpsL8HWzpnji2tehqGfN212gJNJJ45trHTUVoyWppRfi)}]} CtrtId:vsys.Str(CF2tDgr8Fq7E5ekMh1MbiDk7nesnFp34S9y) FuncIdx:vsys.FuncIdx(13) FuncData:vsys.Str(1TeCHmHpy6m4V1qX8PvvqjZRKLZsbZZsTccE59TB6BiGXkHrJ) Attachment:vsys.Str()})
```

#### Collect

Collect the order amount & recipient deposited amount by the recipient when the work is submitted while the payer doesn't either approve or apply to judge in his action duration.

The judge will get judge deposited amount & fee.

Note that it is called by the recipient.

```go
// recipient: Account
// orderId: string - transaction ID of order(create) transaction

resp, err := vec.Collect(recipient, orderId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FNUGubnuEQH99xCmrfYeE5m9QkktkZiCRS3knHdaLrch) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659689469615799000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(J4kte26zsT8US4KWk1JPnDBhaRTQQyaxUQJQfT4Gtryo) Addr:vsys.Str(AU33yXy9pa3jsvEvQE3mFdxsvVSYmWUi2CZ) Signature:vsys.Str(4gKCFgASBWkgZZkuFbJH3VgsDkkVvZNgMecJ2ECYUJLws8fbcvGnsyt1TujVGAtroLxnf971rrb8Ls1jepfc94S7)}]} CtrtId:vsys.Str(CF19EUsitWcZ6ZZkgub7Ng93d9bZ3eaVPMd) FuncIdx:vsys.FuncIdx(14) FuncData:vsys.Str(1TeCHmCSeD3rS3fr25z6B3T3DTuzQeH5EVWELNc42x4x9zzTE) Attachment:vsys.Str()})
```
