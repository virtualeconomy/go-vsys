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

func (p *PaymentTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", p, *p)
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

func (r *RegCtrtTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", r, *r)
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

func (e *ExecCtrtFuncTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", e, *e)
}

type LeaseTxInfoResp struct {
	TxGeneral

	Amount      VSYS `json:"amount"`
	Recipient   Str  `json:"recipient"`
	LeaseStatus Str  `json:"leaseStatus"`
}

func (l *LeaseTxInfoResp) GetTxGeneral() TxGeneral {
	return l.TxGeneral
}

func (l *LeaseTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", l, *l)
}

type MintTxInfoResp struct {
	TxGeneral

	Recipient          Str    `json:"recipient"`
	Amount             VSYS   `json:"amount"`
	CurrentBlockHeight Height `json:"currentBlockHeight"`
}

func (m *MintTxInfoResp) GetTxGeneral() TxGeneral {
	return m.TxGeneral
}

func (m *MintTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", m, *m)
}

type LeaseCancelTxInfoResp struct {
	TxGeneral
	LeaseId Str `json:"leaseId"`
	Lease   struct {
		TxBasic
		Amount    VSYS `json:"amount"`
		Recipient Str  `json:"recipient"`
	} `json:"lease"`
}

func (l *LeaseCancelTxInfoResp) GetTxGeneral() TxGeneral {
	return l.TxGeneral
}

func (l *LeaseCancelTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", l, *l)
}

// genesis block does not have proofs and fee scale, should we use different struct?
type GenesisTxInfoResp struct {
	TxGeneral
	SlotId    int  `json:"slotId"`
	Signature Str  `json:"signature"`
	Recipient Str  `json:"recipient"`
	Amount    VSYS `json:"amount"`
}

func (g *GenesisTxInfoResp) GetTxGeneral() TxGeneral {
	return g.TxGeneral
}

func (g *GenesisTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", g, *g)
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

	// TODO: add missing Tx types
	switch tg.Type {
	case TX_TYPE_PAYMENT:
		res = &PaymentTxInfoResp{}
	case TX_TYPE_LEASE:
		res = &LeaseTxInfoResp{}
	case TX_TYPE_LEASE_CANCEL:
		res = &LeaseCancelTxInfoResp{}
	case TX_TYPE_MINTING:
		res = &MintTxInfoResp{}
	case TX_TYPE_REGISTER_CONTRACT:
		res = &RegCtrtTxInfoResp{}
	case TX_TYPE_EXECUTE_CONTRACT_FUNCTION:
		res = &ExecCtrtFuncTxInfoResp{}
	case TX_TYPE_GENESIS:
		res = &GenesisTxInfoResp{}
	}

	err = json.Unmarshal(resp.Bytes(), res)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}
	return res, nil
}
