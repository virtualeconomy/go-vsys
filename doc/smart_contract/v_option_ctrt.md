# V Option Contract

- [V Option Contract](#v-option-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Base token ID](#base-token-id)
      - [Target token ID](#target-token-id)
      - [Option token ID](#option-token-id)
      - [Proof token ID](#proof-token-id)
      - [Execute time](#execute-time)
      - [Execute deadline](#execute-deadline)
      - [Option status](#option-status)
      - [Max issue num](#max-issue-num)
      - [Reserved option](#reserved-option)
      - [Reserved proof](#reserved-proof)
      - [Price](#price)
      - [Price unit](#price-unit)
      - [Token locked](#token-locked)
      - [Token collected](#token-collected)
      - [Base token balance](#base-token-balance)
      - [Target token balance](#target-token-balance)
      - [Option token balance](#option-token-balance)
      - [Proof token balance](#proof-token-balance)
    - [Actions](#actions)
      - [Supersede](#supersede)
      - [Activate](#activate)
      - [Mint](#mint)
      - [Unlock](#unlock)
      - [Execute](#execute)
      - [Collect](#collect)

## Introduction

[Option contract](https://en.wikipedia.org/wiki/Option_contract) is defined as "a promise which meets the requirements for the formation of a contract and limits the promisor's power to revoke an offer".

Option Contract in VSYS provides an opportunity for the interested parties to buy or sell a VSYS underlying asset based on the determined agreement(e.g., pre-determined price, execute timestamp and so on). It allows users to create option tokens on the VSYS blockchain, and buyers holding these option tokens have the right to buy or sell some underlying asset at some point in the future. Different from the traditional option market, everyone can buy or sell option tokens to join the option market at any time without any contractual relationship with an exchange.

## Usage with Go SDK

### Registration

Register a V Option Contract instance.

```go
// acnt: Account
// baseTokId: string
// targetTokId: string
// optionTokId: string
// proofTokId: string
// executeTime: int64 - Unix timestamp
// executeDeadline: int64 - Unix timestamp

// Register a new contract instance
vo, err := RegisterVOptionCtrt(
    testAcnt0, // by
    "TWuTFmDynQz815ygjgmoiwU1BL3UTA7VC5VFVQLdw", // baseTokId
    "TWsbVKNno1EHdeWfjJpkQkCqgPcba2zk8Ud4BU5dV", // targetTokId
    "TWu14YBZrSKfK3bMpNpoLx4jWjFKPuyeWYdnXNqHt", // optionTokId
    "TWttyxGagVTPWjx2zJ5Wqb67RuvY3B6gXpkkJ39cA", // proofTokId
    time.Now().Unix()+20, // executeTime
    time.Now().Unix()+40, // executeDeadline
    "attachment",
)
fmt.Println(vo.CtrtId);
```

Example output

```
*vsys.CtrtId(vsys.Str(CEtD4tFXLdqrXAUqCXaiYBFUrEUB4GkC91G))
```

### From Existing Contract

Get an object for an existing contract instance.

```go
// ch: *vsys.Chain
// voc_id: string - contract Id of registered V Option contract

voc_id = "CEtD4tFXLdqrXAUqCXaiYBFUrEUB4GkC91G";
voc, err := NewVOptionCtrt(
  voc_id, // ctrtId
  ch, // chain
)
```

### Querying

#### Maker

The address that made this contract instance.

```go
// voc: *vsys.VOptionCtrt

maker, err := voc.Maker()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Base token ID

The base token ID.

```go
// voc: VOptionCtrt

tokId, err := voc.BaseTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Target token ID

The target token ID.

```go
// voc: VOptionCtrt

tokId, err := vec.TargetTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Option token ID

The option token ID.

```go
// voc: VOptionCtrt

tokId, err := vec.OptionTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Proof token ID

The proof token ID.

```go
// voc: VOptionCtrt

tokId, err := vec.ProofTokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Execute time

The execute time.

```go
// voc: VOptionCtrt

time, err := vec.ExecuteTime()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(time)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Execute deadline

The execute deadline.

```go
// voc: VOptionCtrt

time, err := vec.ExecuteDeadline()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(time)
```

Example output

```
vsys.VSYSTimestamp(1659673299000000000)
```

#### Option status

The option contract's status.(check if it is still alive)

```go
// voc: VOptionCtrt

status, err := vo.OptionStatus()
if err != nil {
	log.Fatal(err)
}
fmt.Println(status)
```

Example output

```
false
```

#### Max issue num

The maximum issue of the option tokens.

```go
// voc: VOptionCtrt

maxIssueNum, err := vo.MaxIssueNum()
if err != nil {
    log.Fatal(err)
}
fmt.Println(maxIssueNum)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Reserved option

The reserved option tokens remaining in the pool.

```go
// voc: VOptionCtrt

amount, err := vo.ReservedOption()
if err != nil {
    log.Fatal(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Reserved proof

The reserved proof tokens remaining in the pool.

```go
// voc: VOptionCtrt

amount, err := vo.ReservedProof()
if err != nil {
    log.Fatal(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Price

The price of the contract creator.

```go
// voc: VOptionCtrt

amount, err := vo.Price()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Price unit

The price unit of the contract creator.

```go
// voc: VOptionCtrt

unit, err := vo.PriceUnit()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Token locked

The locked token amount. What kind of token?

```go
// voc: VOptionCtrt

amount, err := vo.TokenLocked()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Token collected

The amount of the base tokens in the pool.

```go
// voc: VOptionCtrt

amount, err := vo.TokenCollected()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Base token balance

Get the balance of the available base tokens.

```go
// voc: VOptionCtrt
// acnt: Account

amount, err := vo.GetBaseTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Target token balance

Get the balance of the available target tokens.

```go
// voc: VOptionCtrt
// acnt: Account

amount, err := vo.GetTargetTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Option token balance

Get the balance of the available option tokens.

```go
// voc: VOptionCtrt
// acnt: Account

amount, err := vo.GetOptionTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

#### Proof token balance

Get the balance of the available proof tokens.

```go
// voc: VOptionCtrt
// acnt: Account

amount, err := vo.GetProofTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amount)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1)})
```

### Actions

#### Supersede

Transfer the ownership of the contract to another account.

```go
// voc: VOptionCtrt
// by: Account
// newIssuer: Account

resp, err := vo.Supersede(by, string(newIssuer.Addr.B58Str()), "attachment")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659668629905596000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6gmM7UxzUyRJXidy2DpXXMvrPqEF9hR1eAqsmh33J6eL) Addr:vsys.Str(AU6BNRK34SLuc27evpzJbAswB6ntHV2hmjD) Signature:vsys.Str(26gn57S3xmf1XVcrhcnmSEp82j6v7sMsskBj1pc8NZt5Gd5jKijkmUwgb52LLsnPepWfj7VH1TurTCcp3GrJSsMf)}]} CtrtId:vsys.Str(CFAAxTu44NsfwMUfpmVd6y4vuN9xQNVFtGa) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1CC6B9Tu94MJrtVckkunxuvwR4ixhCVVLeT4ZX9NUBN6KUifUdbuevxsezvw45po5HFnmyFYAchxWVfwG3zAdK5H729k8VxbmehT2pTXJ1T2xKh) Attachment:vsys.Str()})
```

#### Activate

Activate the V Option contract to store option token and proof token into the pool.

```go
// acnt: *vsys.Account
// voc: *vsys.VOptionCtrt
// maxIssueNum: float64
// price: float64
// priceUnit: int64

resp, err := vo.Activate(acnt, maxIssueNum, price, priceUnit, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(CvTpiBHomsKzQt8kRCr68p2UJwgQNnN1TCfWcj8FEqQ3) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659691634700932000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3axewnrQtwArgQnZKQaNrZscSEwJePij6zQf2tqQD6F6C3GhTKT19hG7AF1c2LZBg6wPvXPBuLuzkwMzS8vVLJaN)}]} CtrtId:vsys.Str(CFE4vYppQvKkVZuA8eeMT2LBArnW31A4zRh) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(12oCrKY2h2JDu8D8RTzEMDhUdp26JXbR2iQJxPy) Attachment:vsys.Str()})
```

#### Mint

Mint target tokens into the pool to get option tokens and proof tokens. Same amount of option, proof and target tokens are used for minting. For example, if we set amount to 200, then 200 proof and option tokens will be given to acnt and 200 target tokens will be locked.

```go
// voc: *vsys.VOptionCtrt
// acnt: *vsys.Account
// amount: float64

resp, err := vo.Mint(testAcnt0, amount, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(C8E1eGeFJKPi6PqdeaNEapcwUkZ4gq9bVjdeenQ9a7ue) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659691774514664000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(oYNs3tYd6V38djFuDkWQwDiEqPsWxMvw8x7va39d4FqciYSoSBTMvo6BP2YQDEDq1EXdL7si8TUAzf5rTcDYTkb)}]} CtrtId:vsys.Str(CFEQdvsbf4ra3333aUo85yrbpsVECDcpb1M) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(14JDCrdo1xwswd) Attachment:vsys.Str()})
```

#### Unlock

Get the remaining option tokens and proof tokens from the pool before the execute time. Amount equals to the amount of Target tokens to be unclocked.

```go
// voc: *vsys.VOptionCtrt
// acnt: *vsys.Account
// amount: float64

resp, err := vo.Unlock(acnt, amount, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(BPGWXWZS3j8KPG4MDuuHVGRApHtEgvfCyMxRetwAuL7a) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659691896760717000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4ryqnEZsvgL7MffTTCTyQxKBSPPHqmPCpEdmsNJqGwTDWmKE89zPRsWzyswXvaFZx1v2h4v3QWmMMzuMfQSw6F7j)}]} CtrtId:vsys.Str(CEsCVgzgR5oousEx3Xj886bG8DWrfHNjy6G) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(14JDCrdo1xwsuu) Attachment:vsys.Str()})
```

#### Execute

Execute the V Option contract to get target token after execute time. Amount equals to the amount of option tokens to be executed.

Note that amount of `price * amount` Base Tokens need to be deposited to V Option contract by executor.

```go
// voc: *vsys.VOptionCtrt
// acnt: *vsys.Account
// amount: float64

exeTx, err := vo.Execute(acnt, amount, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(exeTx)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(8Am1sEQwEiBht2Xa6H9vrbbDwMU8yGrkNmubJgf97Q1x) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659692011434942000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(423b3MoszXK7hnApKHr7Smqs7nfNfj1MjBDbWnzFLVGBDj2twMaHkxvjrFA6nyTnYEKHMp9GXU6RTWSkjbk5cvoD)}]} CtrtId:vsys.Str(CF8UzB8nXwyYmFuTB8xw6GfxAHu9WVCguHo) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(14JDCrdo1xwstM) Attachment:vsys.Str()})
```

#### Collect

Collect the base tokens or/and target tokens from the pool depending on the amount of proof tokens after execute deadline.

```go
// voc: *vsys.VOptionCtrt
// acnt: *vsys.Account
// amount: float64

colTx, err := vo.Collect(acnt, amount, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(colTx)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(8Z56BRwXXYW2UHuRiGBWXvAfAs2NM78KQsFe1HyTHfzb) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659692049016281000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4Xk51dmSqhgXmRCoCRDDNS5WUZT8dFwCSfL3CZu5UXijjEBuBUdsPtujrKR4ep9mK1DsHi79ERzVUnWyEvPy1j63)}]} CtrtId:vsys.Str(CF8UzB8nXwyYmFuTB8xw6GfxAHu9WVCguHo) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(14JDCrdo1xwstM) Attachment:vsys.Str()})
```
