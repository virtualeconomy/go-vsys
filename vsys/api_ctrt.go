package vsys

import (
	"fmt"
)

type BroadcastRegisterPayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	CtrtMeta     Str           `json:"contract"`
	InitData     Str           `json:"initData"`
	Description  Str           `json:"description"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Signature    Str           `json:"signature"`
}

// BroadcastRegisterTxResp is the response for calling endpoint /contract/broadcast/register
type BroadcastRegisterTxResp struct {
	TxBasic

	CtrtId      Str `json:"contractId"`
	InitData    Str `json:"initData"`
	Description Str `json:"description"`

	CtrtMeta CtrtMetaResp `json:"contract"`
}

func (na *NodeAPI) BroadcastRegister(p *BroadcastRegisterPayload) (*BroadcastRegisterTxResp, error) {
	res := &BroadcastRegisterTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/contract/broadcast/register")

	if err != nil {
		return nil, fmt.Errorf("BroadcastRegister: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastRegister: %s", resp.String())
	}

	return res, nil
}

type BroadcastExecutePayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	CtrtId       Str           `json:"contractId"`
	FuncIdx      FuncIdx       `json:"functionIndex"`
	FuncData     Str           `json:"functionData"`
	Attachment   Str           `json:"attachment"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Signature    Str           `json:"signature"`
}

// BroadcastExecuteTxResp is the response for calling endpoint /contract/broadcast/execute
type BroadcastExecuteTxResp struct {
	TxBasic

	CtrtId     Str     `json:"contractId"`
	FuncIdx    FuncIdx `json:"functionIndex"`
	FuncData   Str     `json:"functionData"`
	Attachment Str     `json:"attachment"`
}

func (na *NodeAPI) BroadcastExecute(p *BroadcastExecutePayload) (*BroadcastExecuteTxResp, error) {
	res := &BroadcastExecuteTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/contract/broadcast/execute")

	if err != nil {
		return nil, fmt.Errorf("BroadcastExecute: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastExecute: %s", err)
	}
	return res, nil
}

// CtrtDataResp is the response for calling endpoint /contract/data/{contractId}/{key}
type CtrtDataResp struct {
	CtrtId   Str         `json:"contractId"`
	Key      Str         `json:"key"`
	Height   Height      `json:"height"`
	DbName   Str         `json:"dbName"`
	DataType Str         `json:"dataType"`
	Val      interface{} `json:"value"` // Can be number or string
}

func (na *NodeAPI) GetCtrtData(ctrtId, key string) (*CtrtDataResp, error) {
	res := &CtrtDataResp{}
	resp, err := na.R().SetResult(res).Get(
		fmt.Sprintf("/contract/data/%s/%s", ctrtId, key),
	)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtData: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetCtrtData: %s", resp.String())
	}
	return res, nil
}

// TokInfoResp is the response for calling endpoint /contract/tokenInfo/{tokenId}
type TokInfoResp struct {
	TokId       Str    `json:"tokenId"`
	CtrtId      Str    `json:"contractId"`
	Max         Amount `json:"max"`
	Total       Amount `json:"total"`
	Unit        Unit   `json:"unity"`
	Description Str    `json:"description"`
}

func (na *NodeAPI) GetTokInfo(tokId string) (*TokInfoResp, error) {
	res := &TokInfoResp{}
	resp, err := na.R().SetResult(res).Get(
		fmt.Sprintf("/contract/tokenInfo/%s", tokId),
	)
	if err != nil {
		return nil, fmt.Errorf("GetTokInfo: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetTokInfo: %s", resp.String())
	}
	return res, nil
}

// CtrtInfoResp is the response for calling endpoint /contract/info/{contractId}
type CtrtInfoResp struct {
	CtrtId Str `json:"contractId"`
	TxId   Str `json:"transactionId"`
	Type   Str `json:"type"`
	Info   []struct {
		Data Str `json:"data"`
		Type Str `json:"type"`
		Name Str `json:"name"`
	} `json:"info"`
	Height Height `json:"height"`
}

func (na *NodeAPI) GetCtrtInfo(ctrtId string) (*CtrtInfoResp, error) {
	res := &CtrtInfoResp{}
	resp, err := na.R().SetResult(res).Get(
		fmt.Sprintf("/contract/info/%s", ctrtId),
	)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtInfo: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetCtrtInfo: %s", resp.String())
	}
	return res, nil
}

// TokBalResp is the response for calling endpoint /contract/balance/{addr}/{tokenId}
type TokBalResp struct {
	Addr    Str    `json:"address/contractId"`
	Height  Height `json:"height"`
	TokId   Str    `json:"tokenId"`
	Balance Amount `json:"balance"`
	Unit    Unit   `json:"unity"`
}

// GetTokBal queries and returns response of /contract/balance/{addr}/{tokenId} endpoint
func (na *NodeAPI) GetTokBal(addr, tokenId string) (*TokBalResp, error) {
	res := &TokBalResp{}
	resp, err := na.R().SetResult(res).Get(
		fmt.Sprintf("/contract/balance/%s/%s", addr, tokenId),
	)
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetTokBal: %s", resp.String())
	}
	return res, nil
}
