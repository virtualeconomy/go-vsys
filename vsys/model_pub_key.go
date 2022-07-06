package vsys

import (
	"fmt"
)

type PubKey struct {
	Bytes
}

func NewPubKey(b []byte) (*PubKey, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("NewPubKey: PubKey must be 32 bytes")
	}

	return &PubKey{Bytes(b)}, nil
}

func NewPubKeyFromB58Str(s string) (*PubKey, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewPubKeyFromB58Str: %w", err)
	}

	return NewPubKey(b)
}

func NewPubKeyFromPriKey(p *PriKey) (*PubKey, error) {
	pubBytes, err := GenPubKey(p.Bytes)
	if err != nil {
		return nil, fmt.Errorf("NewPubKeyFromPriKey: %w", err)
	}

	pub, err := NewPubKey(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("NewPubKeyFromPriKey: %w", err)
	}

	return pub, nil
}

func (p *PubKey) Verify(msg, sig []byte) (isValid bool) {
	isValid = Verify(p.Bytes, msg, sig)
	return
}

func (p *PubKey) String() string {
	return fmt.Sprintf("%T(%s)", p, p.B58Str())
}
