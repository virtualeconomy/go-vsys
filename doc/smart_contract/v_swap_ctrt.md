# V Swap Contract

- [V Swap Contract](#v-swap-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Token A's id](#token-as-id)
      - [Token B's id](#token-bs-id)
      - [Liquidity token's id](#liquidity-tokens-id)
      - [Swap status](#swap-status)
      - [Minimum liquidity](#minimum-liquidity)
      - [Token A reserved](#token-a-reserved)
      - [Token B reserved](#token-b-reserved)
      - [Total supply](#total-supply)
      - [Liquidity token left](#liquidity-token-left)
      - [Token A's balance](#token-as-balance)
      - [Token B's balance](#token-bs-balance)
      - [Liquidity's balance](#liquiditys-balance)
    - [Actions](#actions)
      - [Supersede](#supersede)
      - [Set swap](#set-swap)
      - [Add Liquidity](#add-liquidity)
      - [Remove liquidity](#remove-liquidity)
      - [Swap token for exact base token](#swap-token-for-exact-base-token)
      - [Swap exact token for base token](#swap-exact-token-for-base-token)
      - [Swap token for exact target token](#swap-token-for-exact-target-token)
      - [Swap exact token for target token](#swap-exact-token-for-target-token)

## Introduction

V Swap is an automated market making protocol. Prices are regulated by a constant product formula, and requires no action from the liquidity provider to maintain prices.

The contract allows completely decentralised exchanges to be formed, and allows anyone to be a liquidity provider as long as they have tokens on both sides of the swap.

## Usage with Go SDK

### Registration

`tokId` is the token id of the token that deposited into this V Swap contract.

For testing purpose, you can create a new [token contract]() , then [issue]() some tokens and [deposit]() into the V Swap contract.

```go
// acnt: *vsys.Account
// tokAId: string
// tokBId: string
// tokLId: string
// minLiq: int

// Register a new V Swap contract
vs, err := vsys.RegisterVSwapCtrt(
	acnt,
	tokAId,
	tokBId,
	tokLId,
	minLiq,
	"",
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(vs.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CFDVrHkKo5SttNLce7M4LZpwBTXsNj48Z1Q))
```

### From Existing Contract

ncId is the V Swap contract's id.

```go
// ch: *vsys.Chain
// ncId: string

ncId := "CFDVrHkKo5SttNLce7M4LZpwBTXsNj48Z1Q"
vs, err := vsys.NewVSwapCtrt(ncId, ch)
```

### Querying

#### Maker

The address that made this V Swap contract instance.

```go
// vs: VSwapCtrt

maker, err := vs.Maker()
if err != nil {
    log.Fatal(err)
}
fmt.Println(maker)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Token A's id

The token A's id.

```go
// vs: *vsys.VSwapCtrt

tokId, err := vs.TokAId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Token B's id

The token B's id.

```go
// vs: *vsys.VSwapCtrt

tokId, err := vs.TokBId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Liquidity token's id

The liquidity token's id.

```go
// nc: VSwapCtrt

// vs: *vsys.VSwapCtrt

tokId, err := vs.LiqTokId()
if err != nil {
	log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWsi8XxwJqrHZTbjYMj4f3nHCTE37oRXRjfHCwahj))
```

#### Swap status

The swap status of whether or not the swap is currently active.

```go
// vs: *vsys.VSwapCtrt
// orderId: string - TransactionID of escrow order

status, err := vs.IsSwapActive()
if err != nil {
    log.Fatal(err)
}
fmt.Println(status)
```

Example output

```
true
```

#### Minimum liquidity

The minimum liquidity for the pool. This liquidity cannot be withdrawn.

```go
// vs: *vsys.VSwapCtrt

amnt, err := vs.MinLiq()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10) Unit:vsys.Unit(1000)})
```

#### Token A reserved

The amount of token A inside the pool.

```go
// vs: *vsys.VSwapCtrt

amnt, err := vs.TokAReserved()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1000)})
```

#### Token B reserved

the amount of token B inside the pool.

```go
// vs: *vsys.VSwapCtrt

amnt, err := vs.TokBReserved()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1000)})
```

#### Total supply

The total amount of liquidity tokens that can be minted.

```go
// vs: *vsys.VSwapCtrt

amnt, err := vs.TotalLiqTokSupply()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(3000) Unit:vsys.Unit(1000)})
```

#### Liquidity token left

The amount of liquidity tokens left to be minted.

```go
// vs: *vsys.VSwapCtrt

amnt, err := vs.LiqTokLeft()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(3000) Unit:vsys.Unit(1000)})
```

#### Token A's balance

The balance of token A stored within the contract belonging to the given user address.

```go
// vs: *vsys.VSwapCtrt
// acnt: *vsys.Account

amnt, err := vs.GetTokABal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(30000) Unit:vsys.Unit(1000)})
```

#### Token B's balance

The balance of token B stored within the contract belonging to the given user address.

```go
// vs: *vsys.VSwapCtrt
// acnt: *vsys.Account

amnt, err := vs.GetTokBBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(30000) Unit:vsys.Unit(1000)})
```

#### Liquidity's balance

The balance of liquidity token stored within the contract belonging to the given user address.

```go
// vs: *vsys.VSwapCtrt
// acnt: *vsys.Account

amnt, err := vs.GetLiqTokBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(30000) Unit:vsys.Unit(1000)})
```

### Actions

#### Supersede

Transfer the contract rights of the contract to a new account.

```go
// vs: *vsys.VSwapCtrt
// acnt0: *vsys.Account
// acnt1: *vsys.Account

resp, err := vs.Supersede(acnt0, string(acnt1.Addr.B58Str()), "")
if err != nil {
t.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(37uwtrVXojNXYvt2JUsc99zUWdFbn4rEkSieTUK53rWn) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659938725609255000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3bMEN4tZKZEQG78QAFGAxiRb6j6Bb3mRz7XpcTiJgKrULTK9qauTMBbUzr3AULtyqA5hs2UNcFGeqUtw31CicDQu)}]} CtrtId:vsys.Str(CEvN3K9cxAX7c1STjoy7iDyLZLimLBjM5A6) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1bscu1qPwSQ3dpRTmcaVU6cR8yjTQpcJx7S1jy) Attachment:vsys.Str()})
```

#### Set swap

Create a swap and deposit initial amounts into the pool.

```go
// vs: *vsys.VSwapCtrt
// acnt: *vsys.Account
// amountA: float64
// amountB: float64

resp, err := vs.SetSwap(acnt, amountA, amountB, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(8tTs5v5JzJYVN8rz1L8fbaKCwee1NEnSRZLZWChKwHLh) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659938798914079000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5TooaeitKCyxh9UnBetaTTJJwP17fL6rmm3ckjazriYmQPxtDmQY3jNQHtc1dyoNjmr9mwpGfNLn2mSbYCF9L92w)}]} CtrtId:vsys.Str(CEuVv6WCburKU8mYiMuNfgnkKGvVaGikDms) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(1NMvHJqnrwMGBbXnuumgvjmcrT) Attachment:vsys.Str()})
```

#### Add Liquidity

Adds liquidity to the pool. The final added amount of token A & B will be in the same proportion as the pool at that moment as the liquidity provider shouldn't change the price of the token while the price is determined by the ratio between A & B.

```go
// acnt: Account
// amountA: float64
// amountB: float64
// amountAMin: float64
// amountBMin: float64
// deadline: int64

resp, err := vs.AddLiquidity(
	acnt,
    amountA,
    amountB,
    amountAMin,
    amountBMin,
    deadline,
    "",
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(46zbfXANEX74SEdpPfvzfDGLxrqNLn5kdhQ1fWCRVxxY) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659938911050544000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3r4VqaZHgoHk19Tez7UVA1AKSX2FXseiadNRvoh4mYCBX3M9oi9nXzGrqxRGHduDiMeQS9FJzeVgRULpTjRipFf9)}]} CtrtId:vsys.Str(CFFQLA3nWK4nycJ55huDDEvrFKoFHwCoHyK) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1YkDihBAK5wpMKx2NLXXyekqmcAapH1v4MfPirQ2vRyukwHTGgn8pkDsSQ9TQes) Attachment:vsys.Str()})
```

#### Remove liquidity

Remove liquidity from the pool by redeeming token A & B with liquidity tokens.

```go
// acnt: *vsys.Account
// amountLiq: float64
// amountAMin: float64
// amountBMin: float64
// deadline: int64

resp, err := vs.RemoveLiquidity(
    acnt,
    amountLiq,
    amountAMin,
    amountBMin,
    deadline,
    "   ",
)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(5jDa71bbaDu6kF6aP4LfoeKJ2xKAQtuAWi9tLdhgCCjV) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659938996012569000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3PyJGv9uBJsnFg8soWDzz9ZrpvyCMfgmPMLCBT2NYNVdKNBq81Z2cu5K5DKeZhsEbzzVxKmpPH6U93j2brhQ2ri4)}]} CtrtId:vsys.Str(CF3njzLMcwh3H9RdvH8VjT4DdSvoTMS5iCu) FuncIdx:vsys.FuncIdx(3) FuncData:vsys.Str(18oJLXPeLXQ8KEfUZVnERVZVjnNF72HVjpvVdoPsahMDFx5rkKR) Attachment:vsys.Str()})
```

#### Swap token for exact base token

Swap token B for token A where the desired amount of token A is fixed.

```go
// acnt: Account
// amountA: number
// amountBMax: number
// deadline: number

resp, err := vs.SwapBForExactA(acnt, amountA, amountBMax, deadline, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(k3naATjxTe1B75m4YhyLd9AtBZmk3tTmpDyoNNR47YU) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659939167428571000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(A3dir8gr4V4T8EPMLY3FgdijmkdYk4ntdwFAnADaJzvT5rbvM7gBJZg8pVFzwPEvAzKNFXLHfrapZ6cNmooYTkk)}]} CtrtId:vsys.Str(CF2NVhvr1rq5DP7Eb3sxSGGoDSwMnKrAzgt) FuncIdx:vsys.FuncIdx(4) FuncData:vsys.Str(12oCrKY2h2JEPXrBoF7DadHS3eQiurNvuNzuUT1) Attachment:vsys.Str()})
```

#### Swap exact token for base token

Swap token B for token A where the amount of token B to pay is fixed.

```go
// acnt: *vsys.Account
// amountAMin: float64
// amountB: float64
// deadline: int64

resp, err := vs.SwapExactBForA(acnt, amountA, amountBMax, deadline, "")
if err != nil {
log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(5soE3KWF8kv53NmBKDnDyd5y9t1bpKx1fznwzB9Ucb2q) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659939317027008000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(5EBFVSPE1Ey5ExWENqfXxUcQFwycDuViQ1TSXDZGLAsPnK92ZYMvFMZyp9ASXnbcdhwAwUVA36ktsAFuYyjuTo9j)}]} CtrtId:vsys.Str(CF5JKyjuEiuW4s6jr773HtSMKihf6G32S3k) FuncIdx:vsys.FuncIdx(5) FuncData:vsys.Str(12oCrKY2h2JEPXrBoF7DadHS3eQiurNvyKXtgXZ) Attachment:vsys.Str()})
```

#### Swap token for exact target token

Swap token A for token B where the desired amount of token B is fixed.

```go
// acnt: *vsys.Account
// amountB: float64
// amountAMax: float64
// deadline: int64

resp, err := vs.SwapAForExactB(acnt, amountB, amountAMax, deadline, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(DZv6AesccPL3TbRHFfR8GLEPGeyJhWpBnf4q35x36Zj2) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659939575216221000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(4FibMRHDN9YCeRi6VU5AWh2VdXuPUXiax4QCM5xvRSxbvu9Khy7MWgtte5yohFb3UXJEVcTNvEBnRkSVU7ERSuCs)}]} CtrtId:vsys.Str(CEsEKfxdvEAB4pNVvncjHtr6tfFsNuMGVRq) FuncIdx:vsys.FuncIdx(6) FuncData:vsys.Str(12oCrKY2h2JEPXrBoF7DadHS3eQiurNw66cU7wd) Attachment:vsys.Str()})
```

#### Swap exact token for target token

Swap token B for token B where the amount of token A to pay is fixed.

```go
// acnt: *vsys.Account
// amountBMin: float64
// amountA: float64
// deadline: int64

resp, err := vs.SwapExactAForB(acnt, amountBmin, amountA, deadline, "")
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(Hi2vZZLfdU1TscecbbR7rDdEmU9zczeEddh9qmMjgx9H) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659939375934708000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(64i4rzhh6hLv6Enk2cFRe3Ap42T9ZsAo3YG2C2WvEiCBN5csmD6yzdLpJuZG9iwz5ty8PoQACCmHGizXzWmiRPkU)}]} CtrtId:vsys.Str(CFA6Hm7vAqBjuL9cNK1do4mJzBWuA4pwKvn) FuncIdx:vsys.FuncIdx(7) FuncData:vsys.Str(12oCrKY2h2JEPXrBoF7DadHS3eQiurNvzsRX6zs) Attachment:vsys.Str()})
```
