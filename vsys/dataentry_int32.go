package vsys

type DeInt32 struct {
	Idx DeIdx

	Data uint32
}

func NewDeInt32(i uint32) *DeInt32 {
	return &DeInt32{
		Idx:  4,
		Data: i,
	}
}

func (i DeInt32) IdxBytes() Bytes {
	return i.Idx.Serialize()
}

func (i DeInt32) DataBytes() Bytes {
	return PackUInt32(i.Data)
}

func (i DeInt32) Serialize() Bytes {
	return append(i.IdxBytes(), i.DataBytes()...)
}

func (i DeInt32) Size() int {
	return 1 + len(i.DataBytes())
}
