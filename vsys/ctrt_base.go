package vsys

import "fmt"

type Ctrt struct {
	CtrtId *CtrtId
	Chain  *Chain
}

func (c *Ctrt) String() string {
	return fmt.Sprintf("%T(%+v)", c, *c)
}

func (c *Ctrt) QueryDBKey(dbKey Bytes) (*CtrtDataResp, error) {
	resp, err := c.Chain.NodeAPI.GetCtrtData(
		c.CtrtId.B58Str().Str(),
		dbKey.B58Str().Str(),
	)
	if err != nil {
		return nil, fmt.Errorf("QueryDBKey: %w", err)
	}
	return resp, nil
}

type BaseTokCtrt interface {
	Unit() uint64
}

func GetCtrtFromTokId(tokId *TokenId, chain *Chain) (BaseTokCtrt, error) {
	tokInfo, err := chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
	}
	ctrtInfo, err := chain.NodeAPI.GetCtrtInfo(tokInfo.CtrtId.Str())
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
	}
	// Switch statement here to choose constructor
	switch string(ctrtInfo.Type) {
	case "NonFungibleContract":
		n, err := NewNFTCtrt(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
		}
		return n, nil
	//TODO: add other contracts
	case "NFTContractWithBlacklist":
		panic("not implemented!")
		return nil, nil
	case "NFTContractWithWhitelist":
		panic("not implemented!")
		return nil, nil
	case "TokenContract":
		panic("not implemented!")
		return nil, nil
	case "TokenContractWithSplit":
		panic("not implemented!")
		return nil, nil
	case "TokenContractWithWhitelist":
		panic("not implemented!")
		return nil, nil
	case "TokenContractWithBlacklist":
		panic("not implemented!")
		return nil, nil
	}
	// Default fallback here before we implemented other contracts
	return NFTCtrt{}, nil
}
