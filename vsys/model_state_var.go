package vsys

import "fmt"

type StateVar uint8

func (s StateVar) Serialize() Bytes {
	return PackUInt8(uint8(s))
}

func (s StateVar) String() string {
	return fmt.Sprintf("%T(%d)", s, s)
}

const (
	STATE_VAR_NFT_ISSUER = StateVar(iota)
	STATE_VAR_NFT_MAKER  = StateVar(iota)
)
