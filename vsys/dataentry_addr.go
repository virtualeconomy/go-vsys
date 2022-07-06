package vsys

import (
	"fmt"
)

type DeAddr struct {
	Idx DeIdx

	Data *Addr
}

func NewDeAddr(a *Addr) *DeAddr {
	return &DeAddr{
		Idx:  2,
		Data: a,
	}
}

func NewDeAddrFromBytesGeneric(b []byte) (DataEntry, error) {
	a, err := NewAddr(b[1:])
	if err != nil {
		return nil, fmt.Errorf("NewAddrFromBytesGeneric: %w", err)
	}
	return NewDeAddr(a), nil
}

func NewDeAddrFromBytes(b []byte) (*DeAddr, error) {
	a, err := NewDeAddrFromBytesGeneric(b)
	if err != nil {
		return nil, fmt.Errorf("NewAddrFromBytes: %w", err)
	}
	return a.(*DeAddr), nil
}

func (a *DeAddr) IdxBytes() Bytes {
	return a.Idx.Serialize()
}

func (a *DeAddr) DataBytes() Bytes {
	return a.Data.Bytes
}

func (a *DeAddr) Serialize() Bytes {
	return append(a.IdxBytes(), a.DataBytes()...)
}

func (a *DeAddr) Size() int {
	return 1 + len(a.DataBytes())
}

func (a *DeAddr) String() string {
	return fmt.Sprintf("%T(%+v)", a, a.Data)
}
