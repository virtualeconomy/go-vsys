package vsys

import "fmt"

type VEscrowCtrt struct {
	*Ctrt
	tokId   *TokenId
	tokCtrt BaseTokCtrt
}

func NewVEscrowCtrt(ctrtId string, chain *Chain) (*VEscrowCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewVEscrowCtrt: %w", err)
	}

	return &VEscrowCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
		tokId:   nil,
		tokCtrt: nil,
	}, nil
}

func RegisterVEscrowCtrt(
	by *Account,
	tokenId string,
	duration int64,
	judge_duration int64,
	ctrtDescription string,
) (*VEscrowCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForVEscrowCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterVEscrowCtrt: %w", err)
	}

	tokId, err := NewTokenIdFromB58Str(tokenId)
	if err != nil {
		return nil, fmt.Errorf("RegisterVEscrowCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{
			NewDeTokenId(tokId),
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(duration)),
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(judge_duration)),
		},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterVEscrowCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterVEscrowCtrt: %w", err)
	}

	return &VEscrowCtrt{
		&Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
		nil,
		nil,
	}, nil
}

func NewDBKeyVEscrowMaker() Bytes {
	return STATE_VAR_V_ESCROW_MAKER.Serialize()
}

func NewDBKeyVEscrowJudge() Bytes {
	return STATE_VAR_V_ESCROW_JUDGE.Serialize()
}

func NewDBKeyVEscrowTokId() Bytes {
	return STATE_VAR_V_ESCROW_TOKEN_ID.Serialize()
}

func NewDBKeyVEscrowDuration() Bytes {
	return STATE_VAR_V_ESCROW_DURATION.Serialize()
}
func NewDBKeyVEscrowJudgeDuration() Bytes {
	return STATE_VAR_V_ESCROW_JUDGE_DURATION.Serialize()
}

func NewDBKeyVEscrowContractBalance(addr *Addr) Bytes {
	return NewStateMap(
		STATE_MAP_IDX_V_ESCROW_CONTRACT_BALANCE,
		NewDeAddr(addr),
	).Serialize()
}

func NewDBKeyVEscrowOrderPayer(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderPayer: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_PAYER, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipient(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipient: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderAmount(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderAmount: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_AMOUNT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipientDeposit(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipientDeposit: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT_DEPOSIT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderJudgeDeposit(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderJudgeDeposit: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_JUDGE_DEPOSIT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderFee(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderFee: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_FEE, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipientAmount(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipientAmount: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT_AMOUNT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRefund(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRefund: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_REFUND, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipientRefund(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipientRefund: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT_REFUND, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowExpirationTime(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowExpirationTime: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_EXPIRATION_TIME, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_STATUS, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipientDepositStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipientDepositStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT_DEPOSIT_STATUS, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderJudgeDepositStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderJudgeDepositStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_JUDGE_DEPOSIT_STATUS, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderSubmitStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderSubmitStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_SUBMIT_STATUS, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderJudgeStatus(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderJudgeStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_JUDGE_STATUS, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderRecipientLockedAmount(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderRecipientLockedAmount: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_RECIPIENT_LOCKED_AMOUNT, NewDeBytes(b)).Serialize(), nil
}

func NewDBKeyVEscrowOrderJudgeLockedAmount(orderId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyVEscrowOrderJudgeLockedAmount: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_V_ESCROW_ORDER_JUDGE_LOCKED_AMOUNT, NewDeBytes(b)).Serialize(), nil
}

// Maker queries and returns maker Addr of V Escrow Contract.
func (v *VEscrowCtrt) Maker() (*Addr, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVEscrowMaker(),
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

// Judge queries and returns judge Addr of the contract.
func (v *VEscrowCtrt) Judge() (*Addr, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVEscrowJudge(),
	)
	if err != nil {
		return nil, fmt.Errorf("Judge: %w", err)
	}

	addr, err := ctrtDataRespToAddr(resp)
	if err != nil {
		return nil, fmt.Errorf("Judge: %w", err)
	}
	return addr, nil
}

// TokId queries and returns TokenId of the contract's token.
func (v *VEscrowCtrt) TokId() (*TokenId, error) {
	if v.tokId == nil {
		resp, err := v.QueryDBKey(
			NewDBKeyVEscrowTokId(),
		)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}

		tokId, err := ctrtDataRespToTokenId(resp)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		v.tokId = tokId
	}
	return v.tokId, nil
}

// Duration queries & returns the duration where the recipient can take actions in the contract.
func (v *VEscrowCtrt) Duration() (VSYSTimestamp, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVEscrowDuration(),
	)
	if err != nil {
		return 0, fmt.Errorf("Duration: %w", err)
	}
	ts, err := ctrtDataRespToVSYSTimestamp(resp)
	if err != nil {
		return 0, fmt.Errorf("Duration: %w", err)
	}
	return ts, nil
}

// JudgeDuration queries & returns the duration where the judge can take actions in the contract.
func (v *VEscrowCtrt) JudgeDuration() (VSYSTimestamp, error) {
	resp, err := v.QueryDBKey(
		NewDBKeyVEscrowJudgeDuration(),
	)
	if err != nil {
		return 0, fmt.Errorf("JudgeDuration: %w", err)
	}
	ts, err := ctrtDataRespToVSYSTimestamp(resp)
	if err != nil {
		return 0, fmt.Errorf("JudgeDuration: %w", err)
	}
	return ts, nil
}

// Unit queries and returns Unit of the token of contract.
func (v *VEscrowCtrt) Unit() (Unit, error) {
	if v.tokCtrt == nil {
		_, err := v.TokCtrt() // TokCtrt sets a.TokCtrt
		if err != nil {
			return 0, err
		}
	}
	return v.tokCtrt.Unit()
}

// GetCtrtBal queries & returns the balance of the token within this contract belonging to the user address.
func (v *VEscrowCtrt) GetCtrtBal(addr string) (*Token, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	resp, err := v.QueryDBKey(
		NewDBKeyVEscrowContractBalance(addrMd),
	)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}
	return tok, nil
}

// GetOrderPayer queries & returns the payer of the order.
func (v *VEscrowCtrt) GetOrderPayer(orderId string) (*Addr, error) {
	dbKey, err := NewDBKeyVEscrowOrderPayer(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderPayer: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderPayer: %w", err)
	}

	addr, err := ctrtDataRespToAddr(resp)
	if err != nil {
		return nil, fmt.Errorf("GetOrderPayer: %w", err)
	}
	return addr, nil
}

// GetOrderRecipient queries & returns the recipient of the order.
func (v *VEscrowCtrt) GetOrderRecipient(orderId string) (*Addr, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipient(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipient: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipient: %w", err)
	}

	addr, err := ctrtDataRespToAddr(resp)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipient: %w", err)
	}
	return addr, nil
}

// GetOrderAmount queries & returns the amount of the order.
func (v *VEscrowCtrt) GetOrderAmount(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderAmount(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderAmount: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderAmount: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderAmount: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderAmount: %w", err)
	}
	return tok, nil
}

// GetOrderRecipientDeposit queries & returns the amount the recipient should deposit in the order.
func (v *VEscrowCtrt) GetOrderRecipientDeposit(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipientDeposit(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientDeposit: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientDeposit: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientDeposit: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientDeposit: %w", err)
	}
	return tok, nil
}

// GetOrderJudgeDeposit queries & returns the amount the recipient should deposit in the order.
func (v *VEscrowCtrt) GetOrderJudgeDeposit(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderJudgeDeposit(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeDeposit: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeDeposit: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeDeposit: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeDeposit: %w", err)
	}
	return tok, nil
}

// GetOrderFee queries & returns the fee of the order.
func (v *VEscrowCtrt) GetOrderFee(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderFee(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderFee: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderFee: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderFee: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderFee: %w", err)
	}
	return tok, nil
}

// GetOrderRecipientAmount queries & returns how much the recipient will receive
// from the order if the order goes smoothly(i.e. work is submitted & approved).
// The recipient amount = order amount - order fee.
func (v *VEscrowCtrt) GetOrderRecipientAmount(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipientAmount(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientAmount: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientAmount: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientAmount: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientAmount: %w", err)
	}
	return tok, nil
}

// GetOrderRefund queries & returns the refund amount of the order.
// The refund amount means how much the payer will receive if the refund occurs.
// It is defined when the order is created.
func (v *VEscrowCtrt) GetOrderRefund(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderRefund(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRefund: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRefund: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderRefund: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRefund: %w", err)
	}
	return tok, nil
}

// GetOrderRecipientRefund queries & returns the recipient refund amount of the order.
// the recipient refund amount means how much the recipient will receive if the refund occurs.
// The recipient refund amount = The total deposit(order amount + judge deposit + recipient deposit) - payer refund
func (v *VEscrowCtrt) GetOrderRecipientRefund(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipientRefund(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientRefund: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRefund: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientRefund: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientRefund: %w", err)
	}
	return tok, nil
}

// GetOrderExpirationTime queries & returns the expiration time of the order.
func (v *VEscrowCtrt) GetOrderExpirationTime(orderId string) (VSYSTimestamp, error) {
	dbKey, err := NewDBKeyVEscrowExpirationTime(orderId)
	if err != nil {
		return 0, fmt.Errorf("NewDBKeyVEscrowExpirationTime: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return 0, fmt.Errorf("NewDBKeyVEscrowExpirationTime: %w", err)
	}
	ts, err := ctrtDataRespToVSYSTimestamp(resp)
	if err != nil {
		return 0, fmt.Errorf("GetOrderExpirationTime: %w", err)
	}
	return ts, nil
}

// GetOrderStatus queries & returns the status of the order.
// The order status means if the order is active.
// The order is considered active if it is created & it is NOT finished.
func (v *VEscrowCtrt) GetOrderStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVEscrowOrderStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderStatus: %w", err)
	}

	val, err := ctrtDataRespToBool(resp)
	if err != nil {
		return false, fmt.Errorf("GetOrderStatus: %w", err)
	}
	return val, nil
}

// GetOrderRecipientDepositStatus queries & returns the recipient deposit status of the order.
// The order recipient deposit status means if the recipient has deposited into the order.
func (v *VEscrowCtrt) GetOrderRecipientDepositStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipientDepositStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderRecipientDepositStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderRecipientDepositStatus: %w", err)
	}

	val, err := ctrtDataRespToBool(resp)
	if err != nil {
		return false, fmt.Errorf("GetOrderRecipientDepositStatus: %w", err)
	}
	return val, nil
}

// GetOrderJudgeDepositStatus  queries & returns the judge deposit status of the order.
// The order judge deposit status means if the judge has deposited into the order.
func (v *VEscrowCtrt) GetOrderJudgeDepositStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVEscrowOrderJudgeDepositStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeDepositStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeDepositStatus: %w", err)
	}

	val, err := ctrtDataRespToBool(resp)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeDepositStatus: %w", err)
	}
	return val, nil
}

// GetOrderSubmitStatus queries & returns the submit status of the order.
func (v *VEscrowCtrt) GetOrderSubmitStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVEscrowOrderSubmitStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderSubmitStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderSubmitStatus: %w", err)
	}

	val, err := ctrtDataRespToBool(resp)
	if err != nil {
		return false, fmt.Errorf("GetOrderSubmitStatus: %w", err)
	}
	return val, nil
}

// GetOrderJudgeStatus queries & returns the judge status of the order.
func (v *VEscrowCtrt) GetOrderJudgeStatus(orderId string) (bool, error) {
	dbKey, err := NewDBKeyVEscrowOrderJudgeStatus(orderId)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeStatus: %w", err)
	}
	resp, err := v.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeStatus: %w", err)
	}

	val, err := ctrtDataRespToBool(resp)
	if err != nil {
		return false, fmt.Errorf("GetOrderJudgeStatus: %w", err)
	}
	return val, nil
}

// GetOrderRecipientLockedAmount queries & returns the amount from the recipient
// that is locked(deposited) in the order.
func (v *VEscrowCtrt) GetOrderRecipientLockedAmount(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderRecipientLockedAmount(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientLockedAmount: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientLockedAmount: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientLockedAmount: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderRecipientlockedAmount: %w", err)
	}
	return tok, nil
}

// GetOrderJudgeLockedAmount queries & returns the amount from the judge
// that is locked(deposited) in the order.
func (v *VEscrowCtrt) GetOrderJudgeLockedAmount(orderId string) (*Token, error) {
	dbKey, err := NewDBKeyVEscrowOrderJudgeLockedAmount(orderId)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeLockedAmount: %w", err)
	}
	resp, err := v.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeLockedAmount: %w", err)
	}
	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgeLockedAmount: %w", err)
	}
	tok, err := ctrtDataRespToToken(resp, unit)
	if err != nil {
		return nil, fmt.Errorf("GetOrderJudgelockedAmount: %w", err)
	}
	return tok, nil
}

// TokCtrt queries and returns instance of token contract of V Escrow Contract's token.
func (v *VEscrowCtrt) TokCtrt() (BaseTokCtrt, error) {
	if v.tokCtrt == nil {
		tokId, err := v.TokId()
		if err != nil {
			return nil, err
		}
		instance, err := GetCtrtFromTokId(tokId, v.Chain)
		if err != nil {
			return nil, err
		}
		v.tokCtrt = instance
	}
	return v.tokCtrt, nil
}

// Supersede transfers the judge right of the contract to another account.
func (v *VEscrowCtrt) Supersede(by *Account, newJudge string, attachment string) (*BroadcastExecuteTxResp, error) {
	newJudgeAddr, err := NewAddrFromB58Str(newJudge)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_SUPERSEDE,
		DataStack{
			NewDeAddr(newJudgeAddr),
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

// Create creates an escrow order.
// NOTE that the transaction id of this action is the order ID.
func (v *VEscrowCtrt) Create(
	by *Account,
	recipient string,
	amount float64,
	rcpt_deposit_amount float64,
	judge_deposit_amount float64,
	order_fee float64,
	refund_amount float64,
	expire_at int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	rcptAddr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	deVEscrowAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	deRcptDepAmount, err := NewDeAmountForTokAmount(rcpt_deposit_amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	deJudgeDepAmount, err := NewDeAmountForTokAmount(judge_deposit_amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	deFeeAmount, err := NewDeAmountForTokAmount(order_fee, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	deRefundAmount, err := NewDeAmountForTokAmount(refund_amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_CREATE,
		DataStack{
			NewDeAddr(rcptAddr),
			deVEscrowAmount,
			deRcptDepAmount,
			deJudgeDepAmount,
			deFeeAmount,
			deRefundAmount,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(expire_at)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	return resp, nil
}

// RecipientDeposit deposits tokens the recipient deposited into the contract into the order.
func (v *VEscrowCtrt) RecipientDeposit(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("RecipientDeposit: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_DEPOSIT,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RecipientDeposit: %w", err)
	}
	return resp, err
}

// JudgeDeposit deposits tokens the judge deposited into the contract into the order.
func (v *VEscrowCtrt) JudgeDeposit(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("JudgeDeposit: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_JUDGE_DEPOSIT,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("JudgeDeposit: %w", err)
	}
	return resp, err
}

// PayerCancel cancels the order by the payer.
func (v *VEscrowCtrt) PayerCancel(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("PayerCancel: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_PAYER_CANCEL,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("PayerCancel: %w", err)
	}
	return resp, err
}

// RecipientCancel cancels the order by the recipient.
func (v *VEscrowCtrt) RecipientCancel(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("RecipientCancel: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_CANCEL,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RecipientCancel: %w", err)
	}
	return resp, err
}

// JudgeCancel cancels the order by the judge.
func (v *VEscrowCtrt) JudgeCancel(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("JudgeCancel: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_JUDGE_CANCEL,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("JudgeCancel: %w", err)
	}
	return resp, err
}

// SubmitWork submits the work by the recipient.
func (v *VEscrowCtrt) SubmitWork(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("SubmitWork: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_SUBMIT_WORK,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SubmitWork: %w", err)
	}
	return resp, err
}

// ApproveWork approves the work and agrees the amounts are paid by the payer.
func (v *VEscrowCtrt) ApproveWork(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("ApproveWork: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_APPROVE_WORK,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("ApproveWork: %w", err)
	}
	return resp, err
}

// ApplyToJudge applies for the help from judge by the payer.
func (v *VEscrowCtrt) ApplyToJudge(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("ApplyToJudge: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_APPLY_TO_JUDGE,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("ApplyToJudge: %w", err)
	}
	return resp, err
}

// DoJudge judges the work and decides on how much the payer & recipient will receive.
func (v *VEscrowCtrt) DoJudge(
	by *Account,
	orderId string,
	payerAmount float64,
	recipientAmount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("DoJudge: %w", err)
	}

	unit, err := v.Unit()
	if err != nil {
		return nil, fmt.Errorf("DoJudge: %w", err)
	}
	deAmountPayer, err := NewDeAmountForTokAmount(payerAmount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("DoJudge: %w", err)
	}
	deAmountRecipient, err := NewDeAmountForTokAmount(recipientAmount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("DoJudge: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_JUDGE,
		DataStack{
			NewDeBytes(b),
			deAmountPayer,
			deAmountRecipient,
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("DoJudge: %w", err)
	}
	return resp, err
}

// SubmitPenalty submits penalty by the payer for the case where the recipient does not submit
// work before the expiration time. The payer will obtain the recipient deposit amount and the payer amount(fee deducted).
// The judge will still get the fee.
func (v *VEscrowCtrt) SubmitPenalty(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("SubmitPenalty: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_SUBMIT_PENALTY,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("SubmitPenalty: %w", err)
	}
	return resp, err
}

// PayerRefund makes the refund action by the payer when the judge does not judge the work in time
// after the apply_to_judge function is invoked.
// The judge loses his deposit amount and the payer receives the refund amount.
// The recipient receives the rest.
func (v *VEscrowCtrt) PayerRefund(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("PayerRefund: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_PAYER_REFUND,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("PayerRefund: %w", err)
	}
	return resp, err
}

// RecipientRefund makes the refund action by the recipient when the judge does not judge the work in time
// after the apply_to_judge function is invoked.
// The judge loses his deposit amount and the payer receives the refund amount.
// The recipient receives the rest.
func (v *VEscrowCtrt) RecipientRefund(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("RecipientRefund: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_RECIPIENT_REFUND,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RecipientRefund: %w", err)
	}
	return resp, err
}

// Collect collects the order amount & recipient deposited amount by the recipient when the work is submitted
// while the payer doesn't either approve or apply to judge in his action duration.
// The judge will get judge deposited amount & fee.
func (v *VEscrowCtrt) Collect(
	by *Account,
	orderId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(orderId)
	if err != nil {
		return nil, fmt.Errorf("Collect: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		v.CtrtId,
		FUNC_IDX_V_ESCROW_CTRT_COLLECT,
		DataStack{
			NewDeBytes(b),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Collect: %w", err)
	}
	return resp, err
}
