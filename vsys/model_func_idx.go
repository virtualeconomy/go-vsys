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
	FUNC_IDX_ATOMIC_SWAP_LOCK = FuncIdx(iota)
	FUNC_IDX_ATOMIC_SWAP_SOLVE_PUZZLE
	FUNC_IDX_ATOMIC_SWAP_EXPIRE_WITHDRAW
)
