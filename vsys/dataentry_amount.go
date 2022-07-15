package vsys

import "fmt"

type DeAmount struct {
	Idx DeIdx

	Data Amount
}

func NewDeAmount(a Amount) *DeAmount {
	return &DeAmount{
		Idx:  3,
		Data: a,
	}
}

func NewDeAmountForTokAmount(amount float64, unit uint64) (*DeAmount, error) {
	token, err := NewTokenForAmount(amount, unit)
	if err != nil {
		return nil, err
	}
	return NewDeAmount(Amount(token.Data)), nil
}

func (t *DeAmount) IdxBytes() Bytes {
	return t.Idx.Serialize()
}

func (t *DeAmount) DataBytes() Bytes {
	return t.Data.Serialize()
}

func (t *DeAmount) Serialize() Bytes {
	return append(t.IdxBytes(), t.DataBytes()...)
}

func (t *DeAmount) Size() int {
	return 1 + len(t.DataBytes())
}

func (t *DeAmount) String() string {
	return fmt.Sprintf("%T(%+v)", t, t.Data)
}
