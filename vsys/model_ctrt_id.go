package vsys

import "fmt"

type CtrtId struct {
	Bytes
}

func NewCtrtId(b []byte) (*CtrtId, error) {
	if len(b) != 26 {
		return nil, fmt.Errorf("NewCtrtId: CtrtId must be 26 bytes")
	}

	return &CtrtId{Bytes(b)}, nil
}

func NewCtrtIdFromB58Str(s string) (*CtrtId, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtIdFromB58Str: %w", err)
	}

	cid, err := NewCtrtId(b)
	if err != nil {
		return nil, fmt.Errorf("NewCtrtIdFromB58Str: %w", err)
	}

	return cid, nil
}

func (c *CtrtId) String() string {
	return fmt.Sprintf("%T(%s)", c, c.B58Str())
}
