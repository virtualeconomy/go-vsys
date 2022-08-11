# Payment Channel Contract

- [Payment Channel Contract](#payment-channel-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Token id](#token-id)
      - [Contract balance](#contract-balance)
      - [Channel creator](#channel-creator)
      - [Channel creator's public key](#channel-creators-public-key)
      - [Channel recipient](#channel-recipient)
      - [Channel accumulated load](#channel-accumulated-load)
      - [Channel accumulated payment](#channel-accumulated-payment)
      - [Channel expiration time](#channel-expiration-time)
      - [Channel status](#channel-status)
    - [Actions](#actions)
      - [Create and load](#create-and-load)
      - [Extend expiration time](#extend-expiration-time)
      - [Load](#load)
      - [Abort](#abort)
      - [Unload](#unload)
      - [Collect payment](#collect-payment)
      - [Generate the signature for off chain payments](#generate-the-signature-for-off-chain-payments)
      - [Verify the signature](#verify-the-signature)

## Introduction

Payment Channels have been implemented in a large number of blockchains as a method to increase the scalability of any protocol. By taking a large number of the transactions between two parties off-chain, it can significantly reduce the time and money cost of transactions.

The payment channel contract in VSYS is a one-way payment channel, which means that the paying user is considered as `sender` and the receiving user is `receiver`.

## Usage with Go SDK

### Registration

`tokId` is the token id of the token that deposited into this payment channel contract.

For testing purpose, you can create a new [token contract]() , then [issue]() some tokens and [deposit]() into the payment channel contract.

```go
// acnt: *vsys.Account
// tokId: string

// Register a new Payment Channel contract
pc, err := RegisterPayChanCtrt(acnt, tokId, "")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(pc.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CEu3em2sWen1ZHCk5NjgVfeNKechamkbpC5))
```

### From Existing Contract

ncId is the payment channel contract's id.

```go
// ch: *vsys.Chain
// pcId: string

pcId := "CEu3em2sWen1ZHCk5NjgVfeNKechamkbpC5";
pc, err := NewPayChanCtrt(ncId, ch);
```

### Querying

#### Maker

The address that made this payment channel contract instance.

```go
// pc: *vsys.PayChanCtrt

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

#### Token id

The token id of the token that deposited into this payment channel contract.

```go
// pc: *vsys.PayChanCtrt

pcTokId, err := pc.TokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(pcTokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsYHw1uLWJvoikwk3q6Qsk76H2PorauMuqw3LthP))
```

#### Contract balance

The token balance within this contract.

```go
// pc: *vsys.PayChanCtrt
// acnt: *vsys.Account

ctrtBal, err := pc.GetCtrtBal(acnt.Addr.B58Str().Str())
if err != nil {
t.Fatal(err)
}
fmt.Println(ctrtBal)
fmt.Println(ctrtBal.Amount())
```

Example output

```
*vsys.Token({Data:vsys.Amount(100) Unit:vsys.Unit(1)})
100
```

#### Channel creator

The channel creator.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanCreator, err := pc.GetChanCreator(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanCreator)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Channel creator's public key

The channel creator's public key.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanCreatorPubKey, err := pc.GetChanCreatorPubKey(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanCreatorPubKey)
```

Example output

```
*vsys.PubKey(vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv))
```

#### Channel recipient

The recipient of the channel.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanReicpient, err := pc.GetChanRecipient(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanRecipient)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Channel accumulated load

The accumulated load of the channel.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanAccumLoad, err := pc.GetChanAccumLoad(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanAccumLoad)
fmt.Println(chanAccumLoad.Amount())
```

Example output

```
*vsys.Token({Data:vsys.Amount(50) Unit:vsys.Unit(1)})
50
```

#### Channel accumulated payment

The accumulated payment of the channel.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanAccumPay, err := pc.GetChanAccumPay(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanAccumPay)
fmt.Println(chanAccumPay.Amount())
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1)})
0
```

#### Channel expiration time

The expiration time of the channel.

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanExpTime, err := pc.GetChanExpTime(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanExpTime)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Channel status

The channel status (check if the channel is still alive).

```go
// pc: *vsys.PayChanCtrt
// chanId: string e.g. '5dv575QktQMfB9YEi1qyzm5yi9YMMApGNZyddTbsxmpK'
// chanId is the transaction id of the createAndLoad function.

chanStatus, err := pc.GetChanStatus(chanId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(chanStatus)
```

Example output

```
true
```

### Actions

#### Create and load

Create the payment channel and loads an amount into it.

Note that this transaction id is the channel id.

```go
// acnt: *vsys.Account
// recipient: string
// amount: float64
// expiredTime: int64

later := time.Now().Unix() + 60*10 // 10 min from now

resp, err := pc.CreateAndLoad(acnt, recipient, 20.0, later, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(6jwDfUDM6bn5mc7YP3RGuwpF5GHFy7terr4VSLNBf8cg) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677318424957000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(RWcvbZppqZVtb82ouoZbFwWcfDaou6hfc5HFqAfbQfeXmRv8EKtWzgsZFpBamvy2Xzz7QSKUEYxeD4GugiyFcPG)}]} CtrtId:vsys.Str(CF1oxS1GAQFhqhgjwMT64uXz2s1Fu6wyGL7) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1L43p2L9xewtiVCLZN2SFvGTsZkVCUufacstcipWnSs46Hg9NSXEKeQDca6Y4ET) Attachment:vsys.Str()})
```

#### Extend expiration time

Extend the expiration time of the channel to the new input timestamp.

```go
// acnt: *vsys.Account
// chanId: string
// chanExpTimeOld: vsys.VSYSTimestamp

newLater := chanExpTimeOld.UnixTs() + 300
resp, err := pc.ExtendExpTime(acnt, chanId, newLater, "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(59BVN2VeQFf4fh88mwKaWctSwC8KG42cXeeGXxHfKast) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677478578767000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(hV3JqVRSCRqdQx1SA8EdLEaExoyt27AEdFnBctmZ1GeZi5DazaL1z2b5myDANEctNqTw2NPnivwTNHrJFa7LAyR)}]} CtrtId:vsys.Str(CF1wvdQtam16TNZcFr4THhFKMqfcRnG77yP) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(13w3j8AaJFkJpqCnJzHAdbxByEAU3gVXdTkHNY6hAdMNUBsVXHGXEAVQg7TKNf) Attachment:vsys.Str()})
```

#### Load

Load more tokens into the channel.

```go
// acnt: *vsys.Account
// chanId: string
// amount: float64

resp, err := pc.Load(acnt, chanId, amount, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7YHjshxYbQREo8yVH11uSBHm2Lnh82hzKwdFaSY28MB3) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677571069648000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(2pnx5ez4Txk8xcWxBRjd1M2Ar6dMF74KpCpAmsH7EbKvS7XP6af8wKx8qaeUKkLRhBiEzBVy6bLo1t9Trxp9BRgX)}]} CtrtId:vsys.Str(CErZqy3ES5Q9Ag3Yy1bUBiMnqU7MD311YDg) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(13w3j884hdPQCymuFdP6QtAN7EUmuudw671um2L1k5a2yntPcgTT2mvCUHekrC) Attachment:vsys.Str()})
```

#### Abort

Abort the channel, triggering a 2-day grace period where the recipient can still collect payments. After 2 days, the payer can unload all the remaining funds that was locked in the channel.

```go
// acnt: *vsys.Account
// chanId: string

resp, err := pc.Abort(acnt, chanId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(5S1fuZJG6F8C4mF3kgp5rBi5HCqoCJn28PidY6HxKRcv) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677658719821000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(SSPeFokQg5HV3DS726XyogzeToxMcRC5QVYpnsANYRxEYeKS52e2LHf9m96Sb1LFFDbvskWhy4g6AQdTsHZPweP)}]} CtrtId:vsys.Str(CF3d2iG8nQ8czYEMSdf3Ad1d8GyP2GjKjwc) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(1TeCHZaoTMukZfufRc87vo4XjKz19hChpFGc4oUUQUKoxZn2S) Attachment:vsys.Str()})
```

#### Unload

Unload all the funds locked in the channel (only works if the channel has expired).

```go
// acnt: Account
// chanId: string

resp, err = pc.Unload(acnt, chanId, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(GsZnUqU5VUQC3AJNHxjxCtkP1atjCjZKNFBXPYcmRksJ) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677747101430000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(ME1qpKxnu26YGP26njFNsuAmSi5cNUVXCGYidEsvtfFuskUCBf9nQZkX7YCJAHMJosWq8N44g39exL32rxVNsuN)}]} CtrtId:vsys.Str(CEzyMNMXMe1JhaNNM9gT3zufZuUX2SHZKZw) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(1TeCHox5ojbApnrqouBou8V2tyBSsUQsam5QJkkUfc1QFNxYz) Attachment:vsys.Str()})
```

#### Collect payment

Collect the payment from the channel (only works if the channel has expired).

```go
// acnt: *vsys.Account
// chanId: string
// amount: float64
// signature: string

resp, err := pc.CollectPayment(
    acnt,
    chanId,
    amount,
    signature,
    "attachment",
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(B1ttRq8cumDcRMnKvN2rZ3ZpPZhnsZPr1a37FwL2WiAh) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659677861838636000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(4zC56SQpKZp6rGaSBp76i1HYSAAUWuMEmEx8kRFxYYprgeDnDnb2uUXTJi3Gcva3M5N4GDR2PKeTcUPCzfLTR7sW)}]} CtrtId:vsys.Str(CFFUsuqRh13yZxPahuvKJmEoyjpekWXMNDi) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(1a8vFZapK2fiib6xTEkjkQa9ewS3f4uRuMsN1CcKZEzdLj3XECwYobAxRY3PNwDD1VDqv77vYzUgYPLcetus8gptkp9WRGsLZTeFdKBiw6gbvLL1VQ3EAEBPsTzgoBRPrRu9N82rq2dFQvSEkTNWsFsZC) Attachment:vsys.Str()})
```

#### Generate the signature for off chain payments

Generate the offchain payment signature.

```go
// acnt: *vsys.Account
// chanId: string
// amount: float64

sig, err := pc.OffchainPay(
    acnt.PriKey,
    chanId,
    amount,
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(sig)
```

Example output

```
H94c7xnYvB1qkoB1LqRERgacpAuiiP1CeL6yiwnS5gN4fntS3NCz4iYuSMgp7SgU6Fxn3NXN3kHd1tCT8hupEXA
```

#### Verify the signature

Verify the payment signature.

```go
// chanId: string
// amount: float64
// signature: string

ok, err := p.VerifySig(chanId, amount, signature)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(ok)
```

Example output

```
true
```
