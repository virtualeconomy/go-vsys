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

// BaseTokCtrt is the general interface for Token certificates
type BaseTokCtrt interface {
	Unit() (Unit, error)
}

// GetCtrtFromTokId returns instance of token contract corresponding to given tokenId
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
		return nil, fmt.Errorf("not implemented!")
	case "NFTContractWithWhitelist":
		return nil, fmt.Errorf("not implemented!")
	case "TokenContract":
		n, err := NewTokCtrtWithoutSplit(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
		}
		return n, nil
	case "TokenContractWithSplit":
		return nil, fmt.Errorf("not implemented!")
	case "TokenContractWithWhitelist":
		return nil, fmt.Errorf("not implemented!")
	case "TokenContractWithBlacklist":
		return nil, fmt.Errorf("not implemented!")
	default:
		return nil, fmt.Errorf("contract type unexpected: %s", ctrtInfo.Type)
	}
}
