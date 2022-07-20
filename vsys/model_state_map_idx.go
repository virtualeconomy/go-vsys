package vsys

import "fmt"

type StateMapIdx uint8

func (s StateMapIdx) Serialize() Bytes {
	return PackUInt8(uint8(s))
}

func (s StateMapIdx) String() string {
	return fmt.Sprintf("%T(%d)", s, s)
}

// Constants for Atomic Swap Contract state map indexes.
const (
	STATE_MAP_IDX_ATOMIC_SWAP_CONTRACT_BALANCE = StateMapIdx(iota)
	STATE_MAP_IDX_ATOMIC_SWAP_OWNER
	STATE_MAP_IDX_ATOMIC_SWAP_RECIPIENT
	STATE_MAP_IDX_ATOMIC_SWAP_PUZZLE
	STATE_MAP_IDX_ATOMIC_SWAP_AMOUNT
	STATE_MAP_IDX_ATOMIC_SWAP_EXPIRED_TIME
	STATE_MAP_IDX_ATOMIC_SWAP_STATUS
)
