# Atomic Swap Contract

- [Atomic Swap Contract](#atomic-swap-contract)
  - [Introduction](#introduction)
  - [Usage with Go SDK](#usage-with-go-sdk)
    - [Registration](#registration)
    - [From Existing Contract](#from-existing-contract)
    - [Querying](#querying)
      - [Maker](#maker)
      - [Token ID](#token-id)
      - [Unit](#unit)
      - [Contract Balance](#contract-balance)
      - [Swap Owner](#swap-owner)
      - [Swap Recipient](#swap-recipient)
      - [Swap Puzzle](#swap-puzzle)
      - [Swap Amount](#swap-amount)
      - [Swap Expiration Time](#swap-expiration-time)
      - [Swap Status](#swap-status)
    - [Actions](#actions)
      - [Lock](#lock)
      - [solve](#solve)
      - [Withdraw after expiration](#withdraw-after-expiration)

## Introduction

[Atomic Swap](https://en.bitcoin.it/wiki/Atomic_swap) is a general algorithm to achieve the exchange between two parties without having to trust a third party.

Atomic Swap Contract is the VSYS implementation of [Atomic Swap](https://en.bitcoin.it/wiki/Atomic_swap) which supports atomic-swapping tokens on VSYS chain with other tokens(either on VSYS chain or on other atomic-swap-supporting chain like Ethereum).

We have written a helper which exclusively serve when both users' accounts and tokens are on VSYS chain.

## Usage with Go SDK

### Registration

Register an Atomic Swap Contract instance.

```go
// acnt: *vsys.Account
// tokId: string - base58 formatted token id
ac, err := vsys.RegisterAtomicSwapCtrt(acnt, tokId, "")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(ac.CtrtId)
```

Example output

```
*vsys.CtrtId(vsys.Str(CF9C7sfiQYV5bTnWqmR4ZqMzvKPwHtSxeTv))
```

### From Existing Contract

Get an object for an existing contract instance.

```go
// ch: *vsys.Chain
// ctrtId: string - base58 formatted token id
ctrtId := 'CFAAxTu44NsfwMUfpmVd6y4vuN9xQNVFtGa';
atomicCtrt, err = vsys.NewAtomicSwapCtrt(ctrtId, ch);
```

### Querying

#### Maker

The address that made this contract instance.

```go
// ac: *vsys.AtomicSwapCtrt

maker, err := ac.Maker()
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

The token ID of the token deposited into this contract.

```go
// ac: *vsys.AtomicSwapCtrt
tokId, err := ac.TokId()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(tokId)
```

Example output

```
*vsys.TokenId(vsys.Str(TWtPBJN8Sa7i5ZHpMBEeyCE8dCgoAxktMfQRqi1rT))
```

#### Unit

The unit of the token deposited into this contract.

```go
// ac: *vsys.AtomicSwapCtrt
unit, err := ac.Unit()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(unit)
```

Example output

```
vsys.Unit(1)
```

#### Contract Balance

The balance of the token deposited into this contract for the given user.

```go
// ac: *vsys.AtomicSwapCtrt
// acnt: *vsys.Account

bal, err := ac.GetCtrtBal(acnt.Addr.B58Str().Str())
if err != nil {
    log.Fatalln(err)
}
fmt.Println(bal)
```

Example output

```
*vsys.Token({Data:vsys.Amount(0) Unit:vsys.Unit(1000000)}
```

#### Swap Owner

Get the owner of the swap based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: *vsys.AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

owner, err := ac.GetSwapOwner(makerLockTxId)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(owner)
```

Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Swap Recipient

Get the recipient of the swap based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: *vsys.AtomicSwapCtrt
// makerLockTxId: str E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

recipient, err := ac.GetSwapRecipient(makerLockTxId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(recipient)
```

Example output

```
*vsys.Addr(vsys.Str(AU1KWrn3sFwddbZjfeKnauh4zAYiDTmo9gM))
```

#### Swap Puzzle

Get the hashed puzzle(i.e. secret) of the swap based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

puzzle, err := ac.GetSwapPuzzle("FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(puzzle)
```

Example output

```
vsys.Str(DYu3G8aGTMBW1WrTw76zxQJQU4DHLw9MLyy7peG4LKkY)
```

#### Swap Amount

Get the token amount locked into the swap based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

amnt, err := ac.GetSwapAmount("FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(amnt)
```

Example output

```
*vsys.Token({Data:vsys.Amount(10000) Unit:vsys.Unit(100)})
```

#### Swap Expiration Time

Get the expiration time of the swap based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

exp, err := ac.GetSwapExpiredTime("FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(exp)
```

Example output

```
vsys.VSYSTimestamp(1646984339000000000)
```

#### Swap Status

Get the status of the swap(if the swap is active) based on the given token-locking transaction ID(e.g. the transaction ID obtained from taking the maker locking action).

```go
// ac: AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

st, err := ac.GetSwapStatus("FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(st)
```

Example output

```
true
```

### Actions

#### Lock

For this lock function sample code, we consider it as the maker lock of the atomic swap.

Maker is the one who first creates the swap contract, and creates the secret.

If the lock is executed by the taker, we need the hashed secret bytes from the maker lock transaction.

```go
// ac: AtomicSwapCtrt
// acnt0: Account
// acnt1: Account
// lockAmount: number
// maker lock sample code
resp, err := as.Lock(
    acnt0,
    1,
    acnt1.Addr.B58Str().Str(),
    vsys.Sha256Hash([]byte("secret")),
    time.Now().Unix()+30,
    "",
)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(7fVawdkouMq7Yx9qfeoc1ruUViXm17kyzsnuiVmUtdn2) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659664120132480000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5D8J1N6DJ5153oLekSmq4ygoYRCtRWc4AFJP6tDJJdZ9yRdRM8aRtd9xT6RbPeL8uPktuRjRdk5a9wVcLLSv5shS)}]} CtrtId:vsys.Str(CF7BajMvKZLdu9cxx8eWR1csQkWRPJMCg4J) FuncIdx:vsys.FuncIdx(0) FuncData:vsys.Str(1CC6B9Tu94MGKJPwXueR5Ei73n3DcXfRz6cFM1wvTaMrhgqJCAzeP1iHGy4e23fWrqXaaki1XieFmNZV5vyChBUxsHoU7X5z7pbwCzr69NkDJG7) Attachment:vsys.Str()})
```

#### solve

If the solve function is executed by the maker, then the maker takes the tokens locked by the taker and reveals the plain text of the hashed secret.

If the solve function is executed by the taker, then the taker will get the revealed secret and takes the tokens locked by the maker.

The following sample code shows how maker solve runs.

```go
// ac: AtomicSwapCtrt instance of TAKER's
// maker: Account
// takerLockTxId: string E.g. "D5ZPPhw7y4eWcL6zBNWNHdWf9jGxPAi5XCP5KxuZzirP"
// secret: string E.g. "abc"

resp, err = ac.Solve(acnt0, takerLockTxId, "secret", "")
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(A8NX3hvZazVpuEqkqMaLAWp65wWvG3b34dYWaXiNMr4v) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659664283048908000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(4yiCBcT5CBDuchG3gFATxP5m1bHroLQij3EJ1B5mkbxhhP9WtCeEcrR2imM1ohQRsvLwn3nWarAepweGgYhepJbA)}]} CtrtId:vsys.Str(CF3Udwmn9SCxXJBhnHtfjfHdti5hTtq6UvL) FuncIdx:vsys.FuncIdx(1) FuncData:vsys.Str(13w3j8UHjc9FdJqXGwF3Sanigbsm7jAxgNyem48QUAL3EsGjZcs7jTeaXmLLod) Attachment:vsys.Str()})
```

#### Withdraw after expiration

Either the maker or taker withdraws the tokens from the contract after the expiration time.

The example below shows the withdraw after expiration by the maker. It is the same for the taker.

```go
// ac: AtomicSwapCtrt
// makerLockTxId: string E.g. "FHZdvf3yyWuDnNTYeR6MZKTEqLJ1QxKfrDBqFrHDVBeJ"

resp, err = ac.ExpWithdraw(acnt, makerLockTxId, "")
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```

Example output

```
*vsys.BroadcastExecuteTxResp({TxBasic:{Type:vsys.TxType(9) Id:vsys.Str(6NVnr89iL2xZbWeiS7MDi9FdhbN41TWqMM7kU7VQ7EZC) Fee:vsys.VSYS(30000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659664429685848000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(2PvathWebooKLBrJpwPzdTrM1sLLiWHCrJ6TzSfh5Tkm) Addr:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Signature:vsys.Str(5wZnFuttQVjDWttvbai8rkSeJGkaMQyZMhpbn4ej5vKpKksex9ZqRTkabLPbqn1jyNCqHUttLRXuCLCJrNizCUEh)}]} CtrtId:vsys.Str(CF6GHHZdxB3SeffaGyeabBeB6idcTA82fMS) FuncIdx:vsys.FuncIdx(2) FuncData:vsys.Str(1TeCHmc67edpfTGr9A4HsqjZbPtkezcLuNuqf7DBGwM5sNXaQ) Attachment:vsys.Str()})
```
