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

type BroadcastPutDbPayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	DbKey        Str           `json:"dbKey"`
	DataType     Str           `json:"dataType"`
	Data         Str           `json:"data"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Signature    Str           `json:"signature"`
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
