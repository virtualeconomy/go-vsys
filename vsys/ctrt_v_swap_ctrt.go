package vsys

import "fmt"

type VSwapCtrt struct {
	*Ctrt
	tokACtrt   BaseTokCtrt
	tokBCtrt   BaseTokCtrt
	liqTokCtrt BaseTokCtrt
	tokAId     *TokenId
	tokBId     *TokenId
	liqTokId   *TokenId
}

func NewVSwapCtrt(ctrtId string, chain *Chain) (*VSwapCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewVSwapCtrt: %w", err)
	}

	return &VSwapCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

func RegisterVSwapCtrt(
	by *Account,
	tokenAId string,
	tokenBId string,
	liqTokId string,
	minLiq int,
	ctrtDescription string,
) (*VSwapCtrt, error) {
	cm, err := NewCtrtMetaForVSwapCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}

	tokenAIdMd, err := NewTokenIdFromB58Str(tokenAId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}
	tokenBIdMd, err := NewTokenIdFromB58Str(tokenBId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}
	liqTokIdMd, err := NewTokenIdFromB58Str(liqTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			NewDeTokenId(tokenAIdMd),
			NewDeTokenId(tokenBIdMd),
			NewDeTokenId(liqTokIdMd),
			NewDeAmount(Amount(minLiq)),
		},
		cm,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterVSwapCtrt: %w", err)
	}

	return &VSwapCtrt{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

func NewDBKeyVSwapForTokenAId() Bytes {
	return STATE_VAR_V_SWAP_TOKEN_A_ID.Serialize()
}

func NewDBKeyVSwapForTokenBId() Bytes {
	return STATE_VAR_V_SWAP_TOKEN_B_ID.Serialize()
}

func NewDBKeyVSwapForLiqTokId() Bytes {
	return STATE_VAR_V_SWAP_LIQUIDITY_TOKEN_ID.Serialize()
}
func NewDBKeyVSwapForTotalLiqTokSupply() Bytes {
	return STATE_VAR_V_SWAP_TOTAL_SUPPLY.Serialize()
}

func NewDBKeyVSwapForLiqTokLeft() Bytes {
	return STATE_VAR_V_SWAP_LIQUIDITY_TOKEN_LEFT.Serialize()
}

func NewDBKeyVSwapForTokAReserved() Bytes {
	return STATE_VAR_V_SWAP_TOKEN_A_RESERVED.Serialize()
}

func NewDBKeyVSwapForTokBReserved() Bytes {
	return STATE_VAR_V_SWAP_TOKEN_B_RESERVED.Serialize()
}

func NewDBKeyVSwapForMinLiq() Bytes {
	return STATE_VAR_V_SWAP_MINIMUM_LIQUIDITY.Serialize()
}

func NewDBKeyVSwapForSwapStatus() Bytes {
	return STATE_VAR_V_SWAP_SWAP_STATUS.Serialize()
}

func NewDBKeyVSwapForMaker() Bytes {
	return STATE_VAR_V_SWAP_MAKER.Serialize()
}

func NewDBKeyVSwapForGetTokABal(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVSwapForGetTokBal: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_SWAP_TOKEN_A_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}
func NewDBKeyVSwapForGetTokBBal(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVSwapForGetTokBal: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_SWAP_TOKEN_B_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}
func NewDBKeyVSwapForGetLiqTokBal(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVSwapForGetTokBal: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_SWAP_LIQUIDITY_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

func (vs *VSwapCtrt) TokAUnit() (Unit, error) {
	tc, err := vs.TokACtrt()
	if err != nil {
		return 0, fmt.Errorf("TokAUnit: %w", err)
	}
	return tc.Unit()
}

func (vs *VSwapCtrt) TokBUnit() (Unit, error) {
	tc, err := vs.TokBCtrt()
	if err != nil {
		return 0, fmt.Errorf("TokBUnit: %w", err)
	}
	return tc.Unit()
}

func (vs *VSwapCtrt) LiqTokUnit() (Unit, error) {
	tc, err := vs.LiqTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("TokBUnit: %w", err)
	}
	return tc.Unit()
}

func (vs *VSwapCtrt) TokACtrt() (BaseTokCtrt, error) {
	if vs.tokACtrt == nil {
		tokAId, err := vs.TokAId()
		if err != nil {
			return nil, fmt.Errorf("TokACtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(tokAId, vs.Chain)
		if err != nil {
			return nil, fmt.Errorf("TokACtrt: %w", err)
		}
		vs.tokACtrt = tc
	}
	return vs.tokACtrt, nil
}
func (vs *VSwapCtrt) TokBCtrt() (BaseTokCtrt, error) {
	if vs.tokBCtrt == nil {
		tokAId, err := vs.TokBId()
		if err != nil {
			return nil, fmt.Errorf("TokBCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(tokAId, vs.Chain)
		if err != nil {
			return nil, fmt.Errorf("TokBCtrt: %w", err)
		}
		vs.tokBCtrt = tc
	}
	return vs.tokBCtrt, nil
}

func (vs *VSwapCtrt) LiqTokCtrt() (BaseTokCtrt, error) {
	if vs.liqTokCtrt == nil {
		liqTokId, err := vs.LiqTokId()
		if err != nil {
			return nil, fmt.Errorf("LiqTokCtrt: %w", err)
		}
		tc, err := GetCtrtFromTokId(liqTokId, vs.Chain)
		if err != nil {
			return nil, fmt.Errorf("LiqTokCtrt: %w", err)
		}
		vs.liqTokCtrt = tc
	}
	return vs.liqTokCtrt, nil
}

func (vs *VSwapCtrt) TokAId() (*TokenId, error) {

	if vs.tokAId == nil {
		resp, err := vs.QueryDBKey(NewDBKeyVSwapForTokenAId())
		if err != nil {
			return nil, fmt.Errorf("TokAId: %w", err)
		}

		tokId, err := ctrtDataRespToTokenId(resp)
		if err != nil {
			return nil, fmt.Errorf("TokAId: %w", err)
		}
		vs.tokAId = tokId
	}
	return vs.tokAId, nil
}

func (vs *VSwapCtrt) TokBId() (*TokenId, error) {
	if vs.tokBId == nil {
		resp, err := vs.QueryDBKey(NewDBKeyVSwapForTokenBId())
		if err != nil {
			return nil, fmt.Errorf("TokBId: %w", err)
		}

		tokId, err := ctrtDataRespToTokenId(resp)
		if err != nil {
			return nil, fmt.Errorf("TokBId: %w", err)
		}
		vs.tokBId = tokId
	}
	return vs.tokBId, nil
}

func (vs *VSwapCtrt) LiqTokId() (*TokenId, error) {
	if vs.liqTokId == nil {
		resp, err := vs.QueryDBKey(NewDBKeyVSwapForLiqTokId())
		if err != nil {
			return nil, fmt.Errorf("LiqTokId: %w", err)
		}

		tokId, err := ctrtDataRespToTokenId(resp)
		if err != nil {
			return nil, fmt.Errorf("LiqTokId: %w", err)
		}
		vs.liqTokId = tokId
	}
	return vs.liqTokId, nil
}

func (vs *VSwapCtrt) Maker() (*Addr, error) {
	resp, err := vs.QueryDBKey(
		NewDBKeyVSwapForMaker(),
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

func (vs *VSwapCtrt) IsSwapActive() (bool, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForSwapStatus())
	if err != nil {
		return false, fmt.Errorf("IsSwapActive: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("IsSwapActive: CtrtDataResp.Val is %T but string was expected", tokId)
	}
}

func (vs *VSwapCtrt) MinLiq() (*Token, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForMinLiq())
	if err != nil {
		return nil, fmt.Errorf("MinLiq: %w", err)
	}
	unit, err := vs.LiqTokUnit()
	if err != nil {
		return nil, fmt.Errorf("MinLiq: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("MinLiq: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) TokAReserved() (*Token, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForTokAReserved())
	if err != nil {
		return nil, fmt.Errorf("TokAReserved: %w", err)
	}
	unit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("TokAReserved: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("TokAReserved: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) TokBReserved() (*Token, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForTokBReserved())
	if err != nil {
		return nil, fmt.Errorf("TokBReserved: %w", err)
	}
	unit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("TokBReserved: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("TokBReserved: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) TotalLiqTokSupply() (*Token, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForTotalLiqTokSupply())
	if err != nil {
		return nil, fmt.Errorf("TotalLiqTokSupply: %w", err)
	}
	unit, err := vs.LiqTokUnit()
	if err != nil {
		return nil, fmt.Errorf("TotalLiqSupply: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("TotalLiqSupply: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) LiqTokLeft() (*Token, error) {
	resp, err := vs.QueryDBKey(NewDBKeyVSwapForLiqTokLeft())
	if err != nil {
		return nil, fmt.Errorf("LiqTokLeft: %w", err)
	}
	unit, err := vs.LiqTokUnit()
	if err != nil {
		return nil, fmt.Errorf("LiqTokLeft: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("LiqTokLeft: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) GetTokABal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVSwapForGetTokABal(addr)
	if err != nil {
		return nil, fmt.Errorf("GetTokABal: %w", err)
	}
	resp, err := vs.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetTokABal: %w", err)
	}
	unit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("GetTokABal: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetTokABal: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) GetTokBBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVSwapForGetTokBBal(addr)
	if err != nil {
		return nil, fmt.Errorf("GetTokBBal: %w", err)
	}
	resp, err := vs.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetTokBBal: %w", err)
	}
	unit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("GetTokBBal: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetTokBBal: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) GetLiqTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVSwapForGetLiqTokBal(addr)
	if err != nil {
		return nil, fmt.Errorf("GetLiqTokBal: %w", err)
	}
	resp, err := vs.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetLiqTokBal: %w", err)
	}
	unit, err := vs.LiqTokUnit()
	if err != nil {
		return nil, fmt.Errorf("GetLiqTokBal: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetLiqTokBal: %w", err)
	}
	return tok, nil
}

func (vs *VSwapCtrt) Supersede(by *Account, newOwner, attachment string) (*BroadcastExecuteTxResp, error) {
	newOwnerMd, err := NewAddrFromB58Str(newOwner)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SUPERSEDE,
		DataStack{
			NewDeAddr(newOwnerMd),
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

func (vs *VSwapCtrt) SetSwap(by *Account, amountA, amountB float64, attachment string) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("SetSwap: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("SetSwap: %w", err)
	}

	deAmountA, err := NewDeAmountForTokAmount(amountA, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("SetSwap: %w", err)
	}
	deAmountB, err := NewDeAmountForTokAmount(amountB, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("SetSwap: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SET_SWAP,
		DataStack{
			deAmountA,
			deAmountB,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SetSwap: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) AddLiquidity(
	by *Account,
	amountA, amountB, amountAmin, amountBmin float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}

	deAmountA, err := NewDeAmountForTokAmount(amountA, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}
	deAmountB, err := NewDeAmountForTokAmount(amountB, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}
	deAmountAmin, err := NewDeAmountForTokAmount(amountAmin, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}
	deAmountBmin, err := NewDeAmountForTokAmount(amountBmin, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_ADD_LIQUIDITY,
		DataStack{
			deAmountA,
			deAmountB,
			deAmountAmin,
			deAmountBmin,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("AddLiquidity: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) RemoveLiquidity(
	by *Account,
	amountLiq, amountAmin, amountBmin float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}
	liqUnit, err := vs.LiqTokUnit()
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}

	deAmountLiq, err := NewDeAmountForTokAmount(amountLiq, uint64(liqUnit))
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}
	deAmountAmin, err := NewDeAmountForTokAmount(amountAmin, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}
	deAmountBmin, err := NewDeAmountForTokAmount(amountBmin, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_REMOVE_LIQUIDITY,
		DataStack{
			deAmountLiq,
			deAmountAmin,
			deAmountBmin,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RemoveLiquidity: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) SwapBForExactA(
	by *Account,
	amountA, amountBMax float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapBForExactA: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapBForExactA: %w", err)
	}

	deAmountA, err := NewDeAmountForTokAmount(amountA, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapBForExactA: %w", err)
	}
	deAmountBmax, err := NewDeAmountForTokAmount(amountBMax, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapBForExactA: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SWAP_B_FOR_EXACT_A,
		DataStack{deAmountA, deAmountBmax, NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapBForExactA: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) SwapExactBForA(
	by *Account,
	amountAmin, amountB float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapExactBForA: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapExactBForA: %w", err)
	}

	deAmountAmin, err := NewDeAmountForTokAmount(amountAmin, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapExactBForA: %w", err)
	}
	deAmountB, err := NewDeAmountForTokAmount(amountB, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapExactBForA: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SWAP_EXACT_B_FOR_A,
		DataStack{deAmountAmin, deAmountB, NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapExactBForA: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) SwapAForExactB(
	by *Account,
	amountB, amountAMax float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapAForExactB: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapAForExactB: %w", err)
	}

	deAmountAMax, err := NewDeAmountForTokAmount(amountAMax, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapAForExactB: %w", err)
	}
	deAmountB, err := NewDeAmountForTokAmount(amountB, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapAForExactB: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SWAP_A_FOR_EXACT_B,
		DataStack{deAmountB, deAmountAMax, NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapAForExactB: %w", err)
	}
	return resp, nil
}

func (vs *VSwapCtrt) SwapExactAForB(
	by *Account,
	amountBmin, amountA float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	tokAUnit, err := vs.TokAUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapExactAForB: %w", err)
	}
	tokBUnit, err := vs.TokBUnit()
	if err != nil {
		return nil, fmt.Errorf("SwapExactAForB: %w", err)
	}

	deAmountA, err := NewDeAmountForTokAmount(amountA, uint64(tokBUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapExactAForB: %w", err)
	}
	deAmountBmin, err := NewDeAmountForTokAmount(amountBmin, uint64(tokAUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapExactAForB: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		vs.CtrtId,
		FUNC_IDX_V_SWAP_SWAP_EXACT_A_FOR_B,
		DataStack{deAmountBmin, deAmountA, NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapExactAForB: %w", err)
	}
	return resp, nil
}
