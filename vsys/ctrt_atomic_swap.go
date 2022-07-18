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
	if err != nil {
		return nil, fmt.Errorf("RegisterAtomicSwapCtrt: %w", err)
	}

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

func NewDBKeyAtomicSwapGetCtrtBal(addr *Addr) Bytes {
	return NewStateMap(
		STATE_MAP_IDX_ATOMIC_SWAP_CONTRACT_BALANCE,
		NewDeAddr(addr)).Serialize()
}

func (a *AtomicSwapCtrt) GetCtrtBal(addr string) (*Token, error) {
	query_addr, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	data, err := a.QueryDBKey(NewDBKeyAtomicSwapGetCtrtBal(query_addr))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	unit, err := a.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	switch amount := data.Val.(type) {
	case float64:
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetCtrtBal: CtrtDataResp.Val is %T but float64 was expected", amount)
	}
}

func NewDBKeyAtomicSwapOwner(txId string) (Bytes, error) {
	de, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapOwner: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_OWNER, NewDeBytes(de)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapOwner(txId string) (*Addr, error) {
	dbKey, err := NewDBKeyAtomicSwapOwner(txId)
	if err != nil {
		return nil, fmt.Errorf("GetSwapOwner: %w", err)
	}

	data, err := a.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetSwapOwner: %w", err)
	}

	switch addrB58 := data.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("GetSwapOwner: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("GetSwapOwner: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

func NewDBKeyAtomicSwapRecipient(txId string) (Bytes, error) {
	de, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapRecipient: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_RECIPIENT, NewDeBytes(de)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapRecipient(txId string) (*Addr, error) {
	dbKey, err := NewDBKeyAtomicSwapRecipient(txId)
	if err != nil {
		return nil, fmt.Errorf("GetSwapRecipient: %w", err)
	}

	data, err := a.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetSwapRecipient: %w", err)
	}

	switch addrB58 := data.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("GetSwapRecipient: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("GetSwapRecipient: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

func NewDBKeyAtomicSwapPuzzle(txId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapPuzzle: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_PUZZLE, NewDeBytes(b)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapPuzzle(txId string) (Str, error) {
	// Returns b58 encoded puzzle
	dbKey, err := NewDBKeyAtomicSwapPuzzle(txId)
	if err != nil {
		return "", fmt.Errorf("GetSwapPuzzle: %w", err)
	}

	resp, err := a.QueryDBKey(dbKey)
	if err != nil {
		return "", fmt.Errorf("GetSwapPuzzle: %w", err)
	}
	return Str(resp.Val.(string)), nil
}

func NewDBKeyAtomicSwapAmount(txId string) (Bytes, error) {
	de, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapAmount: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_AMOUNT, NewDeBytes(de)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapAmount(txId string) (*Token, error) {
	dbKey, err := NewDBKeyAtomicSwapAmount(txId)
	if err != nil {
		return nil, fmt.Errorf("GetSwapAmount: %w", err)
	}

	data, err := a.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetSwapAmount: %w", err)
	}
	unit, err := a.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetSwapAmount: %w", err)
	}

	switch amount := data.Val.(type) {
	case float64:
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetSwapAmount: CtrtDataResp.Val is %T but float64 was expected", amount)
	}
}

func NewDBKeyAtomicSwapExpiredTime(txId string) (Bytes, error) {
	de, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapExpiredTime: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_EXPIRED_TIME, NewDeBytes(de)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapExpiredTime(txId string) (VSYSTimestamp, error) {
	dbKey, err := NewDBKeyAtomicSwapExpiredTime(txId)
	if err != nil {
		return 0, fmt.Errorf("GetSwapExpiredTime: %w", err)
	}

	resp, err := a.QueryDBKey(dbKey)
	if err != nil {
		return 0, fmt.Errorf("GetSwapExpiredTime: %w", err)
	}

	switch timestamp := resp.Val.(type) {
	case float64:
		return VSYSTimestamp(timestamp), nil
	default:
		return 0, fmt.Errorf("GetSwapExpiredTime: CtrtDataResp.Val is %T but float64 was expected", timestamp)
	}
}

func NewDBKeyAtomicSwapStatus(txId string) (Bytes, error) {
	de, err := NewBytesFromB58Str(txId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyAtomicSwapStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_ATOMIC_SWAP_STATUS, NewDeBytes(de)).Serialize(), nil
}

func (a *AtomicSwapCtrt) GetSwapStatus(txId string) (bool, error) {
	dbKey, err := NewDBKeyAtomicSwapStatus(txId)
	if err != nil {
		return false, fmt.Errorf("GetSwapStatus: %w", err)
	}

	resp, err := a.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetSwapStatus: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("GetSwapExpiredTime: CtrtDataResp.Val is %T but float64 was expected", val)
	}
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
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Lock: %w", err)
	}
	return resp, nil
}

func (a *AtomicSwapCtrt) Solve(
	by *Account,
	lockTxId string,
	secret string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	lockTxIdBytes, err := NewBytesFromB58Str(lockTxId)
	if err != nil {
		return nil, fmt.Errorf("Solve: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		a.CtrtId,
		FUNC_IDX_ATOMIC_SWAP_SOLVE_PUZZLE,
		DataStack{
			NewDeBytes(lockTxIdBytes),
			NewDeBytes(NewBytesFromStr(secret)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Solve: %w", err)
	}
	return resp, nil
}

func (a *AtomicSwapCtrt) ExpWithdraw(
	by *Account,
	lockTxId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	lockTxIdBytes, err := NewBytesFromB58Str(lockTxId)
	if err != nil {
		return nil, fmt.Errorf("ExpWithdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		a.CtrtId,
		FUNC_IDX_ATOMIC_SWAP_EXPIRE_WITHDRAW,
		DataStack{
			NewDeBytes(lockTxIdBytes),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("ExpWithdraw: %w", err)
	}
	return resp, nil
}
