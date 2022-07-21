package vsys

import (
	"fmt"
)

type QueryDBKeyInterface interface {
	QueryDBKey(bytes Bytes) (*CtrtDataResp, error)
	ctrtId() *CtrtId
	Unit() (Unit, error)
}

// internal implementation for Issuer function.
func issuer(t QueryDBKeyInterface, dbKey Bytes) (*Addr, error) {
	resp, err := t.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("Issuer: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Issuer: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// Maker queries and returns maker Addr of the contract.
func maker(t QueryDBKeyInterface, dbKey Bytes) (*Addr, error) {
	resp, err := t.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("Maker: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Maker: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// tokId is internal implementation for TokId.
func tokId(t QueryDBKeyInterface) (*TokenId, error) {
	tokId, err := t.ctrtId().GetTokId(0)
	if err != nil {
		return nil, fmt.Errorf("tokId: %w", err)
	}
	return tokId, nil
}

// supersede is internal implementation for Supersede.
func supersede(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	newIssuer string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	newIssuerAddr, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			NewDeAddr(newIssuerAddr),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	return resp, nil
}

// issue is internal implementation for Issue.
func issue(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Issue: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Issue: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Issue: %w", err)
	}

	return resp, nil
}

// send is internal implementation for Send.
func send(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	recipient string,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	rcpt_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	err = rcpt_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			NewDeAddr(rcpt_addr),
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	return resp, nil
}

// destroy is internal implementation for Destroy.
func destroy(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Destroy: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Destroy: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Destroy: %w", err)
	}

	return resp, nil
}

// Transfer transfers tokens from sender to recipient.
func transfer(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	sender, recipient string,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	sender_addr, err := NewAddrFromB58Str(sender)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	recipient_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	err = sender_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	err = recipient_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			NewDeAddr(sender_addr),
			NewDeAddr(recipient_addr),
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	return resp, nil
}

// deposit is internal implementation for Deposit.
func deposit(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	ctrtId string,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			NewDeAddr(by.Addr),
			NewDeCtrtAddrFromCtrtId(ctrtIdMd),
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	return resp, nil
}

// withdraw is internal implementation for Withdraw.
func withdraw(
	t QueryDBKeyInterface,
	funcIdx FuncIdx,
	by *Account,
	ctrtId string,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		funcIdx,
		DataStack{
			NewDeCtrtAddrFromCtrtId(ctrtIdMd),
			NewDeAddr(by.Addr),
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	return resp, nil
}

// TokCtrtWithoutSplit is the struct for Token Contract Without Split.
type TokCtrtWithoutSplit struct {
	*Ctrt
	tokId *TokenId
}

// NewTokCtrtWithoutSplit creates an instance of TokCtrtWithoutSplit from given contract id.
func NewTokCtrtWithoutSplit(ctrtId string, chain *Chain) (*TokCtrtWithoutSplit, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewTokCtrtWithoutSplit: %w", err)
	}

	return &TokCtrtWithoutSplit{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

// RegisterTokCtrtWithoutSplit registers a token contract without split.
func RegisterTokCtrtWithoutSplit(by *Account, max float64, unit uint64, token_description, ctrt_desciption string) (*TokCtrtWithoutSplit, error) {
	ctrtMeta, err := NewCtrtMetaForTokCtrtWithoutSplit()
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplit: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(max, unit)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplit: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			deAmount,
			NewDeAmount(Amount(unit)),
			NewDeStr(Str(token_description)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrt_desciption),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplit: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplit: %w", err)
	}

	return &TokCtrtWithoutSplit{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

func (t *TokCtrtWithoutSplit) ctrtId() *CtrtId {
	return t.CtrtId
}

// NewDBKeyTokCtrtWithoutSplitMaker returns DB key for querying maker of the contract.
func NewDBKeyTokCtrtWithoutSplitMaker() Bytes {
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_MAKER.Serialize()
}

// NewDBKeyTokCtrtWithoutSplitIssuer returns DB key for querying issuer of the contract.
func NewDBKeyTokCtrtWithoutSplitIssuer() Bytes {
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_ISSUER.Serialize()
}

// Issuer queries and returns maker Addr of the contract.
func (t *TokCtrtWithoutSplit) Issuer() (*Addr, error) {
	return issuer(t, NewDBKeyTokCtrtWithoutSplitIssuer())
}
func (t *TokCtrtWithoutSplit) Maker() (*Addr, error) {
	return maker(t, NewDBKeyTokCtrtWithoutSplitMaker())
}

// TokId returns TokenId of the contract.
func (t *TokCtrtWithoutSplit) TokId() (*TokenId, error) {
	if t.tokId == nil {
		got, err := tokId(t)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		t.tokId = got
	}
	return t.tokId, nil
}

// GetTokBal queries & returns the balance of the token of the contract belonging to the user address.
func (t *TokCtrtWithoutSplit) GetTokBal(addr string) (*Token, error) {
	tokId, err := t.TokId()
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	resp, err := t.Chain.NodeAPI.GetTokBal(addr, string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	tokMd := NewToken(resp.Balance, resp.Unit)
	return tokMd, nil
}

// Unit queries and returns Unit of the token of contract.
func (t *TokCtrtWithoutSplit) Unit() (Unit, error) {
	tokId, err := t.TokId()
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	info, err := t.Chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	return info.Unit, nil
}

// Supersede transfers the issuing right of the contract to another account.
func (t *TokCtrtWithoutSplit) Supersede(by *Account, newIssuer string, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersede(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SUPERSEDE, by, newIssuer, attachment)
}

// Issue issues new Tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplit) Issue(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_ISSUE, by, amount, attachment)
}

// Send sends tokens to another account.
func (t *TokCtrtWithoutSplit) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SEND, by, recipient, amount, attachment)
}

// Destroy destroys an amount of tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplit) Destroy(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return destroy(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DESTROY, by, amount, attachment)
}

// Transfer transfers tokens from sender to recipient.
func (t *TokCtrtWithoutSplit) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_TRANSFER, by, sender, recipient, amount, attachment)
}

// Deposit deposits the tokens into the contract.
func (t *TokCtrtWithoutSplit) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DEPOSIT, by, ctrtId, amount, attachment)
}

// Withdraw withdraws tokens from another contract.
func (t *TokCtrtWithoutSplit) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(t, FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_WITHDRAW, by, ctrtId, amount, attachment)
}

type TokCtrtWithSplit struct {
	*Ctrt
	tokId *TokenId
}

// NewTokCtrtWithSplit creates an instance of TokCtrtWithSplit from given contract id.
func NewTokCtrtWithSplit(ctrtId string, chain *Chain) (*TokCtrtWithSplit, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewTokCtrtWithSplit: %w", err)
	}

	return &TokCtrtWithSplit{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
		tokId: nil,
	}, nil
}

// RegisterTokCtrtWithSplit registers a token contract with split.
func RegisterTokCtrtWithSplit(by *Account, max float64, unit uint64, token_description, ctrt_desciption string) (*TokCtrtWithSplit, error) {
	ctrtMeta, err := NewCtrtMetaForTokCtrtWithSplit()
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithSplit: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(max, unit)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithSplit: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			deAmount,
			NewDeAmount(Amount(unit)),
			NewDeStr(Str(token_description)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrt_desciption),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithSplit: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithSplit: %w", err)
	}

	return &TokCtrtWithSplit{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
		tokId: nil,
	}, nil
}

func (t *TokCtrtWithSplit) ctrtId() *CtrtId {
	return t.CtrtId
}

// NewDBKeyTokCtrtWithSplitMaker returns DB key for querying maker of the contract.
func NewDBKeyTokCtrtWithSplitMaker() Bytes {
	// Uses same state var as TokCtrtWithoutSplit
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_MAKER.Serialize()
}

// NewDBKeyTokCtrtWithSplitIssuer returns DB key for querying issuer of the contract.
func NewDBKeyTokCtrtWithSplitIssuer() Bytes {
	// Uses same state var as TokCtrtWithoutSplit
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_ISSUER.Serialize()
}

// Issuer queries and returns issuer Addr of the contract.
func (t *TokCtrtWithSplit) Issuer() (*Addr, error) {
	return issuer(t, NewDBKeyTokCtrtWithSplitIssuer())
}
func (t *TokCtrtWithSplit) Maker() (*Addr, error) {
	return maker(t, NewDBKeyTokCtrtWithSplitMaker())
}

// TokId returns TokenId of the contract.
func (t *TokCtrtWithSplit) TokId() (*TokenId, error) {
	if t.tokId == nil {
		got, err := tokId(t)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		t.tokId = got
	}
	return t.tokId, nil
}

// GetTokBal queries & returns the balance of the token of the contract belonging to the user address.
func (t *TokCtrtWithSplit) GetTokBal(addr string) (*Token, error) {
	tokId, err := t.TokId()
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	resp, err := t.Chain.NodeAPI.GetTokBal(addr, string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	tokMd := NewToken(resp.Balance, resp.Unit)
	return tokMd, nil
}

// Unit queries and returns Unit of the token of contract.
func (t *TokCtrtWithSplit) Unit() (Unit, error) {
	tokId, err := t.TokId()
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	info, err := t.Chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	return info.Unit, nil
}

// Supersede transfers the issuing right of the contract to another account.
func (t *TokCtrtWithSplit) Supersede(by *Account, newIssuer string, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersede(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_SUPERSEDE, by, newIssuer, attachment)
}

// Issue issues new Tokens by account who has the issuing right.
func (t *TokCtrtWithSplit) Issue(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_ISSUE, by, amount, attachment)
}

// Send sends tokens to another account.
func (t *TokCtrtWithSplit) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_SEND, by, recipient, amount, attachment)
}

// Destroy destroys an amount of tokens by account who has the issuing right.
func (t *TokCtrtWithSplit) Destroy(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return destroy(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_DESTROY, by, amount, attachment)
}

// Transfer transfers tokens from sender to recipient.
func (t *TokCtrtWithSplit) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_TRANSFER, by, sender, recipient, amount, attachment)
}

// Deposit deposits the tokens into the contract.
func (t *TokCtrtWithSplit) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_DEPOSIT, by, ctrtId, amount, attachment)
}

// Withdraw withdraws tokens from another contract.
func (t *TokCtrtWithSplit) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(t, FUNC_IDX_TOK_CTRT_WITH_SPLIT_WITHDRAW, by, ctrtId, amount, attachment)
}

// Split updates the unit of the token contract.
func (t *TokCtrtWithSplit) Split(by *Account, newUnit uint64, attachment string) (*BroadcastExecuteTxResp, error) {
	deAmount := NewDeAmount(Amount(newUnit))

	txReq := NewExecCtrtFuncTxReq(
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITH_SPLIT_SPLIT,
		DataStack{
			deAmount,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Split: %w", err)
	}
	return resp, nil
}

// General functions for Token Contracts with Lists

func isInList(t QueryDBKeyInterface, dbKey Bytes) (bool, error) {
	resp, err := t.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("isInList: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("isInList: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// updateList updates the presence of the address within the given data entry in the list.
//        It's the helper method for UpdateList*.
func updateList(
	t QueryDBKeyInterface,
	by *Account,
	addrDe DataEntry,
	val bool,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		FUNC_IDX_TOK_CTRT_V2_UPDATE_LIST,
		DataStack{addrDe, NewDeBool(val)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("UpdateList: %w", err)
	}
	return resp, nil
}

// supersedeCtrtWithList is internal implementation of Supersede for contracts with lists.
func supersedeCtrtWithList(t QueryDBKeyInterface,
	by *Account,
	newIssuer string,
	newRegulator string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	newIssuerMd, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	newRegulatorMd, err := NewAddrFromB58Str(newRegulator)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		FUNC_IDX_TOK_CTRT_V2_SUPERSEDE,
		DataStack{
			NewDeAddr(newIssuerMd),
			NewDeAddr(newRegulatorMd),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	return resp, nil
}

func NewDBKeyForUserInList(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyForUserInList: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_TOK_CTRT_V2_IS_IN_LIST, NewDeAddr(addrMd)).Serialize(), nil
}

func NewDBKeyForCtrtInList(ctrtId string) (Bytes, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyForCtrtInList: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_TOK_CTRT_V2_IS_IN_LIST, NewDeCtrtAddrFromCtrtId(ctrtIdMd)).Serialize(), nil
}

// TokCtrtWithoutSplitV2Whitelist is the struct for Token Contract V2 with Whitelist.
type TokCtrtWithoutSplitV2Whitelist struct {
	*Ctrt
	tokId *TokenId
}

// NewTokCtrtWithoutSplitV2Whitelist creates an instance of TokCtrtWithoutSplitV2Whitelist from given contract id.
func NewTokCtrtWithoutSplitV2Whitelist(ctrtId string, ch *Chain) (*TokCtrtWithoutSplitV2Whitelist, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewTokCtrtWithoutSplitV2Whitelist: %w", err)
	}

	return &TokCtrtWithoutSplitV2Whitelist{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  ch,
		},
		tokId: nil,
	}, nil
}

// RegisterTokCtrtWithoutSplitV2Whitelist registers a token contract.
func RegisterTokCtrtWithoutSplitV2Whitelist(by *Account, max float64, unit uint64, token_description, ctrt_desciption string) (*TokCtrtWithoutSplitV2Whitelist, error) {
	ctrtMeta, err := NewCtrtMetaForTokCtrtWithoutSplitV2Whitelist()
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Whitelist: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(max, unit)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Whitelist: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			deAmount,
			NewDeAmount(Amount(unit)),
			NewDeStr(Str(token_description)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrt_desciption),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Whitelist: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Whitelist: %w", err)
	}

	return &TokCtrtWithoutSplitV2Whitelist{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
		tokId: nil,
	}, nil
}

func NewDBKeyTokCtrtV2Regulator() Bytes {
	return STATE_VAR_TOK_CTRT_V2_REGULATOR.Serialize()
}

// regulator is internal implementation for Regulator.
func regulator(t QueryDBKeyInterface) (*Addr, error) {
	resp, err := t.QueryDBKey(NewDBKeyTokCtrtV2Regulator())
	if err != nil {
		return nil, fmt.Errorf("Regulator: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("Regulator: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Regulator: CtrtDataResp.Val is %T but string was expected", val)
	}
}

func (t *TokCtrtWithoutSplitV2Whitelist) ctrtId() *CtrtId {
	return t.CtrtId
}

// TokId returns TokenId of the contract.
func (t *TokCtrtWithoutSplitV2Whitelist) TokId() (*TokenId, error) {
	if t.tokId == nil {
		got, err := tokId(t)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		t.tokId = got
	}
	return t.tokId, nil
}

// Unit returns the Unit of token contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Unit() (Unit, error) {
	tokId, err := t.TokId()
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	info, err := t.Chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	return info.Unit, nil
}

// Regulator queries & returns the regulator of the contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Regulator() (*Addr, error) {
	return regulator(t)
}

// IsUserInList queries & returns the status of whether the user address in the whitelist.
func (t *TokCtrtWithoutSplitV2Whitelist) IsUserInList(addr string) (bool, error) {
	dbKey, err := NewDBKeyForUserInList(addr)
	if err != nil {
		return false, fmt.Errorf("IsUserInList: %w", err)
	}
	return isInList(t, dbKey)
}

// IsCtrtInList queries & returns the status of whether the contract address in the whitelist.
func (t *TokCtrtWithoutSplitV2Whitelist) IsCtrtInList(ctrtId string) (bool, error) {
	dbKey, err := NewDBKeyForCtrtInList(ctrtId)
	if err != nil {
		return false, fmt.Errorf("IsCtrtInList: %w", err)
	}
	return isInList(t, dbKey)
}

// UpdateListUser updates the presence of the user address in the list.
func (t *TokCtrtWithoutSplitV2Whitelist) UpdateListUser(by *Account, addr string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("UpdateListUser: %w", err)
	}
	return updateList(t, by, NewDeAddr(addrMd), val, attachment)
}

// UpdateListCtrt updates the presence of the contract address in the list.
func (t *TokCtrtWithoutSplitV2Whitelist) UpdateListCtrt(by *Account, ctrtId string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("UpdateListCtrt: %w", err)
	}
	return updateList(t, by, NewDeCtrtAddrFromCtrtId(ctrtIdMd), val, attachment)
}

// Supersede transfers the issuer role of the contract to a new account.
func (t *TokCtrtWithoutSplitV2Whitelist) Supersede(by *Account, newIssuer, newRegulator string, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersedeCtrtWithList(t, by, newIssuer, newRegulator, attachment)
}

// Issuer queries and returns issuer Addr of the contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Issuer() (*Addr, error) {
	return issuer(t, NewDBKeyTokCtrtWithoutSplitIssuer())
}

// Maker queries and returns maker Addr of the contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Maker() (*Addr, error) {
	return maker(t, NewDBKeyTokCtrtWithoutSplitMaker())
}

// GetTokBal queries & returns the balance of the token of the contract belonging to the user address.
func (t *TokCtrtWithoutSplitV2Whitelist) GetTokBal(addr string) (*Token, error) {
	tokId, err := t.TokId()
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	resp, err := t.Chain.NodeAPI.GetTokBal(addr, string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	tokMd := NewToken(resp.Balance, resp.Unit)
	return tokMd, nil
}

// Issue issues new Tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplitV2Whitelist) Issue(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(t, FUNC_IDX_TOK_CTRT_V2_ISSUE, by, amount, attachment)
}

// Send sends tokens to another account.
func (t *TokCtrtWithoutSplitV2Whitelist) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(t, FUNC_IDX_TOK_CTRT_V2_SEND, by, recipient, amount, attachment)
}

// Destroy destroys an amount of tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplitV2Whitelist) Destroy(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return destroy(t, FUNC_IDX_TOK_CTRT_V2_DESTROY, by, amount, attachment)
}

// Transfer transfers tokens from sender to recipient.
func (t *TokCtrtWithoutSplitV2Whitelist) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(t, FUNC_IDX_TOK_CTRT_V2_TRANSFER, by, sender, recipient, amount, attachment)
}

// Deposit deposits the tokens into the contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(t, FUNC_IDX_TOK_CTRT_V2_DEPOSIT, by, ctrtId, amount, attachment)
}

// Withdraw withdraws tokens from another contract.
func (t *TokCtrtWithoutSplitV2Whitelist) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(t, FUNC_IDX_TOK_CTRT_V2_WITHDRAW, by, ctrtId, amount, attachment)
}

// TokCtrtWithoutSplitV2Blacklist is the struct for Token Contract V2 with Blacklist.
type TokCtrtWithoutSplitV2Blacklist struct {
	*Ctrt
	tokId *TokenId
}

// NewTokCtrtWithoutSplitV2Blacklist creates an instance of TokCtrtWithoutSplitV2Blacklist from given contract id.
func NewTokCtrtWithoutSplitV2Blacklist(ctrtId string, ch *Chain) (*TokCtrtWithoutSplitV2Blacklist, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewTokCtrtWithoutSplitV2Blacklist: %w", err)
	}

	return &TokCtrtWithoutSplitV2Blacklist{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  ch,
		},
		tokId: nil,
	}, nil
}

// RegisterTokCtrtWithoutSplitV2Blacklist registers a token contract.
func RegisterTokCtrtWithoutSplitV2Blacklist(by *Account, max float64, unit uint64, token_description, ctrt_desciption string) (*TokCtrtWithoutSplitV2Blacklist, error) {
	ctrtMeta, err := NewCtrtMetaForTokCtrtWithoutSplitV2Blacklist()
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Blacklist: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(max, unit)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Blacklist: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			deAmount,
			NewDeAmount(Amount(unit)),
			NewDeStr(Str(token_description)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrt_desciption),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Blacklist: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterTokCtrtWithoutSplitV2Blacklist: %w", err)
	}

	return &TokCtrtWithoutSplitV2Blacklist{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
		tokId: nil,
	}, nil
}

func (t *TokCtrtWithoutSplitV2Blacklist) ctrtId() *CtrtId {
	return t.CtrtId
}

// TokId returns TokenId of the contract.
func (t *TokCtrtWithoutSplitV2Blacklist) TokId() (*TokenId, error) {
	if t.tokId == nil {
		got, err := tokId(t)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		t.tokId = got
	}
	return t.tokId, nil
}

// Unit returns the Unit of token contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Unit() (Unit, error) {
	tokId, err := t.TokId()
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	info, err := t.Chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	return info.Unit, nil
}

// Regulator queries & returns the regulator of the contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Regulator() (*Addr, error) {
	return regulator(t)
}

// IsUserInList queries & returns the status of whether the user address in the Blacklist.
func (t *TokCtrtWithoutSplitV2Blacklist) IsUserInList(addr string) (bool, error) {
	dbKey, err := NewDBKeyForUserInList(addr)
	if err != nil {
		return false, fmt.Errorf("IsUserInList: %w", err)
	}
	return isInList(t, dbKey)
}

// IsCtrtInList queries & returns the status of whether the contract address in the Blacklist.
func (t *TokCtrtWithoutSplitV2Blacklist) IsCtrtInList(ctrtId string) (bool, error) {
	dbKey, err := NewDBKeyForCtrtInList(ctrtId)
	if err != nil {
		return false, fmt.Errorf("IsCtrtInList: %w", err)
	}
	return isInList(t, dbKey)
}

// UpdateListUser updates the presence of the user address in the list.
func (t *TokCtrtWithoutSplitV2Blacklist) UpdateListUser(by *Account, addr string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("UpdateListUser: %w", err)
	}
	return updateList(t, by, NewDeAddr(addrMd), val, attachment)
}

// UpdateListCtrt updates the presence of the contract address in the list.
func (t *TokCtrtWithoutSplitV2Blacklist) UpdateListCtrt(by *Account, ctrtId string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("UpdateListCtrt: %w", err)
	}
	return updateList(t, by, NewDeCtrtAddrFromCtrtId(ctrtIdMd), val, attachment)
}

// Supersede transfers the issuer role of the contract to a new account.
func (t *TokCtrtWithoutSplitV2Blacklist) Supersede(by *Account, newIssuer, newRegulator string, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersedeCtrtWithList(t, by, newIssuer, newRegulator, attachment)
}

// Issuer queries and returns issuer Addr of the contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Issuer() (*Addr, error) {
	return issuer(t, NewDBKeyTokCtrtWithoutSplitIssuer())
}

// Maker queries and returns maker Addr of the contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Maker() (*Addr, error) {
	return maker(t, NewDBKeyTokCtrtWithoutSplitMaker())
}

// GetTokBal queries & returns the balance of the token of the contract belonging to the user address.
func (t *TokCtrtWithoutSplitV2Blacklist) GetTokBal(addr string) (*Token, error) {
	tokId, err := t.TokId()
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	resp, err := t.Chain.NodeAPI.GetTokBal(addr, string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetTokBal: %w", err)
	}
	tokMd := NewToken(resp.Balance, resp.Unit)
	return tokMd, nil
}

// Issue issues new Tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplitV2Blacklist) Issue(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(t, FUNC_IDX_TOK_CTRT_V2_ISSUE, by, amount, attachment)
}

// Send sends tokens to another account.
func (t *TokCtrtWithoutSplitV2Blacklist) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(t, FUNC_IDX_TOK_CTRT_V2_SEND, by, recipient, amount, attachment)
}

// Destroy destroys an amount of tokens by account who has the issuing right.
func (t *TokCtrtWithoutSplitV2Blacklist) Destroy(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return destroy(t, FUNC_IDX_TOK_CTRT_V2_DESTROY, by, amount, attachment)
}

// Transfer transfers tokens from sender to recipient.
func (t *TokCtrtWithoutSplitV2Blacklist) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(t, FUNC_IDX_TOK_CTRT_V2_TRANSFER, by, sender, recipient, amount, attachment)
}

// Deposit deposits the tokens into the contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(t, FUNC_IDX_TOK_CTRT_V2_DEPOSIT, by, ctrtId, amount, attachment)
}

// Withdraw withdraws tokens from another contract.
func (t *TokCtrtWithoutSplitV2Blacklist) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(t, FUNC_IDX_TOK_CTRT_V2_WITHDRAW, by, ctrtId, amount, attachment)
}
