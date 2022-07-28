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
	STATE_VAR_NFTV2_ISSUER = StateVar(iota)
	STATE_VAR_NFTV2_MAKER
	STATE_VAR_NFTV2_REGULATOR
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

const (
	STATE_VAR_V_ESCROW_MAKER = StateVar(iota)
	STATE_VAR_V_ESCROW_JUDGE
	STATE_VAR_V_ESCROW_TOKEN_ID
	STATE_VAR_V_ESCROW_DURATION
	STATE_VAR_V_ESCROW_JUDGE_DURATION
)
