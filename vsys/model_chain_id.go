package vsys

import "fmt"

type ChainID string

const (
	MAIN_NET ChainID = "M"
	TEST_NET ChainID = "T"
)

func (c ChainID) Validate() error {
	valid := c == MAIN_NET || c == TEST_NET
	if !valid {
		return fmt.Errorf("Validate: ChainID must be either 'M' or 'T'")
	}
	return nil
}

func (c ChainID) String() string {
	return fmt.Sprintf("%T(%s)", c, string(c))
}
