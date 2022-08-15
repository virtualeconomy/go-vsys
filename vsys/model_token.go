package vsys

import "fmt"

type Token struct {
	Data Amount
	Unit Unit
}

func NewToken(data Amount, unit Unit) *Token {
	return &Token{data, unit}
}

func NewTokenForAmount(amount float64, unit uint64) (*Token, error) {
	data := amount * float64(unit)
	if float64(int(data)) < data {
		return nil, fmt.Errorf("NewTokenForAmount: The minimal valid granularity is %f", 1/float64(unit))
	}
	return &Token{Amount(data), Unit(unit)}, nil
}

func (t *Token) Amount() float64 {
	return float64(t.Data) / float64(t.Unit)
}

func (t *Token) DataUint64() uint64 {
	return uint64(t.Data)
}

func (t *Token) Serialize() Bytes {
	return PackUInt64(t.DataUint64())
}

func (t *Token) String() string {
	return fmt.Sprintf("%T(%+v)", t, *t)
}
