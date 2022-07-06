package vsys

import "fmt"

type VSYS uint64

const VSYS_UNIT = 1_00_000_000

func NewVSYSForAmount(amount float64) (VSYS, error) {
	data := amount * VSYS_UNIT

	// if data has non-zero fraction
	if float64(int(data)) < data {
		return 0, fmt.Errorf("NewVSYSForAmount: the minimal valid amount granularity is %f", 1/float64(VSYS_UNIT))
	}

	return VSYS(data), nil
}

func (v VSYS) Unit() uint64 {
	return VSYS_UNIT
}

func (v VSYS) Amount() float64 {
	return float64(v) / float64(v.Unit())
}

func (v VSYS) Uint64() uint64 {
	return uint64(v)
}

func (v VSYS) Serialize() Bytes {
	return PackUInt64(v.Uint64())
}

func (v VSYS) String() string {
	return fmt.Sprintf("%T(%d)", v, v)
}

const (
	FEE_DEFAULT        = VSYS(0.1 * VSYS_UNIT)
	FEE_PAYMENT        = VSYS(FEE_DEFAULT)
	FEE_LEASING        = VSYS(FEE_DEFAULT)
	FEE_LEASING_CANCEL = VSYS(FEE_DEFAULT)
	FEE_REG_CTRT       = VSYS(100 * VSYS_UNIT)
	FEE_EXEC_CTRT      = VSYS(0.3 * VSYS_UNIT)
	FEE_CONTEND_SLOTS  = VSYS(50_000 * VSYS_UNIT)
	FEE_DBPUT          = VSYS(VSYS_UNIT)
)
