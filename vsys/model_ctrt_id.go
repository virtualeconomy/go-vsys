package vsys

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
)

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

func (c *CtrtId) GetTokId(tokIdx uint32) (*TokenId, error) {
	b := c.Bytes
	raw_CtrtId := b[1 : len(b)-CTRT_META_CHECKSUM_LEN]
	ctrtIdNoChecksum := append(append(PackUInt8(CTRT_META_TOKEN_ADDR_VER), raw_CtrtId...), PackUInt32(tokIdx)...)
	h := Keccak256Hash(Blake2bHash(ctrtIdNoChecksum))
	tokIdBytes := base58.Encode(append(ctrtIdNoChecksum, h[:CTRT_META_CHECKSUM_LEN]...))
	tokId := string(tokIdBytes)
	return NewTokenIdFromB58Str(tokId)
}

func (c *CtrtId) String() string {
	return fmt.Sprintf("%T(%s)", c, c.B58Str())
}
