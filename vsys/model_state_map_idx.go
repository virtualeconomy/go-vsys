package vsys

import "fmt"

type StateMapIdx uint8

func (s StateMapIdx) Serialize() Bytes {
	return PackUInt8(uint8(s))
}

func (s StateMapIdx) String() string {
	return fmt.Sprintf("%T(%d)", s, s)
}
