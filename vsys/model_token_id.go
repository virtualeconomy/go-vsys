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
	// TODO: refactor byte length and link to CtrtID or CtrtMeta
	return t.Bytes[:TOKEN_ID_BYTES_LEN-4]
}

func (t *TokenId) Checksum() Bytes {
	return t.Bytes[26:]
}

func (t *TokenId) VerByte() byte {
	return t.Bytes[0]
}

func (t *TokenId) Ver() TokenIdVer {
	return TokenIdVer(t.VerByte())
}

func (t *TokenId) String() string {
	return fmt.Sprintf("%T(%s)", t, t.B58Str())
}
