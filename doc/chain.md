# Chain

- [Chain](#chain)
    - [Introduction](#introduction)
    - [Usage with Go SDK](#usage-with-go-sdk)
        - [Properties](#properties)
            - [Api](#api)
            - [Chain ID](#chain-id)
            - [Height](#height)
            - [Last block](#last-block)
        - [Actions](#actions)
            - [Get the Block at a Certain Height](#get-the-block-at-a-certain-height)
            - [Get Blocks within a Certain Range](#get-blocks-within-a-certain-range)

## Introduction
Chain is a logical concept that represents the abstract data structure where transactions are packed into a block and blcoks are chained together by including the hash from the last block.

In VSYS, there are 2 types of chains:
- mainnet
- testnet

They have different chain IDs, namely `M` for mainnet & `T` for testnet, which will be used in cases like the address of an account.

In other words, the same pair of seed and nonce will lead to different account addresses on mainnet & testnet.

## Usage with Go SDK
In Go SDK we have an `Chain` Struct that represents the chain.

### Properties

#### Api
The `NodeAPI` object that serves as the API wrapper for calling RESTful APIs that exposed by a node in the VSYS blockchain.

```go
// ch: *vsys.Chain
fmt.Println(ch.NodeAPI)
```
Example output

```
*vsys.NodeAPI(http://veldidina.vos.systems:9928)
```

#### Chain ID
The chain ID.

```go
// ch: *vsys.Chain
fmt.Println(ch.ChainID)
```
Example output

```
vsys.ChainID(T)
```

#### Height
The current height of blocks on the chain.

Note that it is queried by calling RESTful APIs of a node. Technically speaking, the result is of the node. It can be used as of the chain as long as the node synchronises with other nodes well.

```go
// ch: *vsys.Chain
height, err := ch.Height()
if err != nil {
	log.Fatalln(err)
}
fmt.Println(height)
```
Example output

```
vsys.Height(3516500)
```

#### Last block
The last block on the chain.

Note that it is queried by calling RESTful APIs of a node. Technically speaking, the result is of the node. It can be used as of the chain as long as the node synchronises with other nodes well.

```go
// ch: *vsys.Chain

b, err := ch.LastBlock()
if err != nil {
	log.Fatalln(err)
}
fmt.Println(b)
```
Example output

```
*vsys.LastBlockResp({Version:1 Timestamp:vsys.VSYSTimestamp(1659950412012786044) Reference:vsys.Str(4FUv7naXqNHJ9zS3QBU3ykUe6LuD2nc4mPe9jhqnuJvhEAf27yHLV2eKvZ4B314goXvmyBxbXoyM4T6oz3xP1u31) SPOSConsensus:{MintTime:vsys.VSYSTimestamp(1659950412000000000) MintBalance:vsys.VSYS(259298751165120)} ResourcePricingData:{Computation:0 Storage:0 Memory:0 RandomIO:0 SequentialIO:0} TransactionMerkleRoot:vsys.Str(CsAjbGfdTcgfVdQ43AAQg8Cf8sAVUpUo9P8PPQhecwub) Transactions:[{Type:vsys.TxType(5) Id:vsys.Str(AmARPyQWJdTYX3UEvz7uDcGBZ92CUsAsQf6aNCeeMT9w) Recipient:vsys.Str(AU6sMeLdsswqDQrw4RDo5PVxdGh1v6JDv6t) Timestamp:vsys.VSYSTimestamp(1659950412012786044) Amount:vsys.Amount(900000000) CurrentBlockHeight:vsys.Height(3516601) Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0)}] Generator:vsys.Str(AU6sMeLdsswqDQrw4RDo5PVxdGh1v6JDv6t) Signature:vsys.Str(5mHYWeR9bPJDjgC4xV3awWg6uB7gnoSeRSyJWR1TzdAfBmDBwJgHRfiNRw1eha7nxCvTDn4fLWh4Rk7KyWweaMhs) Fee:vsys.VSYS(0) Blocksize:330 Height:vsys.Height(3516601) TransactionCount:1})
```

### Actions

#### Get the Block at a Certain Height
Get the block at a certain height.

```go
// ch: *vsys.Chain

b, err := ch.GetBlockAt(3516601)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(b)
```
Example output

```
*vsys.BlockResp({Version:1 Timestamp:vsys.VSYSTimestamp(1659950412012786044) Reference:vsys.Str(4FUv7naXqNHJ9zS3QBU3ykUe6LuD2nc4mPe9jhqnuJvhEAf27yHLV2eKvZ4B314goXvmyBxbXoyM4T6oz3xP1u31) SPOSConsensus:{MintTime:vsys.VSYSTimestamp(1659950412000000000) MintBalance:vsys.VSYS(259298751165120)} ResourcePricingData:{Computation:0 Storage:0 Memory:0 RandomIO:0 SequentialIO:0} TransactionMerkleRoot:vsys.Str(CsAjbGfdTcgfVdQ43AAQg8Cf8sAVUpUo9P8PPQhecwub) Transactions:[{Type:vsys.TxType(5) Id:vsys.Str(AmARPyQWJdTYX3UEvz7uDcGBZ92CUsAsQf6aNCeeMT9w) Recipient:vsys.Str(AU6sMeLdsswqDQrw4RDo5PVxdGh1v6JDv6t) Timestamp:vsys.VSYSTimestamp(1659950412012786044) Amount:vsys.Amount(900000000) CurrentBlockHeight:vsys.Height(3516601) Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0)}] Generator:vsys.Str(AU6sMeLdsswqDQrw4RDo5PVxdGh1v6JDv6t) Signature:vsys.Str(5mHYWeR9bPJDjgC4xV3awWg6uB7gnoSeRSyJWR1TzdAfBmDBwJgHRfiNRw1eha7nxCvTDn4fLWh4Rk7KyWweaMhs) Fee:vsys.VSYS(0) Blocksize:330 Height:vsys.Height(3516601) TransactionCount:1})
```

#### Get Blocks within a Certain Range
Get blocks within a certain range.

NOTE that the max length of the range is 100.

```go
// ch: *vsys.Chain
start := 3516601
end := 3516601 + 2

blocks := ch.GetBlocksWithin(start, end)
fmt.Println(blocks)

start = 1355645
end = 1355645 + 200

blocks := ch.GetBlocksWithin(start, end)
fmt.Println(blocks)
```
Example output

```
[*vsys.BlockResp({Version:1 Timestamp:vsys.VSYSTimestamp(1646984676006693040) Reference:vsys.Str(2Z6LXeq27kPX1UGZU8eJiz72JVNSGiCC234oGQZbMaD7TniDDPjntSTp3zM3xL2MVvco1Dfm3h3PFizFL4PSRDA4) SPOSConsensus:{MintTime:vsys.VSYSTimestamp(1646984676000000000) MintBalance:vsys.VSYS(50114224369788003)} ResourcePricingData:{Computation:0 Storage:0 Memory:0 RandomIO:0 SequentialIO:0} TransactionMerkleRoot:vsys.Str(cKWBKEtc5XQGjMobespD5cJydJpzhF1SnafBSS1q1is) Transactions:[{Type:vsys.TxType(5) Id:vsys.Str(9e1ToB1zuCPoE7zrrj1L14gvcYJFhX9QdKd3NCyTGRG2) Recipient:vsys.Str(AU7fEwBgHpe6oeH1iuo2mE5TMCrBxPR8LFc) Timestamp:vsys.VSYSTimestamp(1646984676006693040) Amount:vsys.Amount(900000000) CurrentBlockHeight:vsys.Height(1355645) Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0)}] Generator:vsys.Str(AU7fEwBgHpe6oeH1iuo2mE5TMCrBxPR8LFc) Signature:vsys.Str(59enpKJUjVsvtgbChWuinj9Ds5CfUn7ChPxFgfsQZAAmsU5MGDJJGE6sn2n5UpT49URR69MkcD4ofvFf7zLt5BPq) Fee:vsys.VSYS(0) Blocksize:330 Height:vsys.Height(1355645) TransactionCount:1}) *vsys.BlockResp({Version:1 Timestamp:vsys.VSYSTimestamp(1646984682014519750) Reference:vsys.Str(59enpKJUjVsvtgbChWuinj9Ds5CfUn7ChPxFgfsQZAAmsU5MGDJJGE6sn2n5UpT49URR69MkcD4ofvFf7zLt5BPq) SPOSConsensus:{MintTime:vsys.VSYSTimestamp(1646984682000000000) MintBalance:vsys.VSYS(50101287145150374)} ResourcePricingData:{Computation:0 Storage:0 Memory:0 RandomIO:0 SequentialIO:0} TransactionMerkleRoot:vsys.Str(J3WDruFNxu11CKFNkNKP4Lb5YEUDV1fJntXpmmvcaLhV) Transactions:[{Type:vsys.TxType(5) Id:vsys.Str(6YwWoTZBZAungxCyiKJY5VkFsUzb3WhioKTcATSzMQ9) Recipient:vsys.Str(ATxtBDygMvWtvh9xJaGQn5MdaHsbuQxbjiG) Timestamp:vsys.VSYSTimestamp(1646984682014519750) Amount:vsys.Amount(900000000) CurrentBlockHeight:vsys.Height(1355646) Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0)}] Generator:vsys.Str(ATxtBDygMvWtvh9xJaGQn5MdaHsbuQxbjiG) Signature:vsys.Str(nytRh8LMivri3dkjHSvB6ATgKzq7A2R8jBGEtu8NaWB5jjCr4Uj4PVpystLYeQcLQP6ocSXDka2fM26PvwKUpoa) Fee:vsys.VSYS(0) Blocksize:330 Height:vsys.Height(1355646) TransactionCount:1}) *vsys.BlockResp({Version:1 Timestamp:vsys.VSYSTimestamp(1646984688008446521) Reference:vsys.Str(nytRh8LMivri3dkjHSvB6ATgKzq7A2R8jBGEtu8NaWB5jjCr4Uj4PVpystLYeQcLQP6ocSXDka2fM26PvwKUpoa) SPOSConsensus:{MintTime:vsys.VSYSTimestamp(1646984688000000000) MintBalance:vsys.VSYS(50103227168492760)} ResourcePricingData:{Computation:0 Storage:0 Memory:0 RandomIO:0 SequentialIO:0} TransactionMerkleRoot:vsys.Str(64YFFj15L11hw2yGVdiYo71gKA2g1rdaHkqLThs8U8YH) Transactions:[{Type:vsys.TxType(5) Id:vsys.Str(6EsDDPDAhxa3bZuMrospz2BLQyci88LJ3zaD1ViTL3cf) Recipient:vsys.Str(AU1EWbfR8mTwbvzgnY8wdpLy3vEvF64WSEE) Timestamp:vsys.VSYSTimestamp(1646984688008446521) Amount:vsys.Amount(900000000) CurrentBlockHeight:vsys.Height(1355647) Status:vsys.Str(Success) FeeCharged:vsys.VSYS(0)}] Generator:vsys.Str(AU1EWbfR8mTwbvzgnY8wdpLy3vEvF64WSEE) Signature:vsys.Str(3RSNdKZTCaGwKgNBHUTg2cbLGTRZU37kbrNnsp2sp4hdrYA4eNbpyHK4qv4PPiQmcwLDhSVD7A4ng53KPLmoPMvd) Fee:vsys.VSYS(0) Blocksize:330 Height:vsys.Height(1355647) TransactionCount:1})]
{'error': 10, 'message': 'Too big sequences requested'}
```
