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

// GenesisTxInfoResp struct representing response from /transactions/info/{txId}.
// NOTE: Genesis Transaction (Type 1) doesn't have Proofs and FeeScale.
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

type ContentSlotsTxInfoResp struct {
	TxGeneral
	SlotId int `json:"slotId"`
}

func (c *ContentSlotsTxInfoResp) GetTxGeneral() TxGeneral {
	return c.TxGeneral
}

func (c *ContentSlotsTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
}

type ReleaseSlotsTxInfoResp struct {
	TxGeneral
	ContractId string `json:"contractId"`
	Contract   struct {
		LanguageCode    string   `json:"languageCode"`
		LanguageVersion int      `json:"languageVersion"`
		Triggers        []string `json:"triggers"`
		Descriptors     []string `json:"descriptors"`
		StateVariables  []string `json:"stateVariables"`
		StateMaps       []string `json:"stateMaps"`
		Textual         struct {
			Triggers       string `json:"triggers"`
			Descriptors    string `json:"descriptors"`
			StateVariables string `json:"stateVariables"`
			StateMaps      string `json:"stateMaps"`
		} `json:"textual"`
	} `json:"contract"`
	InitData    string `json:"initData"`
	Description string `json:"description"`
}

func (r *ReleaseSlotsTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", r, *r)
}

func (r *ReleaseSlotsTxInfoResp) GetTxGeneral() TxGeneral {
	return r.TxGeneral
}

type RegCtrtTxInfoResp struct {
	TxGeneral
	CtrtId      Str          `json:"contractId"`
	Ctrt        CtrtMetaJSON `json:"contract"`
	InitData    Str          `json:"initData"`
	Description Str          `json:"description"`
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

type DbPutTxInfoResp struct {
	TxGeneral
	DbKey string `json:"dbKey"`
	Entry struct {
		Data string `json:"data"`
		Type string `json:"type"`
	} `json:"entry"`
}

func (d *DbPutTxInfoResp) GetTxGeneral() TxGeneral {
	return d.TxGeneral
}

func (d *DbPutTxInfoResp) String() string {
	return fmt.Sprintf("%T(%+v)", d, *d)
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
	case TX_TYPE_CONTENT_SLOTS:
		res = &ContentSlotsTxInfoResp{}
	case TX_TYPE_RELEASE_SLOTS:
		res = &ReleaseSlotsTxInfoResp{}
	case TX_TYPE_DB_PUT:
		res = &DbPutTxInfoResp{}
	}

	err = json.Unmarshal(resp.Bytes(), res)
	if err != nil {
		return nil, fmt.Errorf("GetTxInfo: %w", err)
	}
	return res, nil
}
