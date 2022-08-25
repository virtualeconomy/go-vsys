package vsys

import (
	"fmt"
)

type DeTokenId struct {
	Idx DeIdx

	Data *TokenId
}

func NewDeTokenId(t *TokenId) *DeTokenId {
	return &DeTokenId{
		Idx:  8,
		Data: t,
	}
}

func NewDeTokenIdFromBytesGeneric(b []byte) (DataEntry, error) {
	a := B58Encode(b[1 : 1+30])
	t, err := NewTokenId([]byte(a))
	if err != nil {
		return nil, fmt.Errorf("NewDeAcntFromBytesGeneric: %w", err)
	}
	return NewDeTokenId(t), nil
}

func (d *DeTokenId) IdxBytes() Bytes {
	return d.Idx.Serialize()
}

func (d *DeTokenId) DataBytes() Bytes {
	return d.Data.Bytes
}

func (d *DeTokenId) Serialize() Bytes {
	return append(d.IdxBytes(), d.DataBytes()...)
}

func (d *DeTokenId) Size() int {
	return 1 + len(d.DataBytes())
}
