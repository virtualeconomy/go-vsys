package vsys

import (
	"fmt"
)

type AtomicSwapCtrt struct {
	*Ctrt
}

func RegisterAtomicSwapCtrt(by *Account, tokenId, ctrtDescription string) (*AtomicSwapCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForAtomicSwapCtrt()
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
	addr, err := NewAddrFromB58Str(resp.Val.(string))
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
	tokId, err := NewTokenIdFromB58Str(resp.Val.(string))
	if err != nil {
		return nil, fmt.Errorf("TokId: %w", err)
	}
	return tokId, nil
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

func (a AtomicSwapCtrt) Unit() (Unit, error) {
	tc, err := a.TokCtrt()
	if err != nil {
		return 0, err
	}
	return tc.Unit(), nil
}

func (a *AtomicSwapCtrt) GetCtrtBal(addr string) (*Token, error) {
	// TODO: separate into StateMap function + mb get statemapconstant
	Addr, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("Lock: %w", err)
	}
	raw_val, err := a.QueryDBKey(append(PackUInt8(uint8(STATE_MAP_IDX_ATOMIC_SWAP_CONTRACT_BALANCE)),
		NewDeAddr(Addr).Serialize()...))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	fmt.Println(raw_val)
	unit, err := a.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	// TODO: verify raw_val format
	amount := raw_val.Val.(float64)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	return NewToken(Amount(amount), unit), nil
}

func (a *AtomicSwapCtrt) Lock(
	by *Account,
	amount float64,
	recipient string,
	hashSecret Bytes,
	expireTime int64,
	attachment string) (*BroadcastExecuteTxResp, error) {
	unit, err := a.Unit()
	if err != nil {
		return nil, err
	}
	deTokAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
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
