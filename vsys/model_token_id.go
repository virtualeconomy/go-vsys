package vsys

import (
	"bytes"
	"fmt"
)

type TokenIdVer uint8

type TokenId struct {
	Bytes
}

func NewTokenId(b []byte) (*TokenId, error) {
	if len(b) != TOKEN_ID_BYTES_LEN {
		return nil, fmt.Errorf("NewTokenId: TokenId must be %d bytes", TOKEN_ID_BYTES_LEN)
	}

	tokId := &TokenId{b}

	checksum := Keccak256Hash(
		Blake2bHash(
			tokId.NonChecksumBytes(),
		),
	)[:TOKEN_ID_CHECKSUM_BYTES_LEN]
	if bytes.Compare(checksum, tokId.Checksum()) != 0 { //nolint:gosimple
		return nil, fmt.Errorf("NewTokenId: TokenId has invalid checksum")
	}

	return tokId, nil
}

const (
	TOKEN_ID_BYTES_LEN          = 30
	TOKEN_ID_CHECKSUM_BYTES_LEN = 4
)

func NewTokenIdFromB58Str(s string) (*TokenId, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewTokenIdFromB58Str: %w", err)
	}
	tokId, err := NewTokenId(b)
	if err != nil {
		return nil, fmt.Errorf("NewTokenIdFromB58Str: %w", err)
	}
	return tokId, nil
}

func (t *TokenId) NonChecksumBytes() Bytes {
	return t.Bytes[:TOKEN_ID_BYTES_LEN-TOKEN_ID_CHECKSUM_BYTES_LEN]
}

func (t *TokenId) Checksum() Bytes {
	return t.Bytes[TOKEN_ID_BYTES_LEN-TOKEN_ID_CHECKSUM_BYTES_LEN:]
}

func (t *TokenId) VerByte() byte {
	return t.Bytes[0]
}

func (t *TokenId) Ver() TokenIdVer {
	return TokenIdVer(t.VerByte())
}

// GetCtrtId computes the contract ID from token ID.
func (t *TokenId) GetCtrtId() (*CtrtId, error) {
	var b = []byte(t.Bytes)
	rawCtrtId := b[1:(len(b) - CTRT_META_TOKEN_IDX_BYTES_LEN - CTRT_META_CHECKSUM_LEN)]
	ctrtIdNoChecksum := append(PackUInt8(CTRT_META_CTRT_ADDR_VER), rawCtrtId...)

	h := Keccak256Hash(Blake2bHash(ctrtIdNoChecksum))

	ctrtIdStr := B58Encode(append(ctrtIdNoChecksum, h[:CTRT_META_CHECKSUM_LEN]...))

	ctrtId, err := NewCtrtIdFromB58Str(ctrtIdStr)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtId: %w", err)
	}
	return ctrtId, nil
}

func (t *TokenId) String() string {
	return fmt.Sprintf("%T(%s)", t, t.B58Str())
}
