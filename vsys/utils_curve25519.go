package vsys

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

func GenPriKey(rand []byte) []byte {
	priKey := make([]byte, 32)
	copy(priKey, rand)

	priKey[0] &= 248
	priKey[31] &= 127
	priKey[31] |= 64

	return priKey
}

func GenPubKey(priKey []byte) ([]byte, error) {
	if len(priKey) != 32 {
		return nil, fmt.Errorf("GenPubKeyFromPriKey: the length of priKey must be 32 bytes")
	}

	pubKey := new([32]byte)
	curve25519.ScalarBaseMult(pubKey, (*[32]byte)(priKey))

	pubKey[31] &= 127

	return pubKey[:], nil
}

func RandBytes(size int) ([]byte, error) {
	r := make([]byte, size)

	_, err := rand.Read(r)
	if err != nil {
		return nil, fmt.Errorf("RandBytes: %w", err)
	}

	return r, nil
}

func Sign(priKey, msg []byte) ([]byte, error) {
	rand, err := RandBytes(64)
	if err != nil {
		return nil, fmt.Errorf("Sign: %w", err)
	}
	return SignImpl(priKey, msg, rand), nil
}

func Verify(pubKey, msg, sig []byte) bool {
	res := VerifyImpl(pubKey, msg, sig)
	return res == 1
}
