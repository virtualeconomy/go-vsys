package vsys

import (
	"crypto/sha256"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func Sha256Hash(b []byte) []byte {
	arr := sha256.Sum256(b)
	return arr[:]
}

func Keccak256Hash(b []byte) []byte {
	d := sha3.NewLegacyKeccak256()
	d.Write(b)
	return d.Sum(nil)
}

func Blake2bHash(b []byte) []byte {
	arr := blake2b.Sum256(b)
	return arr[:]
}
