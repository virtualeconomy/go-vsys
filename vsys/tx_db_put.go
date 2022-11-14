package vsys

import (
	"fmt"
)

type DbPutTxReq struct {
	TxType    TxType
	Key       DBPutKey
	Data      DBPutData
	Timestamp VSYSTimestamp
	Fee       VSYS
}

func NewDbPutTxReq(key DBPutKey, data DBPutData, ts VSYSTimestamp, fee VSYS) *DbPutTxReq {
	return &DbPutTxReq{
		TX_TYPE_DB_PUT,
		key,
		data,
		ts,
		fee,
	}
}

func (d *DbPutTxReq) DataToSign() Bytes {
	txTypeBytes := d.TxType.Serialize()
	feeBytes := d.Fee.Serialize()
	feeScaleBytes := FEE_SCALE_DEFAULT.Serialize()
	tsBytes := d.Timestamp.Serialize()
	keyBytes := d.Key.Serialize()
	dataBytes := d.Data.Serialize()

	size := len(txTypeBytes) +
		len(feeBytes) + len(feeScaleBytes) +
		len(tsBytes) + len(keyBytes) + len(dataBytes)

	b := make([]byte, 0, size)

	b = append(b, txTypeBytes...)
	b = append(b, keyBytes...)
	b = append(b, dataBytes...)
	b = append(b, feeBytes...)
	b = append(b, feeScaleBytes...)
	b = append(b, tsBytes...)

	return b
}

func (d *DbPutTxReq) BroadcastDbPutPayload(priKey *PriKey, pubKey *PubKey) (*BroadcastPutDbPayload, error) {
	sig, err := priKey.Sign(d.DataToSign())

	if err != nil {
		return nil, fmt.Errorf("BroadcastExecutePayload: %w", err)
	}

	return &BroadcastPutDbPayload{
		SenderPubKey: pubKey.B58Str(),
		DbKey:        Str(d.Key),
		DataType:     d.Data.GetDataType(),
		Data:         d.Data.Str(),
		Fee:          d.Fee,
		FeeScale:     FEE_SCALE_DEFAULT,
		Timestamp:    d.Timestamp,
		Signature:    sig.B58Str(),
	}, nil
}

func (d *DbPutTxReq) String() string {
	return fmt.Sprintf("%T(%+v)", d, *d)
}
