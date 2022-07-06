package vsys

import "fmt"

type TxType uint8

const (
	TX_TYPE_GENESIS = TxType(iota + 1)
	TX_TYPE_PAYMENT
	TX_TYPE_LEASE
	TX_TYPE_LEASE_CANCEL
	TX_TYPE_MINTING
	TX_TYPE_CONTENT_SLOTS
	TX_TYPE_RELEASE_SLOTS
	TX_TYPE_REGISTER_CONTRACT
	TX_TYPE_EXECUTE_CONTRACT_FUNCTION
	TX_TYPE_DB_PUT
)

func (t TxType) Serialize() Bytes {
	return PackUInt8(uint8(t))
}

func (t TxType) String() string {
	return fmt.Sprintf("%T(%d)", t, t)
}
