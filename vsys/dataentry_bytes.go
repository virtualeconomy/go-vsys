package vsys

import (
	"fmt"
)

type DeBytes struct {
	Idx DeIdx

	Data Bytes
}

func NewDeBytes(b Bytes) *DeBytes {
	return &DeBytes{
		Idx:  11,
		Data: b,
	}
}

func NewDeBytesFromBytesGeneric(b []byte) (DataEntry, error) {
	size, err := UnpackUInt16(b[1:3])
	if err != nil {
		return nil, fmt.Errorf("NewDeBytesFromBytesGeneric: %w", err)
	}
	return NewDeBytes(b[3 : 3+size]), nil
}

func (b *DeBytes) IdxBytes() Bytes {
	return b.Idx.Serialize()
}

func (b *DeBytes) DataBytes() Bytes {
	return b.Data
}

func (b *DeBytes) LenBytes() Bytes {
	return PackUInt16(uint16(len(b.DataBytes())))
}

func (b *DeBytes) Serialize() Bytes {
	size := len(b.IdxBytes()) +
		len(b.LenBytes()) +
		len(b.DataBytes())

	bs := make([]byte, 0, size)
	bs = append(bs, b.IdxBytes()...)
	bs = append(bs, b.LenBytes()...)
	bs = append(bs, b.DataBytes()...)

	return bs
}

func (b *DeBytes) Size() int {
	return 3 + len(b.DataBytes())
}

func (b *DeBytes) String() string {
	return fmt.Sprintf("%T(%+v)", b, b.Data)
}
