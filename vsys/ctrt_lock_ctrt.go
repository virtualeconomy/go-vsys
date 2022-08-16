package vsys

import "fmt"

// LockCtrt is the struct for Lock Contract.
type LockCtrt struct {
	*Ctrt
	tokId   *TokenId
	tokCtrt BaseTokCtrt
}

// NewLockCtrt creates instance of LockCtrt from given contract id.
func NewLockCtrt(ctrtId string, chain *Chain) (*LockCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewLockCtrt: %w", err)
	}

	return &LockCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

// RegisterLockCtrt registers Lock Contract.
func RegisterLockCtrt(by *Account, tokenId, ctrtDescription string) (*LockCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForLockCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterLockCtrt: %w", err)
	}

	tokIdMd, err := NewTokenIdFromB58Str(tokenId)
	if err != nil {
		return nil, fmt.Errorf("RegisterLockCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{NewDeTokenId(tokIdMd)},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)
	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterLockCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterLockCtrt: %w", err)
	}

	return &LockCtrt{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

func NewDBKeyLockCtrtMaker() Bytes {
	return STATE_VAR_LOCK_CTRT_MAKER.Serialize()
}

// Maker queries and returns Addr of the contract maker.
func (l *LockCtrt) Maker() (*Addr, error) {
	resp, err := l.QueryDBKey(
		NewDBKeyLockCtrtMaker(),
	)
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}

	addr, err := ctrtDataRespToAddr(resp)
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	return addr, nil
}

// NewDBKeyLockCtrtTokId returns DB key to query TokenId of contract's token.
func NewDBKeyLockCtrtTokId() Bytes {
	return STATE_VAR_LOCK_CTRT_TOKEN_ID.Serialize()
}

// TokId queries and returns TokenId of the contract's token.
func (l *LockCtrt) TokId() (*TokenId, error) {
	if l.tokId == nil {
		resp, err := l.QueryDBKey(
			NewDBKeyLockCtrtTokId(),
		)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		tokId, err := ctrtDataRespToTokenId(resp)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		l.tokId = tokId
	}
	return l.tokId, nil
}

// TokCtrt queries and returns instance of token contract of Lock contract's token.
func (l *LockCtrt) TokCtrt() (BaseTokCtrt, error) {
	if l.tokCtrt == nil {
		tokId, err := l.TokId()
		if err != nil {
			return nil, fmt.Errorf("TokCtrt: %w", err)
		}
		instance, err := GetCtrtFromTokId(tokId, l.Chain)
		if err != nil {
			return nil, fmt.Errorf("TokCtrt: %w", err)
		}
		l.tokCtrt = instance
	}
	return l.tokCtrt, nil
}

// Unit queries and returns Unit of the token of contract.
func (l *LockCtrt) Unit() (Unit, error) {
	if l.tokCtrt == nil {
		_, err := l.TokCtrt() // TokCtrt sets l.TokCtrt
		if err != nil {
			return 0, fmt.Errorf("Unit: %w", err)
		}
	}
	return l.tokCtrt.Unit()
}

// NewDBKeyLockCtrtGetCtrtBal returns DB key for querying the contract balance for given address.
func NewDBKeyLockCtrtGetCtrtBal(addr *Addr) Bytes {
	return NewStateMap(
		STATE_MAP_IDX_LOCK_CTRT_CONTRACT_BALANCE,
		NewDeAddr(addr)).Serialize()
}

// GetCtrtBal queries & returns the balance of the token within this contract belonging to the user address.
func (l *LockCtrt) GetCtrtBal(addr string) (*Token, error) {
	query_addr, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	resp, err := l.QueryDBKey(NewDBKeyLockCtrtGetCtrtBal(query_addr))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	unit, err := l.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	return tok, nil
}

// NewDBKeyLockCtrtLockTime returns DB key for querying the contract balance for given address.
func NewDBKeyLockCtrtLockTime(addr *Addr) Bytes {
	return NewStateMap(
		STATE_MAP_IDX_LOCK_CTRT_CONTRACT_LOCK_TIME,
		NewDeAddr(addr)).Serialize()
}

// GetCtrtLockTime queries & returns the balance of the token within this contract belonging to the user address.
func (l *LockCtrt) GetCtrtLockTime(addr string) (VSYSTimestamp, error) {
	query_addr, err := NewAddrFromB58Str(addr)
	if err != nil {
		return 0, fmt.Errorf("GetCtrtLockTime: %w", err)
	}

	resp, err := l.QueryDBKey(NewDBKeyLockCtrtLockTime(query_addr))
	if err != nil {
		return 0, fmt.Errorf("GetCtrtLockTime: %w", err)
	}

	ts, err := ctrtDataRespToVSYSTimestamp(resp)
	if err != nil {
		return 0, fmt.Errorf("GetCtrtLockTime: %w", err)
	}
	return ts, nil
}

// Lock locks the user's deposited tokens in the contract until the given timestamp.
func (l *LockCtrt) Lock(by *Account, expireAt int64, attachment string) (*BroadcastExecuteTxResp, error) {
	txReq := NewExecCtrtFuncTxReq(
		l.CtrtId,
		FUNC_IDX_LOCK_CTRT_LOCK,
		DataStack{
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(expireAt)),
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
