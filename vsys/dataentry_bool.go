package vsys

import (
	"fmt"
)

type DeBool struct {
	Idx DeIdx

	Data bool
}

func NewDeBool(b bool) *DeBool {
	return &DeBool{
		Idx:  10,
		Data: b,
	}
}

func NewDeBoolFromBytesGeneric(b []byte) (DataEntry, error) {
	return NewDeBoolFromBytes(b[1:2])
}

func NewDeBoolFromBytes(b Bytes) (*DeBool, error) {
	v, err := UnpackBool(b)
	if err != nil {
		return nil, fmt.Errorf("NewDeBoolFromBytes: %w", err)
	}
	return NewDeBool(v), nil
}

func (b *DeBool) IdxBytes() Bytes {
	return b.Idx.Serialize()
}

func (b *DeBool) DataBytes() Bytes {
	return PackBool(b.Data)
}

func (b *DeBool) Serialize() Bytes {
	return append(b.IdxBytes(), b.DataBytes()...)
}

func (b *DeBool) Size() int {
	return 1 + len(b.DataBytes())
}

func (b *DeBool) String() string {
	return fmt.Sprintf("DeBool(%v)", b.Data)
}
