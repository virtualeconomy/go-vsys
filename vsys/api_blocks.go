package vsys

import (
	"encoding/json"
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
