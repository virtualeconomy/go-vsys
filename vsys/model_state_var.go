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

const (
	STATE_VAR_V_SWAP_MAKER = StateVar(iota)
	STATE_VAR_V_SWAP_TOKEN_A_ID
	STATE_VAR_V_SWAP_TOKEN_B_ID
	STATE_VAR_V_SWAP_LIQUIDITY_TOKEN_ID
	STATE_VAR_V_SWAP_SWAP_STATUS
	STATE_VAR_V_SWAP_MINIMUM_LIQUIDITY
	STATE_VAR_V_SWAP_TOKEN_A_RESERVED
	STATE_VAR_V_SWAP_TOKEN_B_RESERVED
	STATE_VAR_V_SWAP_TOTAL_SUPPLY
	STATE_VAR_V_SWAP_LIQUIDITY_TOKEN_LEFT
)

const (
	STATE_VAR_V_STABLE_SWAP_MAKER = StateVar(iota)
	STATE_VAR_V_STABLE_SWAP_BASE_TOKEN_ID
	STATE_VAR_V_STABLE_SWAP_TARGET_TOKEN_ID
	STATE_VAR_V_STABLE_SWAP_MAX_ORDER_PER_USER
	STATE_VAR_V_STABLE_SWAP_UNIT_PRICE_BASE
	STATE_VAR_V_STABLE_SWAP_UNIT_PRICE_TARGET
)
