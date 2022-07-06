package vsys

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
)

func B58Encode(b []byte) string {
	return base58.Encode(b)
}

func B58Decode(s string) ([]byte, error) {
	b := base58.Decode(s)
	if len(b) == 0 {
		return nil, fmt.Errorf("B58Decode: failed to decode %s", s)
	}
	return b, nil
}
