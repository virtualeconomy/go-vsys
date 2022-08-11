# Account


- [Account](#account)
    - [Introduction](#introduction)
    - [Usage with Go SDK](#usage-with-go-sdk)
        - [Create Account](#create-account)
        - [From Wallet](#from-wallet)
        - [From Private Key & Public Key](#from-private-key--public-key)
        - [Properties](#properties)
            - [Chain](#chain)
            - [Api](#api)
            - [Key Pair](#key-pair)
            - [Address](#address)
            - [VSYS Balance](#vsys-balance)
            - [VSYS Available Balance](#vsys-available-balance)
            - [VSYS Effective Balance](#vsys-effective-balance)
        - [Actions](#actions)
            - [Get Token Balance](#get-token-balance)
            - [Pay](#pay)
            - [Lease](#lease)
            - [Cancel Lease](#cancel-lease)
            - [DB Put](#db-put)
            - [Register Contract](#register-contract)
            - [Execute Contract](#execute-contract)


## Introduction
Account is the basic entity in the VSYS blockchain that pocesses tokens & can take actions(e.g. send tokens, execute smart contracts).

There are 2 kinds of accounts:
- user account: the most common account.
- contract account: the account for a smart contract instance.

The key difference between them lies in whether they have a private key.

## Usage with Go SDK

In Go SDK we have an `Account` struct that represents a user account on the VSYS blockchain.

### Create Account

#### From Wallet
The `Account` can be instantiated from `Wallet` by `GetAcccount` function from the given `Chain` struct & nonce.

```go
// ch: *vsys.Chain
// wal: *vsys.Wallet
acnt0, err := wal.GetAccount(ch, 0)
```

#### From Private Key & Public Key
The `Account` object can be constructed by a private key & opionally along with a public key.

If the public key is omitted, it will be derived from the private key.
If the public key is provided, it will be verified against the private key.

```go
// ch: *vsys.Chain

acnt0, err := vsys.NewAccountFromPriKeyStr(ch, 'your_private_key')
if err != nil {
	log.Fatalln(err)
}

priKey, err := vsys.NewPriKeyFromB58Str('your private key in base 58 encoded format')
if err != nil {
    log.Fatalln(err)
}
acnt1, err := vsys.NewAccount(ch, priKey)
if err != nil {
    log.Fatalln(err)
}
```

### Struct Fields

#### Chain
The `Chain` object that represents the VSYS blockchain this account is on.

```go
// acnt: *vsys.Account
fmt.Println(acnt.Chain)
```
Example output

```
*vsys.Chain({NodeAPI:*vsys.NodeAPI(http://veldidina.vos.systems:9928) ChainID:vsys.ChainID(T)})
```

#### Address
The address of the account.

```go
// acnt: *vsys.Account
fmt.Println(acnt.Addr)
```
Example output

```
*vsys.Addr(vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP))
```

#### Private/Public keys
The private/public keys of the account.

```go
// acnt: *vsys.Account
fmt.Println(acnt.PriKey)
fmt.Println(acnt.PubKey)
```
Example output

```
Note that these keys are invalid and have been randomly generated.
*vsys.PriKey(vsys.Str(4zsij8MHeYaWhFS5vihLz17aCQNFWc7LWqTLMqAHfsTS))
*vsys.PubKey(vsys.Str(37CrEH855PA3TXLFMycN9nF5sUsobPoaUr9vYJahofFt))
```

### Methods

#### Api
The `NodeAPI` object that serves as the API wrapper for calling RESTful APIs that exposed by a node in the VSYS blockchain.

```go
// acnt: *vsys.Account
fmt.Println(acnt.API())
```
Example output

```
*vsys.NodeAPI(http://veldidina.vos.systems:9928)
```

#### VSYS Balance
The VSYS ledger(regular) balance of the account.

```go
// acnt: *vsys.Account
bal, err := fmt.Println(acnt.Bal())
if err != nil {
	log.Fatal(err)
}
fmt.Println(bal)
```
Example output

```
vsys.VSYS(19127280000000)
```

#### VSYS Available Balance
The VSYS available balance(i.e. the balance that can be spent) of the account.
The amount leased out will be reflected in this balance.

```go
// acnt: *vsys.Account
bal, err := acnt.AvailBal()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(bal)
```
Example output

```
vsys.VSYS(17577440000000)
```

#### VSYS Effective Balance
The VSYS effective balance(i.e. the balance that counts when contending a slot) of the account.
The amount leased in & out will be reflected in this balance.

```go
// acnt: *vsys.Account
bal, err := acnt.EffBal()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(bal)
```
Example output

```
vsys.VSYS(19127280000000)
```


#### Get Token Balance
Get the account balance of the token of the given token ID.

The example below shows querying the token balance of a certain kind of token.
```go
// acnt: *vsys.Account
tokId := "TWu66r3ebS3twXNWh7aiAEWcNAaRPs1JxkAw2A3Hi"

resp, err := acnt.GetTokBal(tokId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.Token({Data:vsys.Amount(999999000000) Unit:vsys.Unit(1000000)})
```

#### Pay
Pay the VSYS coins from the action taker to the recipient.

The example below shows paying 100 VSYS coins to another account.
```go
// acnt0: *vsys.Account
// acnt1: *vsys.Account

resp, err := acnt0.Pay(acnt1.Addr.B58Str().Str(), 100, "")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.BroadcastPaymentTxResp({TxBasic:{Type:vsys.TxType(2) Id:vsys.Str(3PYRWPgVjmsvtonewhXHptjHJug2tU8X2C9qYhJzihXL) Fee:vsys.VSYS(10000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659606163249320000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(3C7b4bevuWrGXrtMTazBUJxzCi6xE3co4pwRLpK3emdcjR5GCtoTSdeqNwUqCDakV6nb6Wb7J3B82xrrnTdQPZYf)}]} Recipient:vsys.Str(ATracVxHwdYF394gXEawdZe9stB9yLH6V7q) Amount:vsys.VSYS(10000000000) Attachment:vsys.Str()})
```

#### Lease
Lease the VSYS coins from the action taker to the recipient(a supernode).

Note that the transaction ID in the response can be used to cancel leasing later.

The example below shows leasing 100 VSYS coins to a supernode.

```go
// acnt: *vsys.Account

supernodeAddr := "AUA1pbbCFyFSte38uENPXSAhZa7TH74V2Tc"
resp, err := acnt0.Lease(supernodeAddr, 1)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.BroadcastLeaseTxResp({TxBasic:{Type:vsys.TxType(3) Id:vsys.Str(4RtnjPL9HxkGGsgXoLrxdi8bd5qFQ5w34TYRrnW6syB2) Fee:vsys.VSYS(10000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1659930691046964000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(iV4bDQSR6HKPaDCF2ZtSsd6vuWZHSqnAr6FGLyf8ZQ7CvGWYv7Zdabnc75oDoGY3iAp88zQey5dhxYvuyPTsttP)}]} Supernode:vsys.Str(AUCUg4dFgn52U2PgZb9YhehBXnSqp8EMRqH) Amount:vsys.VSYS(100000000)})
```

#### Cancel Lease
Cancel the leasing based on the leasing transaction ID.

```go
// acnt: *vsys.Account

leasingTxId := "4RtnjPL9HxkGGsgXoLrxdi8bd5qFQ5w34TYRrnW6syB2"
resp, err := acnt.CancelLease(leasingTxId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
&{{vsys.TxType(4) vsys.Str(A5fEMspmeu1fnEKm3qFg4vSxLJKFCJekCLYbZKQ2TQZZ) vsys.VSYS(10000000) vsys.VSYS(100) vsys.VSYSTimestamp(1659930697536708000) [{vsys.Str(Curve25519) vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) vsys.Str(mjuC7DWe1XPNZ6ftMn7PZ48AcQqmv1S4apRJwfpKGiZySLKtJnaTLvXaz4FdfcNLSdafEa2BQnWzBDk1kB5Ryva)}]} vsys.Str(4RtnjPL9HxkGGsgXoLrxdi8bd5qFQ5w34TYRrnW6syB2)}
```


> Ignore DB Put for now TODO: write docs
#### DB Put
Store the data with a key onto the chain(i.e. treat chain as a key-value store)

```go
# acnt: *vsys.Account

resp = await acnt.db_put(
    db_key="foo",
    data="bar",
)
print(resp)
```
Example output

```
{'type': 10, 'id': 'B5vxEnY1cPQ2GLQVLDRNKoXY2vtacEmqCAdyCxbPwmfK', 'fee': 100000000, 'feeScale': 100, 'timestamp': 1646975057234319104, 'proofs': [{'proofType': 'Curve25519', 'publicKey': '6gmM7UxzUyRJXidy2DpXXMvrPqEF9hR1eAqsmh33J6eL', 'address': 'AU6BNRK34SLuc27evpzJbAswB6ntHV2hmjD', 'signature': '37vhwQYASAwVoUENoo3xHvCJCkaAriAgnYBPdwQkt3brj4yDybhK8H1BsXpMgvvdX7ScwTQP7qtYNGeABoAzL8Qr'}], 'dbKey': 'foo', 'entry': {'data': 'bar', 'type': 'ByteArray'}}
```

The stored data can be queried by calling the node endpoint `/database/get/{addr}/{db_key}`

```bash
curl -X 'GET' \                                                                          
  'http://veldidina.vos.systems:9928/database/get/AU6BNRK34SLuc27evpzJbAswB6ntHV2hmjD/foo' \
  -H 'accept: application/json'
```
Example output

```
{
  "data" : "bar",
  "type" : "ByteArray"
}
```

Or we can use the `NodeAPI` object

```go
# api: *vsys.NodeAPI
# acnt: *vsys.Account

resp = await api.db.get(
    addr=acnt.addr.data,
    db_key="foo",
)
print(resp)
```
Example output

```
{'data': 'bar', 'type': 'ByteArray'}
```

#### Register Contract
There is no public interface for this action through `Account`. Users should always use a specific register function and pass in the `Account` struct as the action taker instead.

See [the example of registering an NFT contract instance](./smart_contract/nft_ctrt.md#registration).

#### Execute Contract
There is no public interface for this action through `Account`. Users should always use a contract methods and pass in the `Account` struct as the action taker instead.

See [the example of executing a function of an NFT instance](./smart_contract/nft_ctrt.md#issue)
