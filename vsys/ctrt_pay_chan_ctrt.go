package vsys

import (
	"fmt"
)

// PayChanCtrt is the struct for VSYS Payment Channel Contract.
type PayChanCtrt struct {
	*Ctrt
	tokId   *TokenId
	tokCtrt BaseTokCtrt
}

// NewPayChanCtrt creates instance of PayChanCtrt from given contract id.
func NewPayChanCtrt(ctrtId string, chain *Chain) (*PayChanCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewPayChanCtrt: %w", err)
	}

	return &PayChanCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
		tokId:   nil,
		tokCtrt: nil,
	}, nil
}

// RegisterPayChanCtrt registers a Payment Channel Contract.
func RegisterPayChanCtrt(by *Account, tokenId, ctrtDescription string) (*PayChanCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForPayChanCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterPayChanCtrt: %w", err)
	}

	tokId, err := NewTokenIdFromB58Str(tokenId)
	if err != nil {
		return nil, fmt.Errorf("RegisterPayChanCtrt: %w", err)
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
		return nil, fmt.Errorf("RegisterPayChanCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}

	return &PayChanCtrt{
		Ctrt: &Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
		tokId:   nil,
		tokCtrt: nil,
	}, nil
}

// CreateAndLoad creates the payment channel and loads an amount into it.
// (This function's transaction id becomes the channel ID)
func (p *PayChanCtrt) CreateAndLoad(
	by *Account,
	recipient string,
	amount float64,
	expireAt int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	rcptMd, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("CreateAndLoad: %w", err)
	}
	rcptMd.MustOn(p.Chain)

	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("CreateAndLoad: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("CreateAndLoad: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_CREATE_AND_LOAD,
		DataStack{
			NewDeAddr(rcptMd),
			deAmount,
			NewDeTimestamp(NewVSYSTimestampFromUnixTs(expireAt)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("CreateAndLoad: %w", err)
	}
	return resp, nil
}

// ExtendExpTime extends the expiration time of the channel to the new input timestamp.
func (p *PayChanCtrt) ExtendExpTime(
	by *Account,
	chanId string,
	expireAt int64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("ExtendExpTime: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_EXTEND_EXPIRATION_TIME,
		DataStack{NewDeBytes(b), NewDeTimestamp(NewVSYSTimestampFromUnixTs(expireAt))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("ExtendExpTime: %w", err)
	}
	return resp, nil
}

// Load loads more tokens into the channel.
func (p *PayChanCtrt) Load(
	by *Account,
	chanId string,
	amount float64,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("Load: %w", err)
	}

	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("Load: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("Load: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_LOAD,
		DataStack{NewDeBytes(b), deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Load: %w", err)
	}
	return resp, nil
}

// Abort aborts the channel, triggering a 2-day grace period where the recipient can still
// collect payments. After 2 days, the payer can unload all the remaining funds that was locked
// in the channel.
func (p *PayChanCtrt) Abort(
	by *Account,
	chanId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("Abort: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_ABORT,
		DataStack{NewDeBytes(b)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Abort: %w", err)
	}
	return resp, nil
}

// Unload unloads all the funcs locked in the channel (only works if the channel has expired).
func (p *PayChanCtrt) Unload(
	by *Account,
	chanId string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("Unload: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_UNLOAD,
		DataStack{NewDeBytes(b)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Unload: %w", err)
	}
	return resp, nil
}

// CollectPayment collects the payment from the channel.
func (p *PayChanCtrt) CollectPayment(
	by *Account,
	chanId string,
	amount float64,
	signature, attachment string,
) (*BroadcastExecuteTxResp, error) {
	ok, err := p.VerifySig(chanId, amount, signature)
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}
	if !ok {
		return nil, fmt.Errorf("CollectPayment: Invalid Payment Channel Contract payment signature")
	}

	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}
	sb, err := NewBytesFromB58Str(signature)
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}

	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		p.CtrtId,
		FUNC_IDX_PAY_CHAN_COLLECT_PAYMENT,
		DataStack{NewDeBytes(b), deAmount, NewDeBytes(sb)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("CollectPayment: %w", err)
	}
	return resp, nil
}

func (p *PayChanCtrt) OffchainPay(key *PriKey, chanId string, amount float64) (string, error) {
	msg, err := p.getPayMsg(chanId, amount)
	if err != nil {
		return "", fmt.Errorf("OffchainPay: %w", err)
	}
	sig_bytes, err := Sign(key.Bytes, msg)
	if err != nil {
		return "", fmt.Errorf("OffchainPay: %w", err)
	}
	return B58Encode(sig_bytes), nil
}

func (p *PayChanCtrt) VerifySig(
	chanId string,
	amount float64,
	signature string,
) (bool, error) {
	msg, err := p.getPayMsg(chanId, amount)
	if err != nil {
		return false, fmt.Errorf("VerifySig: %w", err)
	}
	pubKey, err := p.GetChanCreatorPubKey(chanId)
	if err != nil {
		return false, fmt.Errorf("VerifySig: %w", err)
	}
	sig_bytes, err := B58Decode(signature)
	if err != nil {
		return false, fmt.Errorf("VerifySig: %w", err)
	}
	return Verify(pubKey.Bytes, msg, sig_bytes), nil
}

func (p *PayChanCtrt) getPayMsg(chanId string, amount float64) (Bytes, error) {
	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("getPayMsg: %w", err)
	}
	rawAmount, err := NewTokenForAmount(amount, uint64(unit))
	if err != nil {
		return nil, fmt.Errorf("getPayMsg: %w", err)
	}
	chainIdBytes, err := B58Decode(chanId)
	if err != nil {
		return nil, fmt.Errorf("getPayMsg: %w", err)
	}
	msg := append(PackUInt16(uint16(len(chainIdBytes))), chainIdBytes...)
	msg = append(msg, PackUInt64(uint64(rawAmount.Data))...)
	return msg, nil
}

// NewDBKeyPayChanMaker returns DB key to query Maker of Payment Channel Contract.
func NewDBKeyPayChanMaker() Bytes {
	return STATE_VAR_PAY_CHAN_MAKER.Serialize()
}

// Maker queries and returns Addr of the Maker of contract.
func (p *PayChanCtrt) Maker() (*Addr, error) {
	resp, err := p.QueryDBKey(
		NewDBKeyPayChanMaker(),
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

// NewDBKeyPayChanTokId returns DB key to query TokenId of contract's token.
func NewDBKeyPayChanTokId() Bytes {
	return STATE_VAR_PAY_CHAN_TOKEN_ID.Serialize()
}

// TokId queries and returns TokenId of the contract's token.
func (p *PayChanCtrt) TokId() (*TokenId, error) {
	if p.tokId == nil {
		resp, err := p.QueryDBKey(
			NewDBKeyPayChanTokId(),
		)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		tokId, err := NewTokenIdFromB58Str(resp.Val.(string))
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
		p.tokId = tokId
	}
	return p.tokId, nil
}

// TokCtrt queries and returns instance of token contract of Atomic Swap Contract's token.
func (p *PayChanCtrt) TokCtrt() (BaseTokCtrt, error) {
	if p.tokCtrt == nil {
		tokId, err := p.TokId()
		if err != nil {
			return nil, err
		}
		instance, err := GetCtrtFromTokId(tokId, p.Chain)
		if err != nil {
			return nil, err
		}
		p.tokCtrt = instance
	}
	return p.tokCtrt, nil
}

// Unit queries and returns Unit of the token of contract.
func (p *PayChanCtrt) Unit() (Unit, error) {
	if p.tokCtrt == nil {
		_, err := p.TokCtrt() // TokCtrt sets p.TokCtrt
		if err != nil {
			return 0, fmt.Errorf("Unit: %w", err)
		}
	}
	return p.tokCtrt.Unit()
}

// NewDBKeyPayChanGetCtrtBal returns DB key for querying the contract balance for given address.
func NewDBKeyPayChanGetCtrtBal(addr *Addr) Bytes {
	return NewStateMap(
		STATE_MAP_IDX_PAY_CHAN_CONTRACT_BALANCE,
		NewDeAddr(addr)).Serialize()
}

// GetCtrtBal queries and returns the balance of the token deposited into the contract.
func (p *PayChanCtrt) GetCtrtBal(addr string) (*Token, error) {
	query_addr, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	data, err := p.QueryDBKey(NewDBKeyPayChanGetCtrtBal(query_addr))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtBal: %w", err)
	}

	unit, err := p.Unit()
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

func NewDBKeyPayChanCtrtChanCreator(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChanCreator: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_CREATOR, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanCreator(chanId string) (*Addr, error) {
	dbKey, err := NewDBKeyPayChanCtrtChanCreator(chanId)
	if err != nil {
		return nil, fmt.Errorf("GetChanCreator: %w", err)
	}

	resp, err := p.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetChanCreator: %w", err)
	}

	switch addrB58 := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("GetChanCreator: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("GetChanCreator: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

func NewDBKeyPayChanCtrtChanCreatorPubKey(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChanCreatorPubKey: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_CREATOR_PUBLIC_KEY, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanCreatorPubKey(chanId string) (*PubKey, error) {
	dbKey, err := NewDBKeyPayChanCtrtChanCreatorPubKey(chanId)
	if err != nil {
		return nil, fmt.Errorf("GetChanCreatorPubKey: %w", err)
	}

	resp, err := p.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetChanCreatorPubKey: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		pubKey, err := NewPubKeyFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("GetChanCreatorPubKey: %w", err)
		}
		return pubKey, nil
	default:
		return nil, fmt.Errorf("GetChanCreatorPubKey: CtrtDataResp.Val is %T but string was expected", val)
	}
}

func NewDBKeyPayChanCtrtChanRecipient(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChanRecipient: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_RECIPIENT, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanRecipient(chanId string) (*Addr, error) {
	dbKey, err := NewDBKeyPayChanCtrtChanRecipient(chanId)
	if err != nil {
		return nil, fmt.Errorf("GetChanRecipient: %w", err)
	}

	resp, err := p.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetChanRecipient: %w", err)
	}

	switch addrB58 := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(addrB58)
		if err != nil {
			return nil, fmt.Errorf("GetChanRecipient: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("GetChanRecipient: CtrtDataResp.Val is %T but string was expected", addrB58)
	}
}

func NewDBKeyPayChanCtrtChannelAccumLoad(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChannelAccumLoad: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_ACCUMULATED_LOAD, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanAccumLoad(chanId string) (*Token, error) {
	dbKey, err := NewDBKeyPayChanCtrtChannelAccumLoad(chanId)
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumLoad: %w", err)
	}
	data, err := p.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumLoad: %w", err)
	}

	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumLoad: %w", err)
	}

	switch amount := data.Val.(type) {
	case float64:
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetChanAccumLoad: CtrtDataResp.Val is %T but float64 was expected", amount)
	}
}

func NewDBKeyPayChanCtrtChannelAccumPay(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChannelAccumPay: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_ACCUMULATED_PAYMENT, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanAccumPay(chanId string) (*Token, error) {
	dbKey, err := NewDBKeyPayChanCtrtChannelAccumPay(chanId)
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumPay: %w", err)
	}
	data, err := p.QueryDBKey(dbKey)
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumPay: %w", err)
	}

	unit, err := p.Unit()
	if err != nil {
		return nil, fmt.Errorf("GetChanAccumPay: %w", err)
	}

	switch amount := data.Val.(type) {
	case float64:
		return NewToken(Amount(amount), unit), nil
	default:
		return nil, fmt.Errorf("GetChanAccumPay: CtrtDataResp.Val is %T but float64 was expected", amount)
	}
}

func NewDBKeyPayChanCtrtChanExpTime(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChanExpTime: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_EXPIRATION_TIME, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanExpTime(chanId string) (VSYSTimestamp, error) {
	dbKey, err := NewDBKeyPayChanCtrtChanExpTime(chanId)
	if err != nil {
		return 0, fmt.Errorf("GetChanExpTime: %w", err)
	}

	data, err := p.QueryDBKey(dbKey)
	if err != nil {
		return 0, fmt.Errorf("GetChanExpTime: %w", err)
	}

	switch val := data.Val.(type) {
	case float64:
		return VSYSTimestamp(val), nil
	default:
		return 0, fmt.Errorf("GetChanExpTime: CtrtDataResp.Val is %T but float64 was expected", val)
	}
}

func NewDBKeyPayChanCtrtChanStatus(chanId string) (Bytes, error) {
	b, err := NewBytesFromB58Str(chanId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyPayChanCtrtChanStatus: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_PAY_CHAN_CHANNEL_STATUS, NewDeBytes(b)).Serialize(), nil
}

func (p *PayChanCtrt) GetChanStatus(chanId string) (bool, error) {
	dbKey, err := NewDBKeyPayChanCtrtChanStatus(chanId)
	if err != nil {
		return false, fmt.Errorf("GetChanStatus: %w", err)
	}

	resp, err := p.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("GetChanStatus: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	default:
		return false, fmt.Errorf("GetChanStatus: CtrtDataResp.Val is %T but string was expected", val)
	}
}
