package vsys

import (
	"fmt"
	"time"
)

type VSYSTimestamp uint64

func NewVSYSTimestampForNow() VSYSTimestamp {
	nsec := time.Now().UnixNano()
	return VSYSTimestamp(nsec)
}

func NewVSYSTimestampFromUnixTs(sec int64) VSYSTimestamp {
	nsec := time.Unix(sec, 0).UnixNano()
	return VSYSTimestamp(nsec)
}

func (v VSYSTimestamp) UnixTs() int64 {
	return time.Unix(0, int64(v)).Unix()
}

func (v VSYSTimestamp) Uint64() uint64 {
	return uint64(v)
}

func (v VSYSTimestamp) Serialize() Bytes {
	return PackUInt64(v.Uint64())
}

func (v VSYSTimestamp) String() string {
	return fmt.Sprintf("%T(%d)", v, v)
}
