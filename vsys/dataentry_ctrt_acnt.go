package vsys

import "fmt"

type DeCtrtAcnt struct {
	Idx DeIdx

	Data Bytes
}

func NewDeCtrtAcntFromBytesGeneric(b []byte) (DataEntry, error) {
	a := B58Encode(b[1 : 1+26])
	return NewDeCtrtAcnt(Bytes(a)), nil
}

func (a *DeCtrtAcnt) IdxBytes() Bytes {
	return a.Idx.Serialize()
}

func (a *DeCtrtAcnt) DataBytes() Bytes {
	return a.Data
}

func (a *DeCtrtAcnt) Serialize() Bytes {
	return append(a.IdxBytes(), a.DataBytes()...)
}

func (a *DeCtrtAcnt) Size() int {
	return 1 + len(a.DataBytes())
}

func (a *DeCtrtAcnt) String() string {
	return fmt.Sprintf("%T(%+v)", a, a.Data)
}

func NewDeCtrtAcnt(d Bytes) *DeCtrtAcnt {
	return &DeCtrtAcnt{
		Idx:  6,
		Data: d,
	}
}

func NewDeCtrtAddrFromCtrtId(c *CtrtId) *DeCtrtAcnt {
	return NewDeCtrtAcnt(c.Bytes)
}
