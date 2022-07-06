package vsys

import (
	"fmt"
)

type PaymentTxReq struct {
	TxType     TxType
	Recipient  *Addr
	Amount     VSYS
	Timestamp  VSYSTimestamp
	Attachment Str
	Fee        VSYS
}

func NewPaymentTxReq(
	recipient *Addr,
	amount VSYS,
	timestamp VSYSTimestamp,
	attachment Str,
	fee VSYS,
) *PaymentTxReq {
	return &PaymentTxReq{
		TxType:     TX_TYPE_PAYMENT,
		Recipient:  recipient,
		Amount:     amount,
		Timestamp:  timestamp,
		Attachment: attachment,
		Fee:        fee,
	}
}

func (p *PaymentTxReq) DataToSign() Bytes {
	txTypeBytes := p.TxType.Serialize()
	tsBytes := p.Timestamp.Serialize()
	amountBytes := PackUInt64(uint64(p.Amount))
	feeBytes := p.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	rcptBytes := p.Recipient.ByteSlice()
	attachmentLenBytes := PackUInt16(uint16(p.Attachment.RuneLen()))
	attachmentBytes := p.Attachment.Bytes().ByteSlice()

	size := len(txTypeBytes) +
		len(tsBytes) +
		len(amountBytes) +
		len(feeBytes) +
		len(feeScaleBytes) +
		len(rcptBytes) +
		len(attachmentLenBytes) +
		len(attachmentBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, tsBytes...)
	b = append(b, amountBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, rcptBytes...)
	b = append(b, attachmentLenBytes...)
	b = append(b, attachmentBytes...)

	return b
}

func (p *PaymentTxReq) BroadcastPaymentPayload(priKey *PriKey, pubKey *PubKey) (*BroadcastPaymentPayload, error) {
	sig, err := priKey.Sign(p.DataToSign())

	if err != nil {
		return nil, fmt.Errorf("Payload: %w", err)
	}

	return &BroadcastPaymentPayload{
		SenderPubKey: pubKey.B58Str(),
		Recipient:    p.Recipient.B58Str(),
		Amount:       p.Amount,
		Fee:          p.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    p.Timestamp,
		Attachment:   p.Attachment,
		Signature:    sig.B58Str(),
	}, nil
}

func (p *PaymentTxReq) String() string {
	return fmt.Sprintf("%T(%+v)", p, *p)
}
