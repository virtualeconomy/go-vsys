package vsys

import "fmt"

type LeaseTxReq struct {
	TxType        TxType
	SupernodeAddr *Addr
	Amount        VSYS
	Timestamp     VSYSTimestamp
	Fee           VSYS
}

func NewLeaseTxReq(
	supernodeAddr *Addr,
	amount VSYS,
	timestamp VSYSTimestamp,
	fee VSYS,
) *LeaseTxReq {
	return &LeaseTxReq{
		TxType:        TX_TYPE_LEASE,
		SupernodeAddr: supernodeAddr,
		Amount:        amount,
		Timestamp:     timestamp,
		Fee:           fee,
	}
}

func (l *LeaseTxReq) DataToSign() Bytes {
	txTypeBytes := l.TxType.Serialize()
	supernodeBytes := l.SupernodeAddr.ByteSlice()
	amountBytes := PackUInt64(uint64(l.Amount))
	feeBytes := l.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	tsBytes := l.Timestamp.Serialize()

	size := len(txTypeBytes) +
		len(supernodeBytes) +
		len(amountBytes) +
		len(feeBytes) +
		len(feeScaleBytes) +
		len(tsBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, supernodeBytes...)
	b = append(b, amountBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, tsBytes...)

	return b
}

func (l *LeaseTxReq) BroadcastLeasingPayload(priKey *PriKey, pubKey *PubKey) (*BroadcastLeasingPayload, error) {
	sig, err := priKey.Sign(l.DataToSign())
	if err != nil {
		return nil, fmt.Errorf("BroadcastLeasingPayload: %w", err)
	}

	return &BroadcastLeasingPayload{
		SenderPubKey: pubKey.B58Str(),
		Recipient:    l.SupernodeAddr.B58Str(),
		Amount:       l.Amount,
		Fee:          l.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    l.Timestamp,
		Signature:    sig.B58Str(),
	}, nil
}

type CancelLeaseTxReq struct {
	TxType      TxType
	LeasingTxId Str
	Timestamp   VSYSTimestamp
	Fee         VSYS
}

func NewCancelLeaseTxReq(
	LeasingTXId Str,
	timestamp VSYSTimestamp,
	fee VSYS,
) *CancelLeaseTxReq {
	return &CancelLeaseTxReq{
		TxType:      TX_TYPE_LEASE_CANCEL,
		LeasingTxId: LeasingTXId,
		Timestamp:   timestamp,
		Fee:         fee,
	}
}

func (l *CancelLeaseTxReq) DataToSign() (Bytes, error) {
	txTypeBytes := l.TxType.Serialize()
	feeBytes := l.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	tsBytes := l.Timestamp.Serialize()
	txIdBytes, err := l.LeasingTxId.B58Bytes()
	if err != nil {
		return nil, err
	}

	size := len(txTypeBytes) +
		len(feeBytes) +
		len(feeScaleBytes) +
		len(tsBytes) +
		len(txIdBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, tsBytes...)
	b = append(b, txIdBytes...)

	return b, nil
}

func (l *CancelLeaseTxReq) BroadcastCancelLeasingPayload(priKey *PriKey, pubKey *PubKey) (*BroadcastCancelLeasePayload, error) {
	data, err := l.DataToSign()
	if err != nil {
		return nil, fmt.Errorf("BroadcastCancelLeasingPayload: %w", err)
	}
	sig, err := priKey.Sign(data)
	if err != nil {
		return nil, fmt.Errorf("BroadcastCancelLeasingPayload: %w", err)
	}

	return &BroadcastCancelLeasePayload{
		SenderPubKey: pubKey.B58Str(),
		TxId:         l.LeasingTxId,
		Fee:          l.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    l.Timestamp,
		Signature:    sig.B58Str(),
	}, nil
}
