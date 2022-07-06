package vsys

import (
	"fmt"
)

type RegCtrtTxReq struct {
	TxType TxType

	DataStack   DataStack
	CtrtMeta    *CtrtMeta
	Timestamp   VSYSTimestamp
	Description Str
	Fee         VSYS
}

func NewRegCtrtTxReq(
	dataStack DataStack,
	ctrtMeta *CtrtMeta,
	timestamp VSYSTimestamp,
	description Str,
	fee VSYS,
) *RegCtrtTxReq {
	return &RegCtrtTxReq{
		TxType:      TX_TYPE_REGISTER_CONTRACT,
		DataStack:   dataStack,
		CtrtMeta:    ctrtMeta,
		Timestamp:   timestamp,
		Description: description,
		Fee:         fee,
	}
}

func (r *RegCtrtTxReq) DataToSign() Bytes {
	txTypeBytes := r.TxType.Serialize()
	ctrtMetaBytes := r.CtrtMeta.Serialize()
	ctrtMetaLenBytes := PackUInt16(uint16(len(ctrtMetaBytes)))

	dataStackBytes := r.DataStack.Serialize()
	dataStackLenBytes := PackUInt16(uint16(len(dataStackBytes)))

	descriptionLenBytes := PackUInt16(uint16(r.Description.RuneLen()))
	descriptionBytes := r.Description.Bytes()

	feeBytes := r.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	tsBytes := r.Timestamp.Serialize()

	size := len(txTypeBytes) +
		len(ctrtMetaLenBytes) +
		len(ctrtMetaBytes) +
		len(dataStackLenBytes) +
		len(dataStackBytes) +
		len(descriptionLenBytes) +
		len(descriptionBytes) +
		len(feeBytes) +
		len(feeScaleBytes) +
		len(tsBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, ctrtMetaLenBytes...)
	b = append(b, ctrtMetaBytes...)
	b = append(b, dataStackLenBytes...)
	b = append(b, dataStackBytes...)
	b = append(b, descriptionLenBytes...)
	b = append(b, descriptionBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, tsBytes...)

	return b
}

func (r *RegCtrtTxReq) BroadcastRegisterPayload(priKey *PriKey, pubKey *PubKey) (*BroadcastRegisterPayload, error) {
	sig, err := priKey.Sign(r.DataToSign())

	if err != nil {
		return nil, fmt.Errorf("BroadcastRegisterPayload: %w", err)
	}

	return &BroadcastRegisterPayload{
		SenderPubKey: pubKey.B58Str(),
		CtrtMeta:     r.CtrtMeta.Serialize().B58Str(),
		InitData:     r.DataStack.Serialize().B58Str(),
		Description:  r.Description,
		Fee:          r.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    r.Timestamp,
		Signature:    sig.B58Str(),
	}, nil
}

func (r *RegCtrtTxReq) String() string {
	return fmt.Sprintf("%T(%+v)", r, *r)
}
