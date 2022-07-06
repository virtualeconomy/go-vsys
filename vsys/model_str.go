package vsys

import (
	"fmt"
	"unicode/utf8"
)

type Str string

func (s Str) Bytes() Bytes {
	return Bytes(s)
}

func (s Str) B58Bytes() (Bytes, error) {
	return NewBytesFromB58Str(string(s))
}

func (s Str) B58Str() string {
	return B58Encode([]byte(s))
}

func (s Str) Str() string {
	return string(s)
}

func (s Str) RuneLen() int {
	return utf8.RuneCountInString(string(s))
}

func (s Str) String() string {
	return fmt.Sprintf("%T(%s)", s, string(s))
}
