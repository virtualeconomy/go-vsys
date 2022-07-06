package vsys

import (
	"fmt"
)

type PriKey struct {
	Bytes
}

func NewPriKey(b []byte) (*PriKey, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("NewPriKey: Prikey must be 32 bytes")
	}

	return &PriKey{Bytes(b)}, nil
}

func NewPriKeyFromB58Str(s string) (*PriKey, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewPriKeyFromB58Str: %w", err)
	}

	pk, err := NewPriKey(b)
	if err != nil {
		return nil, fmt.Errorf("NewPriKeyFromB58Str: %w", err)
	}

	return pk, nil
}

func NewPriKeyFromRand(rand []byte) (*PriKey, error) {
	pkBytes := GenPriKey(rand)
	pk, err := NewPriKey(pkBytes)
	if err != nil {
		return nil, fmt.Errorf("NewPriKeyFromRand: %w", err)
	}
	return pk, nil
}

func (p *PriKey) Sign(msg []byte) (Bytes, error) {
	sig, err := Sign(p.Bytes, msg)

	if err != nil {
		return nil, fmt.Errorf("Sign: %w", err)
	}
	return sig, nil
}

func (p *PriKey) String() string {
	return fmt.Sprintf("%T(%s)", p, p.B58Str())
}
