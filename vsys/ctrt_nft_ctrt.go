package vsys

import (
	"fmt"
)

type INFTCtrt interface {
	QueryDBKey(Bytes) (*CtrtDataResp, error)
	ctrtId() *CtrtId
	chain() *Chain
}

// register is internal implementation for Register for NFT contracts.
func register(ctrtMeta *CtrtMeta, by *Account, ctrtDescription string) (*CtrtId, error) {
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

	return cid, nil
}

// maker is internal implementation for Maker.
func maker(n INFTCtrt, dbKey Bytes) (*Addr, error) {
	resp, err := n.QueryDBKey(
		dbKey,
	)
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

// issuer is internal implementation for Issuer.
func issuer(n INFTCtrt, dbKey Bytes) (*Addr, error) {
	resp, err := n.QueryDBKey(
		dbKey,
	)
	if err != nil {
		return nil, fmt.Errorf("Issuer: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("Issuer: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Issuer: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// lastIndex is internal implementation for LastIndex.
func lastIndex(n INFTCtrt) (uint32, error) {
	res, err := n.chain().NodeAPI.GetLastIndex(string(n.ctrtId().B58Str()))
	if err != nil {
		return 0, fmt.Errorf("LastIndex: %w", err)
	}
	return res.LastTokenIdx, nil
}

// issue is internal implementation for Issue.
func issue(n INFTCtrt, funcIdx FuncIdx, by *Account, tokenDescription, attachment string) (*BroadcastExecuteTxResp, error) {
	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		funcIdx,
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

// send is internal implementation of Send functions.
func send(n INFTCtrt, funcIdx FuncIdx, by *Account, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	rcpt_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	err = rcpt_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		funcIdx,
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

// transfer is internal implementation of Transfer functions.
func transfer(n INFTCtrt, funcIdx FuncIdx, by *Account, sender, recipient string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	rcpt_addr, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	sender_addr, err := NewAddrFromB58Str(sender)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	err = sender_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	err = rcpt_addr.MustOn(by.Chain)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		funcIdx,
		DataStack{
			NewDeAddr(sender_addr),
			NewDeAddr(rcpt_addr),
			NewDeInt32(uint32(tokIdx)),
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

// deposit is internal implementation of Deposit functions.
func deposit(
	n INFTCtrt,
	funcIdx FuncIdx,
	by *Account,
	ctrtId string,
	tokIdx int,
	attachment string,
) (*BroadcastExecuteTxResp, error) {

	ctrtID, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	ctrtAccount := NewDeCtrtAddrFromCtrtId(ctrtID)

	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		funcIdx,
		DataStack{NewDeAddr(by.Addr), ctrtAccount, NewDeInt32(uint32(tokIdx))},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}
	return resp, nil
}

// withdraw is internal implementation of Withdraw functions.
func withdraw(n INFTCtrt, funcIdx FuncIdx, by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtID, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		funcIdx,
		DataStack{
			NewDeCtrtAddrFromCtrtId(ctrtID),
			NewDeAddr(by.Addr),
			NewDeInt32(uint32(tokIdx)),
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

type NFTCtrt struct {
	*Ctrt
}

func (n *NFTCtrt) ctrtId() *CtrtId {
	return n.CtrtId
}

func (n *NFTCtrt) chain() *Chain {
	return n.Chain
}

// NewNFTCtrt creates an instance of NewNFTCtrt from given contract id.
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

// RegisterNFTCtrt registers a token contract.
func RegisterNFTCtrt(by *Account, ctrtDescription string) (*NFTCtrt, error) {
	ctrtMeta, err := NewCtrtMetaForNFTCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrt: %w", err)
	}
	cid, err := register(ctrtMeta, by, ctrtDescription)
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

// Unit returns unit of NFT Contract(unit=1 for nft ctrts).
func (n *NFTCtrt) Unit() (Unit, error) {
	// NFT contract have unit of 1
	return 1, nil
}

func NewDBKeyNFTCtrtMaker() Bytes {
	return STATE_VAR_NFT_MAKER.Serialize()
}

func NewDBKeyNFTCtrtIssuer() Bytes {
	return STATE_VAR_NFT_ISSUER.Serialize()
}

// Maker queries & returns the maker of the contract.
func (n *NFTCtrt) Maker() (*Addr, error) {
	return maker(n, NewDBKeyNFTCtrtMaker())
}

// Issuer queries & returns the issuer of the contract.
func (n *NFTCtrt) Issuer() (*Addr, error) {
	return issuer(n, NewDBKeyNFTCtrtIssuer())
}

// LastIndex returns the last index of the NFT contract.
func (n *NFTCtrt) LastIndex() (uint32, error) {
	return lastIndex(n)
}

// Issue issues a token of the NFT contract.
func (n *NFTCtrt) Issue(by *Account, tokenDescription, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(n, FUNC_IDX_NFT_ISSUE, by, tokenDescription, attachment)
}

// Send sends the NFT token from the action taker to the recipient.
func (n *NFTCtrt) Send(by *Account, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(n, FUNC_IDX_NFT_SEND, by, recipient, tok_idx, attachment)
}

// Transfer transfers the NFT token from the sender to the recipient.
func (n *NFTCtrt) Transfer(by *Account, sender, recipient string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(n, FUNC_IDX_NFT_TRANSFER, by, sender, recipient, tokIdx, attachment)
}

// Deposit deposits the NFT token from the action taker to another contract.
func (n *NFTCtrt) Deposit(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(n, FUNC_IDX_NFT_DEPOSIT, by, ctrtId, tokIdx, attachment)
}

// Withdraw withdraws the token from another contract to the action taker.
func (n *NFTCtrt) Withdraw(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(n, FUNC_IDX_NFT_WITHDRAW, by, ctrtId, tokIdx, attachment)
}

// Supersede transfers the issuer role of the contract to a new account.
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

// General functions for V2 NFT contracts

// NewDBKeyNFTCtrtV2ForRegulator returns the DBKey for querying the regulator of contract.
func NewDBKeyNFTCtrtV2ForRegulator() Bytes {
	return STATE_VAR_NFTV2_REGULATOR.Serialize()
}

// NewDBKeyNFTCtrtV2ForUserInList returns the DBKey for querying the status of if the given user address is in the list.
func NewDBKeyNFTCtrtV2ForUserInList(addr string) (Bytes, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyNFTCtrtV2ForUserInList: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_NFTV2_IS_IN_LIST, NewDeAddr(addrMd)).Serialize(), nil
}

// NewDBKeyNFTCtrtV2ForCtrtInList returns the DBKey for querying the status of if the given contract address is in the list.
func NewDBKeyNFTCtrtV2ForCtrtInList(ctrtId string) (Bytes, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewDBKeyNFTCtrtV2ForCtrtInList: %w", err)
	}
	return NewStateMap(STATE_MAP_IDX_NFTV2_IS_IN_LIST, NewDeCtrtAddrFromCtrtId(ctrtIdMd)).Serialize(), nil
}

// isInList queries & returns the status of whether the address is in the list for the given db_key.
func isInList(t INFTCtrt, dbKey Bytes) (bool, error) {
	resp, err := t.QueryDBKey(dbKey)
	if err != nil {
		return false, fmt.Errorf("isInList: %w", err)
	}
	switch val := resp.Val.(type) {
	case string:
		return val == "true", nil
	case float64:
		if val == 0 {
			return false, nil
		}
		return false, fmt.Errorf("isInList: 'dbName:contractNumInfo' but value != 0")
	default:
		return false, fmt.Errorf("isInList: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// regulator is internal implementation for Regulator.
func regulator(t INFTCtrt) (*Addr, error) {
	resp, err := t.QueryDBKey(NewDBKeyNFTCtrtV2ForRegulator())
	if err != nil {
		return nil, fmt.Errorf("Regulator: %w", err)
	}

	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("Regulator: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("Regulator: CtrtDataResp.Val is %T but string was expected", val)
	}
}

// updateList updates the presence of the address within the given data entry in the list.
// It's the helper method for UpdateList*.
func updateList(
	n INFTCtrt,
	by *Account,
	addrDe DataEntry,
	val bool,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	txReq := NewExecCtrtFuncTxReq(
		n.ctrtId(),
		FUNC_IDX_NFTV2_UPDATE_LIST,
		DataStack{addrDe, NewDeBool(val)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)

	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("UpdateList: %w", err)
	}
	return resp, nil
}

// supersedeCtrtWithList is internal implementation of Supersede for contracts with lists.
func supersedeCtrtWithList(
	t INFTCtrt,
	by *Account,
	newIssuer string,
	newRegulator string,
	attachment string,
) (*BroadcastExecuteTxResp, error) {
	newIssuerMd, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	newRegulatorMd, err := NewAddrFromB58Str(newRegulator)
	if err != nil {
		return nil, fmt.Errorf("Supersede: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		t.ctrtId(),
		FUNC_IDX_NFTV2_SUPERSEDE,
		DataStack{
			NewDeAddr(newIssuerMd),
			NewDeAddr(newRegulatorMd),
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

type NFTCtrtV2Whitelist struct {
	*Ctrt
}

func (n *NFTCtrtV2Whitelist) ctrtId() *CtrtId {
	return n.CtrtId
}

func (n *NFTCtrtV2Whitelist) chain() *Chain {
	return n.Chain
}

// Unit returns unit of NFT Contract (unit=1 for nft ctrts).
func (n *NFTCtrtV2Whitelist) Unit() (Unit, error) {
	return 1, nil
}

// NewNFTCtrtV2Whitelist creates an instance of NewNFTCtrtV2Whitelist from given contract id.
func NewNFTCtrtV2Whitelist(ctrtId string, chain *Chain) (*NFTCtrtV2Whitelist, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewNFTCtrt: %w", err)
	}

	return &NFTCtrtV2Whitelist{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

// RegisterNFTCtrtV2Whitelist registers nft contract.
func RegisterNFTCtrtV2Whitelist(by *Account, ctrtDescription string) (*NFTCtrtV2Whitelist, error) {
	ctrtMeta, err := NewCtrtMetaForNFTCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrtV2Whitelist: %w", err)
	}
	cid, err := register(ctrtMeta, by, ctrtDescription)
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrtV2Whitelist: %w", err)
	}
	return &NFTCtrtV2Whitelist{
		&Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

// Maker queries & returns the maker of the contract.
func (n *NFTCtrtV2Whitelist) Maker() (*Addr, error) {
	return maker(n, NewDBKeyNFTCtrtMaker())
}

// Issuer queries & returns the issuer of the contract.
func (n *NFTCtrtV2Whitelist) Issuer() (*Addr, error) {
	return issuer(n, NewDBKeyNFTCtrtIssuer())
}

// LastIndex returns the last index of the NFT contract.
func (n *NFTCtrtV2Whitelist) LastIndex() (uint32, error) {
	return lastIndex(n)
}

// Issue issues a token of the NFT contract.
func (n *NFTCtrtV2Whitelist) Issue(by *Account, tokenDescription, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(n, FUNC_IDX_NFTV2_ISSUE, by, tokenDescription, attachment)
}

// Send sends the NFT token from the action taker to the recipient.
func (n *NFTCtrtV2Whitelist) Send(by *Account, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(n, FUNC_IDX_NFTV2_SEND, by, recipient, tok_idx, attachment)
}

// Transfer transfers the NFT token from the sender to the recipient.
func (n *NFTCtrtV2Whitelist) Transfer(by *Account, sender, recipient string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(n, FUNC_IDX_NFTV2_TRANSFER, by, sender, recipient, tokIdx, attachment)
}

// Deposit deposits the NFT token from the action taker to another contract.
func (n *NFTCtrtV2Whitelist) Deposit(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(n, FUNC_IDX_NFTV2_DEPOSIT, by, ctrtId, tokIdx, attachment)
}

// Withdraw withdraws the token from another contract to the action taker.
func (n *NFTCtrtV2Whitelist) Withdraw(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(n, FUNC_IDX_NFTV2_WITHDRAW, by, ctrtId, tokIdx, attachment)
}

// Regulator queries & returns the regulator of the contract.
func (n *NFTCtrtV2Whitelist) Regulator() (*Addr, error) {
	return regulator(n)
}

// IsUserInList queries & returns the status of whether the user address in the whitelist.
func (n *NFTCtrtV2Whitelist) IsUserInList(addr string) (bool, error) {
	dbKey, err := NewDBKeyNFTCtrtV2ForUserInList(addr)
	if err != nil {
		return false, fmt.Errorf("IsUserInList: %w", err)
	}
	return isInList(n, dbKey)
}

// IsCtrtInList queries & returns the status of whether the contract address in the whitelist.
func (n *NFTCtrtV2Whitelist) IsCtrtInList(ctrtId string) (bool, error) {
	dbKey, err := NewDBKeyNFTCtrtV2ForCtrtInList(ctrtId)
	if err != nil {
		return false, fmt.Errorf("IsCtrtInList: %w", err)
	}
	return isInList(n, dbKey)
}

// UpdateListUser updates the presence of the user address in the list.
func (n *NFTCtrtV2Whitelist) UpdateListUser(by *Account, addr string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("UpdateListUser: %w", err)
	}
	return updateList(n, by, NewDeAddr(addrMd), val, attachment)
}

// UpdateListCtrt updates the presence of the contract address in the list.
func (n *NFTCtrtV2Whitelist) UpdateListCtrt(by *Account, ctrtId string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("UpdateListCtrt: %w", err)
	}
	return updateList(n, by, NewDeCtrtAddrFromCtrtId(ctrtIdMd), val, attachment)
}

// Supersede transfers the issuer role of the contract to a new account.
func (n *NFTCtrtV2Whitelist) Supersede(by *Account, newIssuer, newRegulator, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersedeCtrtWithList(n, by, newIssuer, newRegulator, attachment)
}

type NFTCtrtV2Blacklist struct {
	*Ctrt
}

func (n *NFTCtrtV2Blacklist) ctrtId() *CtrtId {
	return n.CtrtId
}

func (n *NFTCtrtV2Blacklist) chain() *Chain {
	return n.Chain
}

// Unit returns unit of NFT Contract (unit=1 for nft ctrts).
func (n *NFTCtrtV2Blacklist) Unit() (Unit, error) {
	return 1, nil
}

// NewNFTCtrtV2Blacklist creates an instance of NewNFTCtrtV2Blacklist from given contract id.
func NewNFTCtrtV2Blacklist(ctrtId string, chain *Chain) (*NFTCtrtV2Blacklist, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("NewNFTCtrt: %w", err)
	}

	return &NFTCtrtV2Blacklist{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  chain,
		},
	}, nil
}

// RegisterNFTCtrtV2Blacklist registers nft contract.
func RegisterNFTCtrtV2Blacklist(by *Account, ctrtDescription string) (*NFTCtrtV2Blacklist, error) {
	ctrtMeta, err := NewCtrtMetaForNFTCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrtV2Blacklist: %w", err)
	}
	cid, err := register(ctrtMeta, by, ctrtDescription)
	if err != nil {
		return nil, fmt.Errorf("RegisterNFTCtrtV2Blacklist: %w", err)
	}
	return &NFTCtrtV2Blacklist{
		&Ctrt{
			CtrtId: cid,
			Chain:  by.Chain,
		},
	}, nil
}

// Maker queries & returns the maker of the contract.
func (n *NFTCtrtV2Blacklist) Maker() (*Addr, error) {
	return maker(n, NewDBKeyNFTCtrtMaker())
}

// Issuer queries & returns the issuer of the contract.
func (n *NFTCtrtV2Blacklist) Issuer() (*Addr, error) {
	return issuer(n, NewDBKeyNFTCtrtIssuer())
}

// LastIndex returns the last index of the NFT contract.
func (n *NFTCtrtV2Blacklist) LastIndex() (uint32, error) {
	return lastIndex(n)
}

// Issue issues a token of the NFT contract.
func (n *NFTCtrtV2Blacklist) Issue(by *Account, tokenDescription, attachment string) (*BroadcastExecuteTxResp, error) {
	return issue(n, FUNC_IDX_NFTV2_ISSUE, by, tokenDescription, attachment)
}

// Send sends the NFT token from the action taker to the recipient.
func (n *NFTCtrtV2Blacklist) Send(by *Account, recipient string, tok_idx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return send(n, FUNC_IDX_NFTV2_SEND, by, recipient, tok_idx, attachment)
}

// Transfer transfers the NFT token from the sender to the recipient.
func (n *NFTCtrtV2Blacklist) Transfer(by *Account, sender, recipient string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return transfer(n, FUNC_IDX_NFTV2_TRANSFER, by, sender, recipient, tokIdx, attachment)
}

// Deposit deposits the NFT token from the action taker to another contract.
func (n *NFTCtrtV2Blacklist) Deposit(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return deposit(n, FUNC_IDX_NFTV2_DEPOSIT, by, ctrtId, tokIdx, attachment)
}

// Withdraw withdraws the token from another contract to the action taker.
func (n *NFTCtrtV2Blacklist) Withdraw(by *Account, ctrtId string, tokIdx int, attachment string) (*BroadcastExecuteTxResp, error) {
	return withdraw(n, FUNC_IDX_NFTV2_WITHDRAW, by, ctrtId, tokIdx, attachment)
}

// Regulator queries & returns the regulator of the contract.
func (n *NFTCtrtV2Blacklist) Regulator() (*Addr, error) {
	return regulator(n)
}

// IsUserInList queries & returns the status of whether the user address in the black list.
func (n *NFTCtrtV2Blacklist) IsUserInList(addr string) (bool, error) {
	dbKey, err := NewDBKeyNFTCtrtV2ForUserInList(addr)
	if err != nil {
		return false, fmt.Errorf("IsUserInList: %w", err)
	}
	return isInList(n, dbKey)
}

// IsCtrtInList queries & returns the status of whether the contract address in the blacklist.
func (n *NFTCtrtV2Blacklist) IsCtrtInList(ctrtId string) (bool, error) {
	dbKey, err := NewDBKeyNFTCtrtV2ForCtrtInList(ctrtId)
	if err != nil {
		return false, fmt.Errorf("IsCtrtInList: %w", err)
	}
	return isInList(n, dbKey)
}

// UpdateListUser updates the presence of the user address in the list.
func (n *NFTCtrtV2Blacklist) UpdateListUser(by *Account, addr string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	addrMd, err := NewAddrFromB58Str(addr)
	if err != nil {
		return nil, fmt.Errorf("UpdateListUser: %w", err)
	}
	return updateList(n, by, NewDeAddr(addrMd), val, attachment)
}

// UpdateListCtrt updates the presence of the contract address in the list.
func (n *NFTCtrtV2Blacklist) UpdateListCtrt(by *Account, ctrtId string, val bool, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrtIdMd, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("UpdateListCtrt: %w", err)
	}
	return updateList(n, by, NewDeCtrtAddrFromCtrtId(ctrtIdMd), val, attachment)
}

// Supersede transfers the issuer role of the contract to a new account.
func (n *NFTCtrtV2Blacklist) Supersede(by *Account, newIssuer, newRegulator, attachment string) (*BroadcastExecuteTxResp, error) {
	return supersedeCtrtWithList(n, by, newIssuer, newRegulator, attachment)
}
