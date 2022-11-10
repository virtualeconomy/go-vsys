package vsys

import (
	"fmt"
)

type GetDBResponse struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func (g *GetDBResponse) String() string {
	return fmt.Sprintf("%T(%+v)", g, *g)
}

// GetDB get broadcasts the DB Put request.
func (na *NodeAPI) GetDB(addr string, key string) (*GetDBResponse, error) {
	res := &GetDBResponse{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/database/get/%s/%s", addr, key))
	if err != nil {
		return nil, fmt.Errorf("GetDB: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetDB: %s", resp.String())
	}
	return res, nil
}

type DBPutKey Str

func (s DBPutKey) Serialize() Bytes {
	l := len(s)
	return append(PackUInt16(uint16(l)), []byte(s)...)
}

type DBPutData interface {
	Serialize() Bytes
	Str() Str
	GetDataType() Str
}

const (
	DB_PUT_DATA_ID = iota
	DB_PUT_DATA_BYTE_ARRAY
)

type ByteArray struct {
	data string
}

func NewDbPutByteArray(d string) *ByteArray {
	return &ByteArray{d}
}

func (ba *ByteArray) IdBytes() Bytes {
	return PackUInt8(DB_PUT_DATA_BYTE_ARRAY)
}

func (ba *ByteArray) DataBytes() Bytes {
	return NewBytesFromStr(ba.data)
}

func (ba *ByteArray) Serialize() Bytes {
	l := len(ba.data) + 1
	b := PackUInt16(uint16(l))
	b = append(b, ba.IdBytes()...)
	b = append(b, ba.DataBytes()...)
	return b
}

func (ba *ByteArray) Str() Str {
	return Str(ba.data)
}

func (ba *ByteArray) GetDataType() Str {
	return "ByteArray"
}

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

type BroadcastPutDbPayload struct {
	SenderPubKey Str `json:"senderPublicKey"`
	DbKey        Str `json:"dbKey"`
	DataType     Str `json:"dataType"`
	Data         Str `json:"data"`

	Fee       VSYS          `json:"fee"`
	FeeScale  FeeScale      `json:"feeScale"`
	Timestamp VSYSTimestamp `json:"timestamp"`
	Signature Str           `json:"signature"`
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

type BroadcastPutDbTxResp struct {
	TxBasic

	DbKey Str `json:"dbKey"`
	Entry struct {
		Data Str `json:"data"`
		Type Str `json:"type"`
	} `json:"entry"`
}

func (b *BroadcastPutDbTxResp) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}

// BroadcastDbPut broadcasts the DB Put request.
func (na *NodeAPI) BroadcastDbPut(p *BroadcastPutDbPayload) (*BroadcastPutDbTxResp, error) {
	res := &BroadcastPutDbTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/database/broadcast/put")

	if err != nil {
		return nil, fmt.Errorf("BroadcastDbPut: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastDbPut: %s", resp.String())
	}
	return res, nil
}
