package vsys

import "fmt"

type Height uint64

func (h Height) String() string {
	return fmt.Sprintf("%T(%d)", h, h)
}
