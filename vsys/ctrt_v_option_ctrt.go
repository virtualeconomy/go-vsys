package vsys

import (
	"fmt"
)

type VOptionCtrt struct {
	*Ctrt
	baseTokId   *TokenId
	targetTokId *TokenId
	optionTokId *TokenId
	proofTokId  *TokenId

	baseTokCtrt   BaseTokCtrt
	targetTokCtrt BaseTokCtrt
	optionTokCtrt BaseTokCtrt
	proofTokCtrt  BaseTokCtrt
}

func NewVOptionCtrt(ctrtId string, chain *Chain) (*VOptionCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewVOptioCtrt: %w", err)
	}

	return &VOptionCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

func RegisterVOptionCtrt(
	by *Account,
	baseTokId, targetTokId, optionTokId, proofTokId string,
	executeTime, executeDeadline int64,
	ctrtDescription string,
) (*VOptionCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForVOptionCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}

	baseTokIdMd, err := NewTokenIdFromB58Str(baseTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}
	targetTokIdMd, err := NewTokenIdFromB58Str(targetTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}
	optionTokIdMd, err := NewTokenIdFromB58Str(optionTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}
	proofTokIdMd, err := NewTokenIdFromB58Str(proofTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			NewDeTokenId(baseTokIdMd),
			NewDeTokenId(targetTokIdMd),
			NewDeTokenId(optionTokIdMd),
			NewDeTokenId(proofTokIdMd),
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(executeTime)),
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(executeDeadline)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)
	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterVOptionCtrt: %w", err)
	}
	return &VOptionCtrt{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

// Supersede transfers the ownership of the contract to another Addr.
func (v *VOptionCtrt) Supersede(by *Account, newIssuer, attachment string) (*BroadcastExecuteTxResp, error) {
	addr, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_SUPERSEDE,
		DataStack{NewDeAddr(addr)},
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

// Activate activates V Option contract to store option and proof token into the pool.
func (v *VOptionCtrt) Activate(
	by *Account,
	maxIssueNum, price, priceUnit float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	optionTokUnit, err := v.OptionTokUnit()
	if err != nil {
		return nil, fmt.Errorf("Activate: %w", err)
	}
	deAmountMaxIssueNum, err := NewDeAmountForTokAmount(maxIssueNum, uint64(optionTokUnit))
	if err != nil {
		return nil, fmt.Errorf("Activate: %w", err)
	}
	deAmountPrice, err := NewDeAmountForTokAmount(maxIssueNum, uint64(optionTokUnit))
	if err != nil {
		return nil, fmt.Errorf("Activate: %w", err)
	}
	deAmountPriceUnit, err := NewDeAmountForTokAmount(maxIssueNum, uint64(optionTokUnit))
	if err != nil {
		return nil, fmt.Errorf("Activate: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_ACTIVATE,
		DataStack{deAmountMaxIssueNum, deAmountPrice, deAmountPriceUnit},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Activate: %w", err)
	}
	return resp, nil
}

func (v *VOptionCtrt) Mint(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("Mint: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("Mint: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_MINT,
		DataStack{deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Mint: %w", err)
	}
	return resp, nil
}

func (v *VOptionCtrt) Unlock(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("Unlock: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("Unlock: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_UNLOCK,
		DataStack{deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Unlock: %w", err)
	}
	return resp, nil
}

func (v *VOptionCtrt) Execute(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("Execute: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("Execute: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_EXECUTE,
		DataStack{deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Execute: %w", err)
	}
	return resp, nil
}

func (v *VOptionCtrt) Collect(by *Account, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	optionUnit, err := v.OptionTokUnit()
	if err != nil {
		return nil, fmt.Errorf("Collect: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(optionUnit))
	if err != nil {
		return nil, fmt.Errorf("Collect: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_OPTION_COLLECT,
		DataStack{deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Collect: %w", err)
	}
	return resp, nil
}

func NewDBKeyVOptionForMaker() Bytes {
	return STATE_VAR_V_OPTION_MAKER.Serialize()
}

func (v *VOptionCtrt) Maker() (*Addr, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForMaker())
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

func (v *VOptionCtrt) BaseTokUnit() (Unit, error) {
	tc, err := v.BaseTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("BaseTokUnit: %w", err)
	}
	return tc.Unit()
}

func (v *VOptionCtrt) BaseTokCtrt() (BaseTokCtrt, error) {
	if v.baseTokCtrt == nil {
		baseTokId, err := v.BaseTokId()
		if err != nil {
			return nil, fmt.Errorf("BaseTokCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(baseTokId, v.Chain)
		if err != nil {
			return nil, fmt.Errorf("BaseTokCtrt: %w", err)
		}
		v.baseTokCtrt = tc
	}
	return v.baseTokCtrt, nil
}

func NewDBKeyVOptionCtrtBaseTokId() Bytes {
	return STATE_VAR_V_OPTION_BASE_TOKEN_ID.Serialize()
}

func (v *VOptionCtrt) BaseTokId() (*TokenId, error) {
	if v.baseTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtBaseTokId())
		if err != nil {
			return nil, fmt.Errorf("BaseTokId: %w", err)
		}

		switch val := resp.Val.(type) {
		case string:
			tokId, err := NewTokenIdFromB58Str(val)
			if err != nil {
				return nil, fmt.Errorf("BaseTokId: %w", err)
			}
			v.baseTokId = tokId
		default:
			return nil, fmt.Errorf("BaseTokId: CtrtDataResp.Val is %T but string was expected.", val)
		}
	}
	return v.baseTokId, nil
}

func (v *VOptionCtrt) TargetTokUnit() (Unit, error) {
	tc, err := v.TargetTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("TargetTokUnit: %w", err)
	}
	return tc.Unit()
}

func (v *VOptionCtrt) TargetTokCtrt() (BaseTokCtrt, error) {
	if v.targetTokCtrt == nil {
		targetTokId, err := v.TargetTokId()
		if err != nil {
			return nil, fmt.Errorf("TargetTokCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(targetTokId, v.Chain)
		if err != nil {
			return nil, fmt.Errorf("TargetTokCtrt: %w", err)
		}
		v.targetTokCtrt = tc
	}
	return v.targetTokCtrt, nil
}

func NewDBKeyVOptionCtrtTargetTokId() Bytes {
	return STATE_VAR_V_OPTION_TARGET_TOKEN_ID.Serialize()
}

func (v *VOptionCtrt) TargetTokId() (*TokenId, error) {
	if v.targetTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtTargetTokId())
		if err != nil {
			return nil, fmt.Errorf("TargetTokId: %w", err)
		}

		switch val := resp.Val.(type) {
		case string:
			tokId, err := NewTokenIdFromB58Str(val)
			if err != nil {
				return nil, fmt.Errorf("TargetTokId: %w", err)
			}
			v.targetTokId = tokId
		default:
			return nil, fmt.Errorf("TargetTokId: CtrtDataResp.Val is %T but string was expected.", val)
		}
	}
	return v.targetTokId, nil
}

func (v *VOptionCtrt) OptionTokUnit() (Unit, error) {
	tc, err := v.OptionTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("OptionTokUnit: %w", err)
	}
	return tc.Unit()
}

func (v *VOptionCtrt) OptionTokCtrt() (BaseTokCtrt, error) {
	if v.optionTokCtrt == nil {
		optionTokId, err := v.OptionTokId()
		if err != nil {
			return nil, fmt.Errorf("OptionTokCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(optionTokId, v.Chain)
		if err != nil {
			return nil, fmt.Errorf("OptionTokCtrt: %w", err)
		}
		v.optionTokCtrt = tc
	}
	return v.optionTokCtrt, nil
}

func NewDBKeyVOptionCtrtOptionTokId() Bytes {
	return STATE_VAR_V_OPTION_OPTION_TOKEN_ID.Serialize()
}

func (v *VOptionCtrt) OptionTokId() (*TokenId, error) {
	if v.optionTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtOptionTokId())
		if err != nil {
			return nil, fmt.Errorf("OptionTokId: %w", err)
		}

		switch val := resp.Val.(type) {
		case string:
			tokId, err := NewTokenIdFromB58Str(val)
			if err != nil {
				return nil, fmt.Errorf("OptionTokId: %w", err)
			}
			v.optionTokId = tokId
		default:
			return nil, fmt.Errorf("OptionTokId: CtrtDataResp.Val is %T but string was expected.", val)
		}
	}
	return v.optionTokId, nil
}

func (v *VOptionCtrt) ProofTokUnit() (Unit, error) {
	tc, err := v.ProofTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("ProofTokUnit: %w", err)
	}
	return tc.Unit()
}

func (v *VOptionCtrt) ProofTokCtrt() (BaseTokCtrt, error) {
	if v.proofTokCtrt == nil {
		proofTokId, err := v.ProofTokId()
		if err != nil {
			return nil, fmt.Errorf("ProofTokCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(proofTokId, v.Chain)
		if err != nil {
			return nil, fmt.Errorf("ProofTokCtrt: %w", err)
		}
		v.proofTokCtrt = tc
	}
	return v.proofTokCtrt, nil
}

func NewDBKeyVOptionCtrtProofTokId() Bytes {
	return STATE_VAR_V_OPTION_PROOF_TOKEN_ID.Serialize()
}

func (v *VOptionCtrt) ProofTokId() (*TokenId, error) {
	if v.proofTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtProofTokId())
		if err != nil {
			return nil, fmt.Errorf("ProofTokId: %w", err)
		}

		switch val := resp.Val.(type) {
		case string:
			tokId, err := NewTokenIdFromB58Str(val)
			if err != nil {
				return nil, fmt.Errorf("ProofTokId: %w", err)
			}
			v.proofTokId = tokId
		default:
			return nil, fmt.Errorf("ProofTokId: CtrtDataResp.Val is %T but string was expected.", val)
		}
	}
	return v.proofTokId, nil
}

func NewDBKeyVOptionCtrtForExecuteTime() Bytes {
	return STATE_VAR_V_OPTION_EXECUTE_TIME.Serialize()
}

func (v *VOptionCtrt) ExecuteTime() (VSYSTimestamp, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtForExecuteTime())
	if err != nil {
		return 0, fmt.Errorf("ExecuteTime: %w", err)
	}

	switch timestamp := resp.Val.(type) {
	case float64:
		return VSYSTimestamp(timestamp), nil
	default:
		return 0, fmt.Errorf("ExecuteTime: CtrtDataResp.Val is %T but float64 was expected", timestamp)
	}
}

func NewDBKeyVOptionCtrtForExecuteDeadline() Bytes {
	return STATE_VAR_V_OPTION_EXECUTE_DEADLINE.Serialize()
}

func (v *VOptionCtrt) ExecuteDeadline() (VSYSTimestamp, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionCtrtForExecuteDeadline())
	if err != nil {
		return 0, fmt.Errorf("ExecuteDeadline: %w", err)
	}

	switch timestamp := resp.Val.(type) {
	case float64:
		return VSYSTimestamp(timestamp), nil
	default:
		return 0, fmt.Errorf("ExecuteDeadline: CtrtDataResp.Val is %T but float64 was expected", timestamp)
	}
}

func NewDBKeyVOptionForOptionStatus() Bytes {
	return STATE_VAR_V_OPTION_OPTION_STATUS.Serialize()
}

func (v *VOptionCtrt) OptionStatus() (bool, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForOptionStatus())
	if err != nil {
		return false, fmt.Errorf("OptionStatus: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("OptionStatus: CtrtDataResp.Val is %T but string was expected", val)
	}
}

func NewDBKeyVOptionForMaxIssueNum() Bytes {
	return STATE_VAR_V_OPTION_MAX_ISSUE_NUM.Serialize()
}

func (v *VOptionCtrt) MaxIssueNum() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForMaxIssueNum())
	if err != nil {
		return nil, fmt.Errorf("MaxIssueNum: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.BaseTokUnit()
		if err != nil {
			return nil, fmt.Errorf("MaxIssueNum: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("MaxIssueNum: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForReservedOption() Bytes {
	return STATE_VAR_V_OPTION_RESERVED_OPTION.Serialize()
}

func (v *VOptionCtrt) ReservedOption() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForReservedOption())
	if err != nil {
		return nil, fmt.Errorf("ReservedOption: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.OptionTokUnit()
		if err != nil {
			return nil, fmt.Errorf("ReservedOption: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("ReservedOption: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForReservedProof() Bytes {
	return STATE_VAR_V_OPTION_RESERVED_PROOF.Serialize()
}

func (v *VOptionCtrt) ReservedProof() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForReservedProof())
	if err != nil {
		return nil, fmt.Errorf("ReservedProof: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.ProofTokUnit()
		if err != nil {
			return nil, fmt.Errorf("ReservedProof: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("ReservedProof: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForPrice() Bytes {
	return STATE_VAR_V_OPTION_PRICE.Serialize()
}

func (v *VOptionCtrt) Price() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForPrice())
	if err != nil {
		return nil, fmt.Errorf("Price: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		return NewToken(Amount(amount), 1), nil
	default:
		return nil, fmt.Errorf("Price: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForPriceUnit() Bytes {
	return STATE_VAR_V_OPTION_PRICE_UNIT.Serialize()
}

func (v *VOptionCtrt) PriceUnit() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForPriceUnit())
	if err != nil {
		return nil, fmt.Errorf("PriceUnit: %w", err)
	}
	switch priceUnit := resp.Val.(type) {
	case float64:
		return NewToken(Amount(priceUnit), 1), nil
	default:
		return nil, fmt.Errorf("PriceUnit: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForTokenLocked() Bytes {
	return STATE_VAR_V_OPTION_TOKEN_LOCKED.Serialize()
}

func (v *VOptionCtrt) TokenLocked() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForTokenLocked())
	if err != nil {
		return nil, fmt.Errorf("TokenLocked: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("TokenLocked: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("TokenLocked: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionForTokenCollected() Bytes {
	return STATE_VAR_V_OPTION_TOKEN_COLLECTED.Serialize()
}

func (v *VOptionCtrt) TokenCollected() (*Token, error) {
	resp, err := v.QueryDBKey(NewDBKeyVOptionForTokenCollected())
	if err != nil {
		return nil, fmt.Errorf("TokenCollected: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.BaseTokUnit()
		if err != nil {
			return nil, fmt.Errorf("TokenCollected: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("TokenCollected: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionCtrtForBaseTokenBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVOptionForBaseTokenBalance")
	}
	return NewStateMap(STATE_MAP_IDX_V_OPTION_BASE_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

func (v *VOptionCtrt) GetBaseTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVOptionCtrtForBaseTokenBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokBal: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.BaseTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetBaseTokBal: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetBaseTokBal: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionCtrtForTargetTokBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVOptionForTargetTokBalance")
	}
	return NewStateMap(STATE_MAP_IDX_V_OPTION_TARGET_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

func (v *VOptionCtrt) GetTargetTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVOptionCtrtForTargetTokBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokBal: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetTargetTokBal: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetTargetTokBal: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionCtrtForOptionTokBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVOptionForOptionTokBalance")
	}
	return NewStateMap(STATE_MAP_IDX_V_OPTION_OPTION_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

func (v *VOptionCtrt) GetOptionTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVOptionCtrtForOptionTokBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetOptionTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetOptionTokBal: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.OptionTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetOptionTokBal: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetOptionTokBal: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVOptionCtrtForProofTokBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVOptionForProofTokBalance")
	}
	return NewStateMap(STATE_MAP_IDX_V_OPTION_PROOF_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

func (v *VOptionCtrt) GetProofTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVOptionCtrtForProofTokBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetProofTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetProofTokBal: %w", err)
	}
	switch amount := resp.Val.(type) {
	case float64:
		unit, err := v.ProofTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetProofTokBal: %w", err)
		}
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetProofTokBal: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}
