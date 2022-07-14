package vsys

import "fmt"

type DeAmount struct {
	Idx DeIdx
	// use model token instead model amount
	Data Token
}

func NewDeAmount(t Token) *DeAmount {
	return &DeAmount{
		Idx:  3,
		Data: t,
	}
}

func NewDeAmountForTokAmount(amount float64, unit uint64) (*DeAmount, error) {
	token, err := NewTokenForAmount(amount, unit)
	if err != nil {
		return nil, err
	}
	return NewDeAmount(*token), nil
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
