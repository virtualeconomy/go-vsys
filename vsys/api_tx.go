package vsys

import (
	"encoding/json"
	"fmt"
)

type TxGeneral struct {
	TxBasic

	Status     Str    `json:"status"`
	FeeCharged VSYS   `json:"feeCharged"`
	Height     Height `json:"height"`
}

type TxInfoResp interface {
	GetTxGeneral() TxGeneral
}

type PaymentTxInfoResp struct {
	TxGeneral

	Recipient  Str  `json:"recipient"`
	Amount     VSYS `json:"amount"`
	Attachment Str  `json:"attachment"`
}

func (p *PaymentTxInfoResp) GetTxGeneral() TxGeneral {
	return p.TxGeneral
}

type RegCtrtTxInfoResp struct {
	TxGeneral

	CtrtId      Str      `json:"contractId"`
	Ctrt        CtrtMeta `json:"contract"`
	InitData    Str      `json:"initData"`
	Description Str      `json:"description"`
}

func (r *RegCtrtTxInfoResp) GetTxGeneral() TxGeneral {
	return r.TxGeneral
}

type ExecCtrtFuncTxInfoResp struct {
	TxGeneral

	CtrtId     Str     `json:"contractId"`
	FuncIdx    FuncIdx `json:"functionIndex"`
	FuncData   Str     `json:"functionData"`
	Attachment Str     `json:"attachment"`
}

func (e *ExecCtrtFuncTxInfoResp) GetTxGeneral() TxGeneral {
	return e.TxGeneral
}

func (na *NodeAPI) GetTxInfo(txId string) (TxInfoResp, error) {
	resp, err := na.R().Get("/transactions/info/" + txId)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetTxInfo: %s", resp.String())
	}

	tg := &TxGeneral{}
	err = json.Unmarshal(resp.Bytes(), tg)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}

	var res TxInfoResp

	switch tg.Type {
	case TX_TYPE_PAYMENT:
		res = &PaymentTxInfoResp{}
	case TX_TYPE_REGISTER_CONTRACT:
		res = &RegCtrtTxInfoResp{}
	case TX_TYPE_EXECUTE_CONTRACT_FUNCTION:
		res = &ExecCtrtFuncTxInfoResp{}
	}

	err = json.Unmarshal(resp.Bytes(), res)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}
	return res, nil
}
