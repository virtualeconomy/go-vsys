package vsys

type DataEntry interface {
	IdxBytes() Bytes

	DataBytes() Bytes
	Serialize() Bytes

	Size() int
}

type DeIdx uint8

func (d DeIdx) Serialize() Bytes {
	return PackUInt8(uint8(d))
}
