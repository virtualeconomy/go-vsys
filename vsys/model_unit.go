package vsys

import "fmt"

type Unit uint64

func (u Unit) String() string {
	return fmt.Sprintf("%T(%d)", u, u)
}
