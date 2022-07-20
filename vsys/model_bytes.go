package vsys

import (
	"fmt"
)

type Bytes []byte

func NewBytesFromB58Str(s string) (Bytes, error) {
	b, err := B58Decode(s)
	if err != nil {
		return nil, fmt.Errorf("NewBytesFromB58Str: %w", err)
	}
	return Bytes(b), nil
}

// NewBytesFromStr converts string data type to Bytes type.
func NewBytesFromStr(s string) Bytes {
	return Bytes(s)
}

func (b Bytes) B58Str() Str {
	return Str(B58Encode(b))
}

func (b Bytes) Str() Str {
	return Str(b)
}

func (b Bytes) ByteSlice() []byte {
	return []byte(b)
}

func (b Bytes) Size() int {
	return len(b)
}

func (b Bytes) String() string {
	return fmt.Sprintf("%T(%s)", b, string(b))
}
