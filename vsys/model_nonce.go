package vsys

import "fmt"

type Nonce uint

func (n Nonce) String() string {
	return fmt.Sprintf("%T(%d)", n, n)
}
