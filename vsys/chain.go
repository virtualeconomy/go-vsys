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

func (c *Chain) LastBlock() (*BlockResp, error) {
	res, err := c.NodeAPI.GetLast()
	if err != nil {
		return nil, fmt.Errorf("LastBlock: %w", err)
	}
	return res, nil
}

func (c *Chain) GetBlockAt(h int) (*BlockResp, error) {
	res, err := c.NodeAPI.GetBlockAt(h)
	if err != nil {
		return nil, fmt.Errorf("GetBlockAt: %w", err)
	}
	return res, nil
}

func (c *Chain) GetBlocksWithin(start, end int) ([]*BlockResp, error) {
	res, err := c.NodeAPI.GetBlocksWithin(start, end)
	if err != nil {
		return nil, fmt.Errorf("GetBlocksWithin: %w", err)
	}
	return res, nil
}
