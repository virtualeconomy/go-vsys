package vsys

import (
	"encoding/json"
	"fmt"
)

// BlockHeightResp is the result of GETting height.
type BlockHeightResp struct {
	Height Height `json:"height"`
}

// GetCurBlockHeight gets the current block height of chain.
func (na *NodeAPI) GetCurBlockHeight() (*BlockHeightResp, error) {
	res := &BlockHeightResp{}
	resp, err := na.R().SetResult(res).Get("/blocks/height")
	if err != nil {
		return nil, fmt.Errorf("GetCurBlockHeight: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetCurBlockHeight: %s", resp.String())
	}
	return res, nil
}

// GetHeightBySignature gets the height of a block as per its signature.
func (na *NodeAPI) GetHeightBySignature(sig string) (*BlockHeightResp, error) {
	res := &BlockHeightResp{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/blocks/height/%s", sig))
	if err != nil {
		return nil, fmt.Errorf("GetHeightBySignature: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetHeightBySignature: %s", resp.String())
	}
	return res, nil
}

// BlockResp is a struct representing response of GETting block or array of blocks.
type BlockResp struct {
	Version       int           `json:"version"`
	Timestamp     VSYSTimestamp `json:"timestamp"`
	Reference     Str           `json:"reference"`
	SPOSConsensus struct {
		MintTime    VSYSTimestamp `json:"mintTime"`
		MintBalance VSYS          `json:"mintBalance"`
	} `json:"SPOSConsensus"`
	ResourcePricingData struct {
		Computation  int `json:"computation"`
		Storage      int `json:"storage"`
		Memory       int `json:"memory"`
		RandomIO     int `json:"randomIO"`
		SequentialIO int `json:"sequentialIO"`
	} `json:"resourcePricingData"`
	TransactionMerkleRoot Str `json:"TransactionMerkleRoot"`
	Transactions          []struct {
		Type               TxType        `json:"type"`
		Id                 Str           `json:"id"`
		Recipient          Str           `json:"recipient"`
		Timestamp          VSYSTimestamp `json:"timestamp"`
		Amount             Amount        `json:"amount"`
		CurrentBlockHeight Height        `json:"currentBlockHeight"`
		Status             Str           `json:"status"`
		FeeCharged         VSYS          `json:"feeCharged"`
	} `json:"transactions"`
	Generator        Str    `json:"generator"`
	Signature        Str    `json:"signature"`
	Fee              VSYS   `json:"fee"`
	Blocksize        int    `json:"blocksize"`
	Height           Height `json:"height"`
	TransactionCount int    `json:"transaction count"`
}

// GetLast gets the last block of the chain.
func (na *NodeAPI) GetLast() (*BlockResp, error) {
	res := &BlockResp{}
	resp, err := na.R().SetResult(res).Get("/blocks/last")
	if err != nil {
		return nil, fmt.Errorf("GetCurBlockHeight: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetCurBlockHeight: %s", resp.String())
	}
	return res, nil
}

// GetBlockAt gets the block at the given height.
func (na *NodeAPI) GetBlockAt(h int) (*BlockResp, error) {
	res := &BlockResp{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/blocks/at/%d", h))
	if err != nil {
		return nil, fmt.Errorf("GetBlockAt: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetBlockAt: %s", resp.String())
	}
	return res, nil
}

// GetBlocksWithin gets blocks fall in the given range.
func (na *NodeAPI) GetBlocksWithin(start, end int) ([]*BlockResp, error) {
	res := make([]*BlockResp, 0)
	resp, err := na.R().Get(fmt.Sprintf("/blocks/seq/%d/%d", start, end))
	if err != nil {
		return nil, fmt.Errorf("GetBlocksWithin: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetBlocksWithin: %s", resp.String())
	}
	err = json.Unmarshal(resp.Bytes(), &res)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}
	return res, nil
}

func (b *BlockResp) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}

// AvgDelayResp is response struct for GET /blocks/delay/{sig}/{num}
type AvgDelayResp struct {
	Delay int64 `json:"delay"`
}

// GetAvgDelay gets the average delay in milliseconds for a few blocks starting from the block of which the signature is given.
func (na *NodeAPI) GetAvgDelay(sig string, num int) (*AvgDelayResp, error) {
	res := &AvgDelayResp{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/blocks/delay/%s/%d", sig, num))
	if err != nil {
		return nil, fmt.Errorf("GetAvgDelay: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetAvgDelay: %s", resp.String())
	}
	return res, nil
}
