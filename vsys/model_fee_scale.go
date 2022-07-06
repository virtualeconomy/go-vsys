package vsys

import "fmt"

type FeeScale uint16

func (f FeeScale) Uint16() uint16 {
	return uint16(f)
}

func (f FeeScale) String() string {
	return fmt.Sprintf("%T(%d)", f, f)
}

func (f FeeScale) Serialize() Bytes {
	return PackUInt16(f.Uint16())
}

const FEE_SCALE_DEFAULT = FeeScale(100)
