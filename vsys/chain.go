package vsys

import (
	"fmt"
)

type Chain struct {
	NodeAPI *NodeAPI
	ChainID ChainID
}

func NewChain(nodeAPI *NodeAPI, chainID ChainID) *Chain {
	return &Chain{NodeAPI: nodeAPI, ChainID: chainID}
}

func (c *Chain) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
}

func (c *Chain) Height() (Height, error) {
	res, err := c.NodeAPI.GetCurBlockHeight()
	if err != nil {
		return 0, fmt.Errorf("Height: %w", err)
	}
	return res.Height, nil
}
