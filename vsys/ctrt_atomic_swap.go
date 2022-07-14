package vsys

import "fmt"

type AtomicSwapCtrt struct {
	*Ctrt
}

func RegisterAtomicSwapCtrt(by *Account, tokenId string, ctrtDescription string) (*AtomicSwapCtrt, error) {
	ctrtMeta, err := newCtrtMetaForAtomicSwapCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterAtomicSwapCtrt: %w", err)
	}

	tokId, err := NewTokenIdFromB58Str(tokenId)

	txReq := NewRegCtrtTxReq(
		DataStack{NewDeTokenId(tokId)},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)
	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterAtomicSwapCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}

	return &AtomicSwapCtrt{
		&Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

func NewDBKeyAtomicSwapMaker() Bytes {
	return STATE_VAR_ATOMIC_SWAP_MAKER.Serialize()
}

func (a *AtomicSwapCtrt) Maker() (*Addr, error) {
	resp, err := a.QueryDBKey(
		NewDBKeyAtomicSwapMaker(),
	)
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	addr, err := NewAddrFromB58Str(resp.Val.Str())
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	return addr, nil
}

func NewDBKeyAtomicSwapTokId() Bytes {
	return STATE_VAR_ATOMIC_SWAP_TOKEN_ID.Serialize()
}

func (a *AtomicSwapCtrt) TokId() (*TokenId, error) {
	resp, err := a.QueryDBKey(
		NewDBKeyAtomicSwapTokId(),
	)
	if err != nil {
		return nil, fmt.Errorf("TokId: %w", err)
	}
	tokId, err := NewTokenIdFromB58Str(resp.Val.Str())
	if err != nil {
		return nil, fmt.Errorf("TokId: %w", err)
	}
	return tokId, nil
}

func (a AtomicSwapCtrt) Unit() (uint64, error) {
	tc, err := a.TokCtrt()
	if err != nil {
		return 0, err
	}
	return tc.Unit(), nil
}

func (a *AtomicSwapCtrt) Lock(
	by *Account,
	amount float64,
	recipient string,
	hashSecret Bytes,
	expireTime int64,
	attachment string) (*BroadcastExecuteTxResp, error) {
	// TODO: unit function
	unit, err := a.Unit()
	if err != nil {
		return nil, err
	}
	deTokAmount, err := NewDeAmountForTokAmount(amount, unit)
	if err != nil {
		return nil, fmt.Errorf("Lock: %w", err)
	}
	recipientAddr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Lock: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		a.CtrtId,
		FUNC_IDX_ATOMIC_SWAP_LOCK,
		DataStack{
			deTokAmount,
			NewDeAddr(recipientAddr),
			NewDeBytes(hashSecret),
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(expireTime)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	return by.ExecuteCtrt(txReq)
}

func (a *AtomicSwapCtrt) TokCtrt() (BaseTokCtrt, error) {
	tokId, err := a.TokId()
	if err != nil {
		return nil, err
	}
	instance, err := GetCtrtFromTokId(tokId, a.Chain)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
