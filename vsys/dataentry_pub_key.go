package vsys

import (
	"fmt"
)

type DePubKey struct {
	Idx DeIdx

	Data *PubKey
}

func NewDePubKey(p *PubKey) *DePubKey {
	return &DePubKey{
		Idx:  1,
		Data: p,
	}
}

func NewDePubKeyFromBytesGeneric(b []byte) (DataEntry, error) {
	p, err := NewPubKey(b[1:])
	if err != nil {
		return nil, fmt.Errorf("NewDePubKeyFromBytesGeneric: %w", err)
	}
	return NewDePubKey(p), nil
}

func NewDePubKeyFromBytes(b []byte) (*DePubKey, error) {
	d, err := NewDePubKeyFromBytesGeneric(b)
	if err != nil {
		return nil, fmt.Errorf("NewDePubKeyFromBytes: %w", err)
	}
	return d.(*DePubKey), nil

}

func (p *DePubKey) IdxBytes() Bytes {
	return p.Idx.Serialize()
}

func (p *DePubKey) DataBytes() Bytes {
	return p.Data.Bytes
}

func (p *DePubKey) Serialize() Bytes {
	return append(p.IdxBytes(), p.DataBytes()...)
}

func (p *DePubKey) Size() int {
	return 1 + len(p.DataBytes())
}

func (p *DePubKey) String() string {
	return fmt.Sprintf("%T(%+v)", p, p.Data)
}
