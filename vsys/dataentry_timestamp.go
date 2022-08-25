package vsys

import "fmt"

type DeTimestamp struct {
	Idx DeIdx

	Data VSYSTimestamp
}

func NewDeTimestamp(t VSYSTimestamp) *DeTimestamp {
	return &DeTimestamp{
		Idx:  9,
		Data: t,
	}
}

func NewDeTimestampFromBytesGeneric(b []byte) (DataEntry, error) {
	i, err := UnpackUInt64(b[1 : 1+8])
	if err != nil {
		return nil, fmt.Errorf("NewDeTimestampFromBytesGeneric: %w", err)
	}
	return NewDeTimestamp(VSYSTimestamp(i)), nil
}

func NewDeTimestampForNow() *DeTimestamp {
	return NewDeTimestamp(NewVSYSTimestampForNow())
}

func (t *DeTimestamp) IdxBytes() Bytes {
	return t.Idx.Serialize()
}

func (t *DeTimestamp) DataBytes() Bytes {
	return t.Data.Serialize()
}

func (t *DeTimestamp) Serialize() Bytes {
	return append(t.IdxBytes(), t.DataBytes()...)
}

func (t *DeTimestamp) Size() int {
	return 1 + len(t.DataBytes())
}

func (t *DeTimestamp) String() string {
	return fmt.Sprintf("%T(%+v)", t, t.Data)
}
