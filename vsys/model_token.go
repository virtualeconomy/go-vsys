package vsys

import "fmt"

type Token struct {
	Data uint64
	Unit Unit
}

// TODO: Refactor carefully and add methods.
func NewToken(data uint64, unit Unit) *Token {
	return &Token{data, unit}
}

func NewTokenForAmount(amount float64, unit uint64) (*Token, error) {
	data := amount * float64(unit)
	if float64(int(data)) < data {
		return nil, fmt.Errorf("NewTokenForAmount: The minimal valid granularity is %f", 1/unit)
	}
	return &Token{uint64(data), Unit(unit)}, nil
}

func (t *Token) Amount() float64 {
	return float64(t.Data) / float64(t.Unit)
}

func (t *Token) DataUint64() uint64 {
	return t.Data
}

func (t *Token) Serialize() Bytes {
	return PackUInt64(t.DataUint64())
}

func (t *Token) String() string {
	return fmt.Sprintf("%T(%+v)", t, *t)
}
