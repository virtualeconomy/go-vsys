package vsys

import (
	"fmt"
)

type ExecCtrtFuncTxReq struct {
	TxType TxType

	CtrtId     *CtrtId
	FuncIdx    FuncIdx
	DataStack  DataStack
	Timestamp  VSYSTimestamp
	Attachment Str
	Fee        VSYS
}

func NewExecCtrtFuncTxReq(
	ctrtId *CtrtId,
	funcIdx FuncIdx,
	dataStack DataStack,
	timestamp VSYSTimestamp,
	attachment Str,
	fee VSYS,
) *ExecCtrtFuncTxReq {
	return &ExecCtrtFuncTxReq{
		TxType:     TX_TYPE_EXECUTE_CONTRACT_FUNCTION,
		CtrtId:     ctrtId,
		FuncIdx:    funcIdx,
		DataStack:  dataStack,
		Timestamp:  timestamp,
		Attachment: attachment,
		Fee:        fee,
	}
}

func (e *ExecCtrtFuncTxReq) DataToSign() Bytes {
	txTypeBytes := e.TxType.Serialize()
	ctrtIdBytes := e.CtrtId.Bytes
	funcIdxBytes := e.FuncIdx.Serialize()

	dataStackBytes := e.DataStack.Serialize()
	dataStackLenBytes := PackUInt16(uint16(len(dataStackBytes)))

	attachmentLenBytes := PackUInt16(uint16(e.Attachment.RuneLen()))
	attachmentBytes := e.Attachment.Bytes()

	feeBytes := e.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	tsBytes := e.Timestamp.Serialize()

	size := len(txTypeBytes) +
		len(ctrtIdBytes) +
		len(funcIdxBytes) +
		len(dataStackLenBytes) +
		len(dataStackBytes) +
		len(attachmentLenBytes) +
		len(attachmentBytes) +
		len(feeBytes) +
		len(feeScaleBytes) +
		len(tsBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, ctrtIdBytes...)
	b = append(b, funcIdxBytes...)
	b = append(b, dataStackLenBytes...)
	b = append(b, dataStackBytes...)
	b = append(b, attachmentLenBytes...)
	b = append(b, attachmentBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, tsBytes...)

	return b
}

func (e *ExecCtrtFuncTxReq) BroadcastExecutePayload(priKey *PriKey, pubKey *PubKey) (*BroadcastExecutePayload, error) {
	sig, err := priKey.Sign(e.DataToSign())

	if err != nil {
		return nil, fmt.Errorf("BroadcastExecutePayload: %w", err)
	}

	return &BroadcastExecutePayload{
		SenderPubKey: pubKey.B58Str(),
		CtrtId:       e.CtrtId.B58Str(),
		FuncIdx:      e.FuncIdx,
		FuncData:     e.DataStack.Serialize().B58Str(),
		Attachment:   e.Attachment,
		Fee:          e.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    e.Timestamp,
		Signature:    sig.B58Str(),
	}, nil
}

func (e *ExecCtrtFuncTxReq) String() string {
	return fmt.Sprintf("%T(%+v)", e, *e)
}
