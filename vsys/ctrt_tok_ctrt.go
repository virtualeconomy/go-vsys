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
	// TODO: move to MustOn() bool function
	if rcpt_addr.ChainID() != by.Chain.ChainID {
		return nil, fmt.Errorf("Send: Adress must be on same chain")
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
