package vsys

import (
	"fmt"
)

type NFTCtrt struct {
	*Ctrt
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

	return by.ExecuteCtrt(txReq)
}

func (n *NFTCtrt) Supersede(by *Account, newIssuer string, attachment string) (*BroadcastExecuteTxResp, error) {
	addr, err := NewAddrFromB58Str(newIssuer)
	if err != nil {
		return nil, err
	}

	txReq := NewExecCtrtFuncTxReq(
		n.CtrtId,
		FUNC_IDX_NFT_SUPERSEDE,
		DataStack{NewDeAddr(addr)},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	return by.ExecuteCtrt(txReq)
}
