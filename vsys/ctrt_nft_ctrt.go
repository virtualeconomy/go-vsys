package vsys

import (
	"fmt"
)

type NFTCtrt struct {
	*Ctrt
}

func (n NFTCtrt) Unit() Unit {
	// NFT contract have unit of 1
	return 1
}

func NewNFTCtrt(ctrtId string, chain *Chain) (*NFTCtrt, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewNFTCtrt: %w", err)
	}

	return &NFTCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

func RegisterNFTCtrt(by *Account, ctrtDescription string) (*NFTCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForNFTCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}

	txReq := NewRegCtrtTxReq(
		DataStack{},
		ctrtMeta,
		NewVSYSTimestampForNow(),
		Str(ctrtDescription),
		FEE_REG_CTRT,
	)

	resp, err := by.RegisterCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}

	cid, err := NewCtrtIdFromB58Str(resp.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}

	return &NFTCtrt{
		&Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

func NewDBKeyNFTCtrtIssuer() Bytes {
	return STATE_VAR_NFT_ISSUER.Serialize()
}

func (n *NFTCtrt) Issuer() (*Addr, error) {
	resp, err := n.QueryDBKey(
		NewDBKeyNFTCtrtIssuer(),
	)
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}
	addr, err := NewAddrFromB58Str(resp.Val.Str())
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}
	return addr, nil
}

func (n *NFTCtrt) Issue(by *Account, tokenDescription, attachment string) (*BroadcastExecuteTxResp, error) {
	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_ISSUE,
		DataStack{NewDeStrFromString(tokenDescription)},
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

func (n *NFTCtrt) Supersede(by *Account, newIssuer, attachment string) (*BroadcastExecuteTxResp, error) {
	addr, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_SUPERSEDE,
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

func (n *NFTCtrt) Send(by *Account, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	rcpt_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}
	// TODO: move to MustOn() bool function
	if rcpt_addr.ChainID() != by.Chain.ChainID {
		return nil, fmt.Errorf("Send: Adress must be on same chain")
	}

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_SEND,
		DataStack{
			NewDeAddr(rcpt_addr),
			NewDeInt32(uint32(tok_idx)),
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

func (n *NFTCtrt) Transfer(by *Account, sender, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	rcpt_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	sender_addr, err := NewAddrFromB58Str(sender)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	// TODO: move to MustOn() bool function
	if rcpt_addr.ChainID() != by.Chain.ChainID {
		return nil, fmt.Errorf("Transfer: Adress must be on same chain")
	}
	if sender_addr.ChainID() != by.Chain.ChainID {
		return nil, fmt.Errorf("Transfer: Adress must be on same chain")
	}

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_TRANSFER,
		DataStack{
			NewDeAddr(sender_addr),
			NewDeAddr(rcpt_addr),
			NewDeInt32(uint32(tok_idx)),
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

func (n *NFTCtrt) Deposit(
	by *Account,
	ctrtId string,
	tokIdx int,
	attachment string,
) (*BroadcastExecuteTxResp, error) {

	// TODO: dataentry_ctrtAccount
	ctrtID, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	ctrtAccount := NewDeCtrtAddrFromCtrtId(ctrtID)

	// model.TokenIdx not needed?

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_DEPOSIT,
		DataStack{NewDeAddr(by.Addr), ctrtAccount, NewDeInt32(uint32(tokIdx))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	return by.ExecuteCtrt(txReq)
}

func (n *NFTCtrt) Withdraw(by *Account, ctrtId string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtID, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_TRANSFER,
		DataStack{
			NewDeCtrtAddrFromCtrtId(ctrtID),
			NewDeAddr(by.Addr),
			NewDeInt32(uint32(tok_idx)),
		},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}
	return resp, nil
}
