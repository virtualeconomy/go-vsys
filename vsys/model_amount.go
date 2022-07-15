package vsys

import "fmt"

type Amount uint64

func (a Amount) String() string {
	return fmt.Sprintf("%T(%d)", a, a)
}

func (a Amount) Serialize() Bytes {
	return PackUInt64(uint64(a))
}
