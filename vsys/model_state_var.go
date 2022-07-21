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
	STATE_VAR_NFT_MAKER
)

const (
	STATE_VAR_ATOMIC_SWAP_MAKER = StateVar(iota)
	STATE_VAR_ATOMIC_SWAP_TOKEN_ID
)

const (
	STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_ISSUER = StateVar(iota)
	STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_MAKER
)

const (
	STATE_VAR_TOK_CTRT_V2_ISSUER = StateVar(iota)
	STATE_VAR_TOK_CTRT_V2_MAKER
	STATE_VAR_TOK_CTRT_V2_REGULATOR
)
