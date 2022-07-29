package vsys

import "fmt"

// VStableSwapCtrt is the struct for VSYS Stable Swap Contract.
type VStableSwapCtrt struct {
	*Ctrt
	baseTokId     *TokenId
	targetTokId   *TokenId
	baseTokCtrt   BaseTokCtrt
	targetTokCtrt BaseTokCtrt
}

// NewVStableSwapCtrt creates instance of VStableSwapCtrt from given contract id.
func NewVStableSwapCtrt(ctrtId string, chain *Chain) (*VStableSwapCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewVStableSwapCtrt: %w", err)
	}

	return &VStableSwapCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

// RegisterVStableSwapCtrt registers a Stable Swap Contract.
func RegisterVStableSwapCtrt(
	by *Account,
	baseTokId, targetTokId string,
	maxOrderPerUser int,
	basePriceUnit int,
	targetPriceUnit int,
	ctrtDescription string,
) (*VStableSwapCtrt, error) {
	cm, err := NewCtrtMetaForVStableSwapCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterVStableSwapCtrt: %w", err)
	}

	baseTokIdMd, err := NewTokenIdFromB58Str(baseTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVStableSwapCtrt: %w", err)
	}
	targetTokIdMd, err := NewTokenIdFromB58Str(targetTokId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVStableSwapCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			NewDeTokenId(baseTokIdMd),
			NewDeTokenId(targetTokIdMd),
			NewDeAmount(Amount(maxOrderPerUser)),
			NewDeAmount(Amount(basePriceUnit)),
			NewDeAmount(Amount(targetPriceUnit)),
		},
		cm,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterVStableSwap: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterVStableSwap: %w", err)
	}

	return &VStableSwapCtrt{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

// Supersede transfers the ownership of the contract to another account.
func (v *VStableSwapCtrt) Supersede(by *Account, newOwner string, attachment string) (*BroadcastExecuteTxResp, error) {
	newOwnerMd, err := NewAddrFromB58Str(newOwner)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_SUPERSEDE,
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

// SetOrder creates the order.
func (v *VStableSwapCtrt) SetOrder(
	by *Account,
	feeBase, feeTarget, minBase, maxBase, minTarget, maxTarget, priceBase, priceTarget, baseDeposit, targetDeposit float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	baseUnit, err := v.BaseTokUnit()
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}

	basePriceUnit, err := v.BasePriceUnit()
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	targetPriceUnit, err := v.TargetPriceUnit()
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}

	de1, err := NewDeAmountForTokAmount(feeBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de2, err := NewDeAmountForTokAmount(feeTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de3, err := NewDeAmountForTokAmount(minBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de4, err := NewDeAmountForTokAmount(maxBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de5, err := NewDeAmountForTokAmount(minTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de6, err := NewDeAmountForTokAmount(maxTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de7, err := NewDeAmountForTokAmount(priceBase, uint64(basePriceUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de8, err := NewDeAmountForTokAmount(priceTarget, uint64(targetPriceUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de9, err := NewDeAmountForTokAmount(baseDeposit, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	de10, err := NewDeAmountForTokAmount(targetDeposit, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_SET_ORDER,
		DataStack{
			de1,
			de2,
			de3,
			de4,
			de5,
			de6,
			de7,
			de8,
			de9,
			de10,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SetOrder: %w", err)
	}
	return resp, nil
}

// UpdateOrder updates the order settings.
func (v *VStableSwapCtrt) UpdateOrder(
	by *Account,
	feeBase, feeTarget, minBase, maxBase, minTarget, maxTarget, priceBase, priceTarget, baseDeposit, targetDeposit float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	baseUnit, err := v.BaseTokUnit()
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}

	basePriceUnit, err := v.BasePriceUnit()
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	targetPriceUnit, err := v.TargetPriceUnit()
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}

	de1, err := NewDeAmountForTokAmount(feeBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de2, err := NewDeAmountForTokAmount(feeTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de3, err := NewDeAmountForTokAmount(minBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de4, err := NewDeAmountForTokAmount(maxBase, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de5, err := NewDeAmountForTokAmount(minTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de6, err := NewDeAmountForTokAmount(maxTarget, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de7, err := NewDeAmountForTokAmount(priceBase, uint64(basePriceUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de8, err := NewDeAmountForTokAmount(priceTarget, uint64(targetPriceUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de9, err := NewDeAmountForTokAmount(baseDeposit, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	de10, err := NewDeAmountForTokAmount(targetDeposit, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_SET_ORDER,
		DataStack{
			de1,
			de2,
			de3,
			de4,
			de5,
			de6,
			de7,
			de8,
			de9,
			de10,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: %w", err)
	}
	return resp, nil
}

// OrderDeposit locks the tokens.
func (v *VStableSwapCtrt) OrderDeposit(
	by *Account,
	orderId string,
	baseDeposit, targetDeposit float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	baseUnit, err := v.BaseTokUnit()
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}

	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}
	deBase, err := NewDeAmountForTokAmount(baseDeposit, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}
	deTarget, err := NewDeAmountForTokAmount(targetDeposit, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_ORDER_DEPOSIT,
		DataStack{NewDeBytes(b), deBase, deTarget},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("OrderDeposit: %w", err)
	}
	return resp, nil
}

// OrderWithdraw unlocks the tokens.
func (v *VStableSwapCtrt) OrderWithdraw(
	by *Account,
	orderId string,
	baseWithdraw, targetWithdraw float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	baseUnit, err := v.BaseTokUnit()
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}
	targetUnit, err := v.TargetTokUnit()
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}

	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}
	deBase, err := NewDeAmountForTokAmount(baseWithdraw, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}
	deTarget, err := NewDeAmountForTokAmount(targetWithdraw, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_ORDER_WITHDRAW,
		DataStack{NewDeBytes(b), deBase, deTarget},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("OrderWithdraw: %w", err)
	}
	return resp, nil
}

// CloseOrder closes the order.
func (v *VStableSwapCtrt) CloseOrder(by *Account, orderId, attachment string) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("CloseOrder: %w", err)
	}
	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_CLOSE_ORDER,
		DataStack{NewDeBytes(b)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("CloseOrder: %w", err)
	}
	return resp, nil
}

// SwapBaseToTarget swaps base token to target token.
func (v *VStableSwapCtrt) SwapBaseToTarget(
	by *Account,
	orderId string,
	amount, swapFee, price float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	baseUnit, err := v.BaseTokUnit()
	basePriceUnit, err := v.BasePriceUnit()

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapBaseToTarget: %w", err)
	}
	deFee, err := NewDeAmountForTokAmount(swapFee, uint64(baseUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapBaseToTarget: %w", err)
	}
	dePrice, err := NewDeAmountForTokAmount(price, uint64(basePriceUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapBaseToTarget: %w", err)
	}
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("SwapBaseToTarget: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_SWAP_BASE_TO_TARGET,
		DataStack{
			NewDeBytes(b),
			deAmount,
			deFee,
			dePrice,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapBaseToTarget: %w", err)
	}
	return resp, nil
}

// SwapTargetToBase swaps target token to base token.
func (v *VStableSwapCtrt) SwapTargetToBase(
	by *Account,
	orderId string,
	amount, swapFee, price float64,
	deadline int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	targetUnit, err := v.TargetTokUnit()
	targetPriceUnit, err := v.TargetPriceUnit()

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapTargetToBase: %w", err)
	}
	deFee, err := NewDeAmountForTokAmount(swapFee, uint64(targetUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapTargetToBase: %w", err)
	}
	dePrice, err := NewDeAmountForTokAmount(price, uint64(targetPriceUnit))
	if err != nil {
		return nil, fmt.Errorf("SwapTargetToBase: %w", err)
	}
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("SwapTargetToBase: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_STABLE_SWAP_SWAP_TARGET_TO_BASE,
		DataStack{
			NewDeBytes(b),
			deAmount,
			deFee,
			dePrice,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(deadline)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SwapTargetToBase: %w", err)
	}
	return resp, nil
}

// Maker queries & returns the maker of the contract.
func (v *VStableSwapCtrt) Maker() (*Addr, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVSwapForMaker(),
	)
	if err != nil {
		return nil, fmt.Errorf("Maker: %w", err)
	}
	switch addrB58 := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("Maker: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Maker: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

// BaseTokUnit queries & returns the base token id.
func (v *VStableSwapCtrt) BaseTokUnit() (Unit, error) {
	tc, err := v.BaseTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("BaseTokUnit: %w", err)
	}
	return tc.Unit()
}

// TargetTokUnit queries & returns the target token id.
func (v *VStableSwapCtrt) TargetTokUnit() (Unit, error) {
	tc, err := v.TargetTokCtrt()
	if err != nil {
		return 0, fmt.Errorf("TargetTokUnit: %w", err)
	}
	return tc.Unit()
}

// BaseTokCtrt returns the token contract instance for base token.
func (v *VStableSwapCtrt) BaseTokCtrt() (BaseTokCtrt, error) {
	if v.baseTokCtrt == nil {
		// Note that this BaseTokId() is not related to Base token of VStableSwap
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

// TargetTokCtrt returns the token contract instance for target token.
func (v *VStableSwapCtrt) TargetTokCtrt() (BaseTokCtrt, error) {
	if v.targetTokCtrt == nil {
		// Note that this BaseTokId() is not related to Base token of VStableSwap
		targetTokId, err := v.BaseTokId()
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

func NewDBKeyVStableSwapBaseTokId() Bytes {
	return STATE_VAR_V_STABLE_SWAP_BASE_TOKEN_ID.Serialize()
}

func NewDBKeyVStableSwapTargetTokId() Bytes {
	return STATE_VAR_V_STABLE_SWAP_TARGET_TOKEN_ID.Serialize()
}

// BaseTokId returns token id of base token.
func (v *VStableSwapCtrt) BaseTokId() (*TokenId, error) {
	if v.baseTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVStableSwapBaseTokId())
		if err != nil {
			return nil, fmt.Errorf("BaseTokId: %w", err)
		}
		switch tokId := resp.Val.(type) {
		case string:
			tokIdMd, err := NewTokenIdFromB58Str(tokId)
			if err != nil {
				return nil, fmt.Errorf("BaseTokId: %w", err)
			}
			v.baseTokId = tokIdMd
			return tokIdMd, nil
		default:
			return nil, fmt.Errorf("BaseTokId: CtrtDataResp.Val is %T but string was expected", tokId)
		}
	}
	return v.baseTokId, nil
}

// TargetTokId returns token id of target token.
func (v *VStableSwapCtrt) TargetTokId() (*TokenId, error) {
	if v.targetTokId == nil {
		resp, err := v.QueryDBKey(NewDBKeyVStableSwapTargetTokId())
		if err != nil {
			return nil, fmt.Errorf("TargetTokId: %w", err)
		}
		switch tokId := resp.Val.(type) {
		case string:
			tokIdMd, err := NewTokenIdFromB58Str(tokId)
			if err != nil {
				return nil, fmt.Errorf("TargetTokId: %w", err)
			}
			v.targetTokId = tokIdMd
			return tokIdMd, nil
		default:
			return nil, fmt.Errorf("TargetTokId: CtrtDataResp.Val is %T but string was expected", tokId)
		}
	}
	return v.targetTokId, nil
}

func NewDBKeyVStableSwapBasePriceUnit() Bytes {
	return STATE_VAR_V_STABLE_SWAP_UNIT_PRICE_BASE.Serialize()
}
func NewDBKeyVStableSwapTargetPriceUnit() Bytes {
	return STATE_VAR_V_STABLE_SWAP_UNIT_PRICE_TARGET.Serialize()
}

// BasePriceUnit queries & returns the price unit of base token.
func (v *VStableSwapCtrt) BasePriceUnit() (Unit, error) {
	resp, err := v.QueryDBKey(NewDBKeyVStableSwapBasePriceUnit())
	if err != nil {
		return 0, fmt.Errorf("BasePriceUnit: %w", err)
	}
	switch unit := resp.Val.(type) {
	case float64:
		return Unit(unit), nil
	default:
		return 0, fmt.Errorf("BasePriceUnit: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

// TargetPriceUnit  queries & returns the price unit of target token.
func (v *VStableSwapCtrt) TargetPriceUnit() (Unit, error) {
	resp, err := v.QueryDBKey(NewDBKeyVStableSwapTargetPriceUnit())
	if err != nil {
		return 0, fmt.Errorf("TargetPriceUnit: %w", err)
	}
	switch unit := resp.Val.(type) {
	case float64:
		return Unit(unit), nil
	default:
		return 0, fmt.Errorf("TargetPriceUnit: CtrtDataResp.Val is %T but float64 was expected", tokId)
	}
}

func NewDBKeyVStableSwapMaxOrderPerUser() Bytes {
	return STATE_VAR_V_STABLE_SWAP_MAX_ORDER_PER_USER.Serialize()
}

// MaxOrderPerUser queries & returns the maximum order number that each user can create.
func (v *VStableSwapCtrt) MaxOrderPerUser() (int, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVStableSwapMaxOrderPerUser(),
	)
	if err != nil {
		return 0, fmt.Errorf("Maker: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		return int(val), nil
	default:
		return 0, fmt.Errorf("Maker: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapBaseTokenBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyForBaseTokenBalance: %w", err)
	}

	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_BASE_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

// GetBaseTokBal queries & returns the balance of the available base tokens.
func (v *VStableSwapCtrt) GetBaseTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapBaseTokenBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokBal: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.BaseTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetBaseTokBal: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetBaseTokBal: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapTargetTokenBalance(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyForTargetTokenBalance: %w", err)
	}

	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_TARGET_TOKEN_BALANCE, NewDeAddr(addrMd)).Serialize(), nil
}

// GetTargetTokBal queries & returns the balance of the available target tokens.
func (v *VStableSwapCtrt) GetTargetTokBal(addr string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapTargetTokenBalance(addr)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokBal: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokBal: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetTargetTokBal: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetTargetTokBal: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapGetUserOrders(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyForTargetTokenBalance: %w", err)
	}

	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_USER_ORDERS, NewDeAddr(addrMd)).Serialize(), nil
}

// GetUserOrders queries & returns the number of user orders.
func (v *VStableSwapCtrt) GetUserOrders(addr string) (int, error) {
	dbKey, err := NewDBKeyVStableSwapGetUserOrders(addr)
	if err != nil {
		return 0, fmt.Errorf("GetUserOrders: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return 0, fmt.Errorf("GetUserOrders: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		return int(val), nil
	default:
		return 0, fmt.Errorf("GetUserOrders: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapOrderOwner(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapOrderOwner: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_ORDER_OWNER, NewDeBytes(b)).Serialize(), nil
}

// GetOrderOwner queries & returns the address of the order owner.
func (v *VStableSwapCtrt) GetOrderOwner(orderId string) (*Addr, error) {
	dbKey, err := NewDBKeyVStableSwapOrderOwner(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderOwner: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetOrderOwner: %w", err)
	}
	switch addrB58 := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("GetOrderOwner: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("GetOrderOwner: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

func NewDBKeyVStableSwapFeeBase(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapFeeBase: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_FEE_BASE, NewDeBytes(b)).Serialize(), nil
}

// GetFeeBase queries & returns the base fee.
func (v *VStableSwapCtrt) GetFeeBase(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapFeeBase(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetFeeBase: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetFeeBase: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetFeeBase: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetFeeBase: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapFeeTarget(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapFeeTarget: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_FEE_TARGET, NewDeBytes(b)).Serialize(), nil
}

// GetFeeTarget queries and returns target fee.
func (v *VStableSwapCtrt) GetFeeTarget(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapFeeTarget(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetFeeTarget: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetFeeTarget: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetFeeTarget: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetFeeTarget: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapMinBase(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapMinBase: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_MIN_BASE, NewDeBytes(b)).Serialize(), nil
}

// GetMinBase queries & returns the minimum amount of base token.
func (v *VStableSwapCtrt) GetMinBase(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapMinBase(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetMinBase: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetMinBase: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetMinBase: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetMinBase: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapMinTarget(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapMinTarget: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_MIN_TARGET, NewDeBytes(b)).Serialize(), nil
}

// GetMinTarget  queries & returns the minimum amount of target token.
func (v *VStableSwapCtrt) GetMinTarget(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapMinTarget(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetMinTarget: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetMinTarget: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetMinTarget: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetMinTarget: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapMaxBase(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapMaxBase: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_MAX_BASE, NewDeBytes(b)).Serialize(), nil
}

// GetMaxBase  queries & returns the maximum amount of base token.
func (v *VStableSwapCtrt) GetMaxBase(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapMaxBase(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetMaxBase: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetMaxBase: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetMaxBase: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetMaxBase: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapMaxTarget(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapMaxTarget: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_MAX_TARGET, NewDeBytes(b)).Serialize(), nil
}

// GetMaxTarget  queries & returns the maximum amount of target token.
func (v *VStableSwapCtrt) GetMaxTarget(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapMaxTarget(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetMaxTarget: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetMaxTarget: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetMaxTarget: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetMaxTarget: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}
func NewDBKeyVStableSwapPriceBase(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapPriceBase: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_PRICE_BASE, NewDeBytes(b)).Serialize(), nil
}

// GetPriceBase  queries & returns the price of base token.
func (v *VStableSwapCtrt) GetPriceBase(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapPriceBase(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetPriceBase: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetPriceBase: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetPriceBase: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetPriceBase: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapPriceTarget(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapPriceTarget: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_PRICE_TARGET, NewDeBytes(b)).Serialize(), nil
}

// GetPriceTarget  queries & returns the price of the target token.
func (v *VStableSwapCtrt) GetPriceTarget(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapPriceTarget(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetPriceTarget: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetPriceTarget: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetPriceTarget: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetPriceTarget: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapBaseTokenLocked(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapBaseTokenLocked: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_BASE_TOKEN_LOCKED, NewDeBytes(b)).Serialize(), nil
}

// GetBaseTokLocked  queries & returns the amount of locked base tokens.
func (v *VStableSwapCtrt) GetBaseTokLocked(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapBaseTokenLocked(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokLocked: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetBaseTokLocked: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetBaseTokLocked: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetBaseTokLocked: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapTargetTokenLocked(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapTargetTokenLocked: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_TARGET_TOKEN_LOCKED, NewDeBytes(b)).Serialize(), nil
}

// GetTargetTokLocked  queries & returns the amount of locked target tokens.
func (v *VStableSwapCtrt) GetTargetTokLocked(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVStableSwapTargetTokenLocked(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokLocked: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetTargetTokLocked: %w", err)
	}
	switch val := resp.Val.(type) {
	case float64:
		unit, err := v.TargetTokUnit()
		if err != nil {
			return nil, fmt.Errorf("GetTargetTokLocked: %w", err)
		}
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("GetTargetTokLocked: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyVStableSwapOrderStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVStableSwapOrderStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_STABLE_SWAP_ORDER_STATUS, NewDeBytes(b)).Serialize(), nil
}

// GetOrderStatus  queries & returns the status of the order.
func (v *VStableSwapCtrt) GetOrderStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVStableSwapOrderStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderStatus: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("GetOrderStatus: CtrtDataResp.Val is %T but string was expected", val)
	}
}
