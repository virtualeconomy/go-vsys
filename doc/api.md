# Api

- [Api](#api)
    - [Introduction](#introduction)
    - [Usage with Go SDK](#usage-with-go-sdk)
        - [Instantiation](#instantiation)
        - [Properties](#properties)
            - [Session](#session)
            - [API Group: Blocks](#api-group-blocks)
            - [API Group: Utils](#api-group-utils)
            - [API Group: Node](#api-group-node)
            - [API Group: Transactions](#api-group-transactions)
            - [API Group: Contract](#api-group-contract)
            - [API Group: Addresses](#api-group-addresses)
            - [API Group: Database](#api-group-database)
            - [API Group: Leasing](#api-group-leasing)
            - [API Group: VSYS](#api-group-vsys)
        - [Actions](#actions)
            - [Make HTTP GET Request](#make-http-get-request)
            - [Make HTTP POST Request](#make-http-post-request)

## Introduction
Nodes in VSYS net can expose RESTful APIs for users to interact with the chain(e.g. query states, broadcast transactions).

## Usage with Go SDK
In Go SDK we have `NodeAPI` struct that serves as an API wrapper for calling node APIs.

### Instantiation
Create an object of `NodeAPI`

```go
const HOST = "http://veldidina.vos.systems:9928"
api := vsys.NewNodeAPI(HOST)
fmt.Println(api)
```
Example output

```
*vsys.NodeAPI(http://veldidina.vos.systems:9928)
```

### Properties

#### Session
The `*req.ClientSession` structs that records the HTTP session(e.g. host) is embedded in `NodeAPI` struct.

```go
// api: *vsys.NodeAPI
fmt.Printf("%T\n", api.Client)
```
Example output

```
*req.Client
```

### Groups

#### API Group: Blocks
The group of APIs that share the prefix `/blocks` in file `api_blocks.go`.

* GetCurBlockHeight
* GetHeightBySignature
* GetLast
* GetBlockAt
* GetBlocksWithin
* GetAvgDelay

#### API Group: Utils
The group of APIs that share the prefix `/utils` in file `api_utils.go`.

```go
// api: pv.NodeAPI
// /utils/hash/fast
resp, err := api.FastHash("foo")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp.Hash)
```
Example output

```
vsys.Str(DT9CxyH887V4WJoNq9KxcpnF68622oK3BNJ41C2TvESx)
```

#### API Group: Node
The group of APIs that share the prefix `/node` in `api_node.go` file.

* GetNodeStatus
* GetNodeVersion

```go
// /node/status
resp, err := api.GetNodeStatus()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)

// /node/version
resp1, err := api.GetNodeVersion()
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp1.Version)
```
Example output

```
*vsys.NodeStatusResp({BlockchainHeight:vsys.Height(3526717) StateHeight:vsys.Height(3526717) UpdatedTimestamp:vsys.VSYSTimestamp(1660011108009563339) UpdatedDate:2022-08-09 02:11:48.009 +0000 UTC})
vsys.Str(VSYS Core v0.4.1)
```

#### API Group: Transactions
The group of APIs that share the prefix `/transactions` in file `api_tx.go`.

* GetTxInfo

```go
// api: *vsys.NodeAPI

txId := "Eui1yaRcE4jCnf4yBawroxSvqGa54WyQV9LjHkRHVvPd"
// /transactions/info/{tx_id}
resp, err := api.GetTxInfo(txId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.GenesisTxInfoResp({TxGeneral:{TxBasic:{Type:vsys.TxType(1) Id:vsys.Str(7YWD62iUS2oSrRnSggNjHBSqmfpXorrYtM3hYSCXKrWY) Fee:vsys.VSYS(0) FeeScale:vsys.VSYS(0) Timestamp:vsys.VSYSTimestamp(1631506399220088162) Proofs:[]} Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0) Height:vsys.Height(1)} SlotId:0 Signature:vsys.Str(33meMP9e6W6F8EjEGQiJZ4rdWEHcRYyEF3zL2ox28eJEeZCgGaMjYkqR35gea1165a3FdfoTj2NXYgedHuGCyiE5) Recipient:vsys.Str(ATwAPYdriV1aRXAWYmLViW7Y6K5Jb5bZDkT) Amount:vsys.VSYS(50000000000000000)})
```

#### API Group: Contract
The group of APIs that share the prefix `/contract` in file `api_ctrt.go`.

```go
// api: *vsys.NodeAPI

tokId := "TWu2qeuPdfjFQ7HdZGqjSYCSTh3m9k7kCttv7NmSx"

// /contract/tokenInfo/{tok_id}
resp, err := api.GetTokInfo(tokId)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.TokInfoResp({TokId:vsys.Str(TWu2qeuPdfjFQ7HdZGqjSYCSTh3m9k7kCttv7NmSx) CtrtId:vsys.Str(CF6sVHb2Y8i5Cqcw5yZL1m2PmaTvk1KdB2T) Max:vsys.Amount(10000) Total:vsys.Amount(3000) Unit:vsys.Unit(100) Description:vsys.Str()})
```

#### API Group: Addresses
The group of APIs that share the prefix `/addresses` in file `api_addr.go`.

* GetBal
* GetBalDetails
* GetAddr

```go
// api: pv.NodeAPI

addr := "AUA1pbbCFyFSte38uENPXSAhZa7TH74V2Tc"

// /addresses/balance/{addr}
resp, err := api.GetBal(addr)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.BalResp({Address:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Confirmations:0 Balance:vsys.VSYS(16749520000000)})
```

#### API Group: Leasing
The group of APIs that share the prefix `/leasing` in `api_leasing.go` file.

* Broadcast Lease
* BroadcastCancelLease

```go
// api: *vsys.NodeAPI

var p = &vsys.BroadcastLeasingPayload{
	SenderPubKey: vsys.Str("6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv"),
	Recipient: vsys.Str("AUCUg4dFgn52U2PgZb9YhehBXnSqp8EMRqH"),
	Amount: vsys.VSYS(100000000),
	Fee: vsys.VSYS(10000000),
	FeeScale: vsys.FeeScale(100),
	Timestamp: vsys.VSYSTimestamp(1660024867890262000),
	Signature: vsys.Str("4LA5m6LRkKudWG8Mk7r3M3YR2X235arY6mWKnJKc11Cb14H76QXdn2AKa833g27kCjNv9ua6mxu9ujcmPmh8aouq"),
}

// /leasing/broadcast/lease
resp, err := ch.NodeAPI.BroadcastLease(p)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
*vsys.BroadcastLeaseTxResp({TxBasic:{Type:vsys.TxType(3) Id:vsys.Str(DBuEk9uME62bJgQYn2bb6XNpAub9XgfWWAFcgJDQmkos) Fee:vsys.VSYS(10000000) FeeScale:vsys.VSYS(100) Timestamp:vsys.VSYSTimestamp(1660025413566348000) Proofs:[{ProofType:vsys.Str(Curve25519) PubKey:vsys.Str(6VH5QC2ktUA5UK4j6c4hxQTZi4cm9jdNYhnCQV2rT4Wv) Addr:vsys.Str(AU8xJNjE5RNo8hmPYA1bSgQzPKYNgejytiP) Signature:vsys.Str(5o8PeTfS2cRKa5Qb1MDyVok6k9vngM8ZTgiu9JaywDoTR3vA9w7TMixkj2ySYJy5fQhsubpzJjkusuvbJJQnCaSK)}]} Supernode:vsys.Str(AUCUg4dFgn52U2PgZb9YhehBXnSqp8EMRqH) Amount:vsys.VSYS(100000000)})
```

#### API Group: VSYS
The group of APIs that share the prefix `/vsys` in file `api_vsys.go`.

```go
// TODO: implement BroadcastPayment
// api: pv.NodeAPI

payload = {
    'senderPublicKey': '6gmM7UxzUyRJXidy2DpXXMvrPqEF9hR1eAqsmh33J6eL',
    'recipient': 'AU5NsHE8eC2guo3JobD8jrGvnEDQhBP8GtW',
    'amount': 10000000000,
    'fee': 10000000,
    'feeScale': 100,
    'timestamp': 1646993201931712000,
    'attachment': '',
    'signature': 'mjmu9CwQiUhtUkgVXFefpw8GM9Zypjf64pAufuUK5SvaGc9x8m9qZo8aRprnw7DmWRT4YQyPStTCERomGncRSMd',
}
resp = await api.vsys.broadcast_payment(payload)
print(resp)
```
Example output

```
{'type': 2, 'id': 'D2UUnSX9gWnsTWW2tEs1BoF4dgZeyZRXQETrEUjUNrns', 'fee': 10000000, 'feeScale': 100, 'timestamp': 1646993201931712000, 'proofs': [{'proofType': 'Curve25519', 'publicKey': '6gmM7UxzUyRJXidy2DpXXMvrPqEF9hR1eAqsmh33J6eL', 'address': 'AU6BNRK34SLuc27evpzJbAswB6ntHV2hmjD', 'signature': 'mjmu9CwQiUhtUkgVXFefpw8GM9Zypjf64pAufuUK5SvaGc9x8m9qZo8aRprnw7DmWRT4YQyPStTCERomGncRSMd'}], 'recipient': 'AU5NsHE8eC2guo3JobD8jrGvnEDQhBP8GtW', 'amount': 10000000000, 'attachment': ''}
```

### Actions

#### Make HTTP GET Request
Make an HTTP GET request to given endpoint.

```go
// api: *vsys.NodeAPI

resp, err := api.Get("/node/version")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
{
  "version" : "VSYS Core v0.4.1"
}
```

#### Make HTTP POST Request
Make an HTTP POST request to given endpoint.

```go
// api: *vsys.NodeAPI

resp, err := api.Post("/utils/hash/fast", "foo")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(resp)
```
Example output

```
{
  "message" : "foo",
  "hash" : "DT9CxyH887V4WJoNq9KxcpnF68622oK3BNJ41C2TvESx"
}
```
