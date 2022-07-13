package vsys

import "fmt"

type AtomicSwapCtrt struct {
	*Ctrt
}

func RegisterAtomicSwapCtrt(by *Account, tokenId string, ctrtDescription string) (*AtomicSwapCtrt, error) {
	ctrtMeta, err := newCtrtMetaForAtomicSwapCtrt()
	if err != nil {
		return nil, fmt.Errorf("RegisterAtomicSwapCtrt: %w", err)
	}

	tokId, err := NewTokenIdFromB58Str(tokenId)

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

	//fmt.Println(resp)

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
