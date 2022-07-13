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

type CtrtDataResp struct {
	CtrtId   Str    `json:"contractId"`
	Key      Str    `json:"key"`
	Height   Height `json:"height"`
	DbName   Str    `json:"dbName"`
	DataType Str    `json:"dataType"`
	Val      Str    `json:"value"`
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
