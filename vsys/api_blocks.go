package vsys

import (
	"fmt"
)

// BlockHeightResp is the result of GET /blocks/height
type CurBlockHeightResp struct {
	Height Height `json:"height"`
}

// GetCurBlockHeight gets the current block height of chain.
func (na *NodeAPI) GetCurBlockHeight() (*CurBlockHeightResp, error) {
	res := &CurBlockHeightResp{}
	resp, err := na.R().SetResult(res).Get("/blocks/height")
	if err != nil {
		return nil, fmt.Errorf("GetCurBlockHeight: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetCurBlockHeight: %s", resp.String())
	}
	return res, nil
}
