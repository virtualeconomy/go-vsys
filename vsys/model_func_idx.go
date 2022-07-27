package vsys

import "fmt"

type FuncIdx uint16

func (f FuncIdx) Serialize() Bytes {
	return PackUInt16(uint16(f))
}

func (f FuncIdx) String() string {
	return fmt.Sprintf("%T(%d)", f, f)
}

const (
	FUNC_IDX_NFT_SUPERSEDE = FuncIdx(iota)
	FUNC_IDX_NFT_ISSUE
	FUNC_IDX_NFT_SEND
	FUNC_IDX_NFT_TRANSFER
	FUNC_IDX_NFT_DEPOSIT
	FUNC_IDX_NFT_WITHDRAW
)

const (
	FUNC_IDX_NFTV2_SUPERSEDE = FuncIdx(iota)
	FUNC_IDX_NFTV2_ISSUE
	FUNC_IDX_NFTV2_UPDATE_LIST
	FUNC_IDX_NFTV2_SEND
	FUNC_IDX_NFTV2_TRANSFER
	FUNC_IDX_NFTV2_DEPOSIT
	FUNC_IDX_NFTV2_WITHDRAW
)

const (
	FUNC_IDX_ATOMIC_SWAP_LOCK = FuncIdx(iota)
	FUNC_IDX_ATOMIC_SWAP_SOLVE_PUZZLE
	FUNC_IDX_ATOMIC_SWAP_EXPIRE_WITHDRAW
)

const (
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SUPERSEDE = FuncIdx(iota)
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_ISSUE
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DESTROY
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SEND
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_TRANSFER
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DEPOSIT
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_WITHDRAW
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_TOTAL_SUPPLY
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_MAX_SUPPLY
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_BALANCE_OF
	FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_GET_ISSUER
)

const (
	FUNC_IDX_V_ESCROW_CTRT_SUPERSEDE = FuncIdx(iota)
	FUNC_IDX_V_ESCROW_CTRT_CREATE
	FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_DEPOSIT
	FUNC_IDX_V_ESCROW_CTRT_JUDGE_DEPOSIT
	FUNC_IDX_V_ESCROW_CTRT_PAYER_CANCEL
	FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_CANCEL
	FUNC_IDX_V_ESCROW_CTRT_JUDGE_CANCEL
	FUNC_IDX_V_ESCROW_CTRT_SUBMIT_WORK
	FUNC_IDX_V_ESCROW_CTRT_APPROVE_WORK
	FUNC_IDX_V_ESCROW_CTRT_APPLY_TO_JUDGE
	FUNC_IDX_V_ESCROW_CTRT_JUDGE
	FUNC_IDX_V_ESCROW_CTRT_SUBMIT_PENALTY
	FUNC_IDX_V_ESCROW_CTRT_PAYER_REFUND
	FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_REFUND
	FUNC_IDX_V_ESCROW_CTRT_COLLECT
)
