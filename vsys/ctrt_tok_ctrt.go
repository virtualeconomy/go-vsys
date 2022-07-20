package vsys

import "fmt"

type TokCtrtWithoutSplit struct {
	*Ctrt
}

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

func (t TokCtrtWithoutSplit) Unit() (Unit, error) {
	tokId, err := t.CtrtId.GetTokId(0)
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	info, err := t.Chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return 0, fmt.Errorf("Unit: %w", err)
	}
	return info.Unit, nil
}

func (t *TokCtrtWithoutSplit) Supersede(by *Account, newIssuer string, attachment string) (*BroadcastExecuteTxResp, error) {
	newIssuerAddr, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SUPERSEDE,
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

func (t *TokCtrtWithoutSplit) Issue(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Issue: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Issue: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_ISSUE,
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

func (t *TokCtrtWithoutSplit) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
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
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_SEND,
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

func (t *TokCtrtWithoutSplit) Destroy(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	unit, err := t.Unit()
	if err != nil {
		return nil, fmt.Errorf("Destroy: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Destroy: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DESTROY,
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

func (t *TokCtrtWithoutSplit) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
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
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_TRANSFER,
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

func (t *TokCtrtWithoutSplit) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
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
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_DEPOSIT,
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

func (t *TokCtrtWithoutSplit) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
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
		t.CtrtId,
		FUNC_IDX_TOK_CTRT_WITHOUT_SPLIT_WITHDRAW,
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

func (t *TokCtrtWithoutSplit) TokId() (*TokenId, error) {
	tokId, err := t.CtrtId.GetTokId(0)
	if err != nil {
		return nil, fmt.Errorf("TokId: %w", err)
	}
	return tokId, nil
}

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

func NewDBKeyTokCtrtWithoutSplitMaker() Bytes {
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_MAKER.Serialize()
}

func (t *TokCtrtWithoutSplit) Maker() (*Addr, error) {
	resp, err := t.QueryDBKey(NewDBKeyTokCtrtWithoutSplitMaker())
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	addr, err := NewAddrFromB58Str(resp.Val.Str())
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	return addr, nil
}
func NewDBKeyTokCtrtWithoutSplitIssuer() Bytes {
	return STATE_VAR_TOK_CTRT_WITHOUT_SPLIT_ISSUER.Serialize()
}

func (t *TokCtrtWithoutSplit) Issuer() (*Addr, error) {
	resp, err := t.QueryDBKey(NewDBKeyTokCtrtWithoutSplitIssuer())
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}
	addr, err := NewAddrFromB58Str(resp.Val.Str())
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}
	return addr, nil
}
