package vsys

import (
	"bytes"
	"fmt"
)

type AddrVer uint8

type Addr struct {
	Bytes
}

func NewAddr(b []byte) (*Addr, error) {
	if len(b) != ADDR_BYTES_LEN {
		return nil, fmt.Errorf("NewAddr: Addr must be %d bytes", ADDR_BYTES_LEN)
	}

	addr := &Addr{b}

	err := addr.ChainID().Validate()
	if err != nil {
		return nil, fmt.Errorf("NewAddr: %w", err)
	}

	checksum := Keccak256Hash(
		Blake2bHash(
			addr.NonChecksumBytes(),
		),
	)[:ADDR_CHECKSUM_BYTES_LEN]
	if bytes.Compare(checksum, addr.Checksum()) != 0 {
		return nil, fmt.Errorf("NewAddr: Addr has invalid checksum")
	}

	return addr, nil
}

func NewAddrFromB58Str(s string) (*Addr, error) {
	b, err := NewBytesFromB58Str(s)
	if err != nil {
		return nil, fmt.Errorf("NewAddrFromB58Str: %w", err)
	}
	addr, err := NewAddr(b)
	if err != nil {
		return nil, fmt.Errorf("NewAddrFromB58Str: %w", err)
	}
	return addr, nil
}

func NewAddrFromPubKey(pubKey *PubKey, chainID ChainID) (*Addr, error) {
	const ADDR_VER AddrVer = 5

	keBlaHash := func(b []byte) []byte {
		return Keccak256Hash(Blake2bHash(b))
	}

	pubKeyHash := string(keBlaHash(pubKey.Bytes))[:20]
	rawAddr := string(byte(ADDR_VER)) + string(chainID) + pubKeyHash

	checksum := string(
		keBlaHash(
			[]byte(rawAddr),
		),
	)[:4]

	b := []byte(rawAddr + checksum)
	addr, err := NewAddr(b)
	if err != nil {
		return nil, fmt.Errorf("NewAddrFromPubkey: %w", err)
	}
	return addr, nil
}

const (
	ADDR_VER_BYTES_LEN         = 1
	ADDR_CHAIN_ID_BYTES_LEN    = 1
	ADDR_PUBKEY_HASH_BYTES_LEN = 20
	ADDR_CHECKSUM_BYTES_LEN    = 4
	ADDR_BYTES_LEN             = ADDR_VER_BYTES_LEN + ADDR_CHAIN_ID_BYTES_LEN + ADDR_PUBKEY_HASH_BYTES_LEN + ADDR_CHECKSUM_BYTES_LEN
)

func (a *Addr) VerByte() byte {
	return a.Bytes[0]
}

func (a *Addr) Ver() AddrVer {
	return AddrVer(a.VerByte())
}

func (a *Addr) ChainIDByte() byte {
	return a.Bytes[1]
}

func (a *Addr) ChainID() ChainID {
	return ChainID(a.ChainIDByte())
}

func (a *Addr) PubKeyHash() Bytes {
	offset := ADDR_VER_BYTES_LEN + ADDR_CHAIN_ID_BYTES_LEN
	return a.Bytes[offset:][:ADDR_PUBKEY_HASH_BYTES_LEN]
}

func (a *Addr) NonChecksumBytes() Bytes {
	end := ADDR_VER_BYTES_LEN + ADDR_CHAIN_ID_BYTES_LEN + ADDR_PUBKEY_HASH_BYTES_LEN
	return a.Bytes[:end]
}

func (a *Addr) Checksum() Bytes {
	offset := ADDR_VER_BYTES_LEN + ADDR_CHAIN_ID_BYTES_LEN + ADDR_PUBKEY_HASH_BYTES_LEN
	return a.Bytes[offset:]
}

func (a *Addr) String() string {
	return fmt.Sprintf("%T(%s)", a, a.B58Str())
}
