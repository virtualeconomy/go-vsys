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
	// Check if token id is vsys system contract
	sysCtrt := NewSysCtrt(chain)
	sysCtrtTokId, err := sysCtrt.TokId()
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
	}
	if string(tokId.B58Str()) == string(sysCtrtTokId.B58Str()) {
		return sysCtrt, nil
	}

	tokInfo, err := chain.NodeAPI.GetTokInfo(string(tokId.B58Str()))
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
	}
	ctrt, err := GetCtrtFromCtrtId(tokInfo.CtrtId.Str(), chain)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromTokId: %w", err)
	}
	return ctrt, nil
}

// GetCtrtFromCtrtId return instance of token contract corresponding to given tokenId
func GetCtrtFromCtrtId(ctrtId string, chain *Chain) (BaseTokCtrt, error) {
	ctrtInfo, err := chain.NodeAPI.GetCtrtInfo(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
	}
	switch string(ctrtInfo.Type) {
	case "NonFungibleContract":
		n, err := NewNFTCtrt(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "NFTContractWithBlacklist":
		n, err := NewNFTCtrtV2Blacklist(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "NFTContractWithWhitelist":
		n, err := NewNFTCtrtV2Whitelist(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "TokenContract":
		n, err := NewTokCtrtWithoutSplit(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "TokenContractWithSplit":
		n, err := NewTokCtrtWithSplit(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "TokenContractWithWhitelist":
		n, err := NewTokCtrtWithoutSplitV2Whitelist(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	case "TokenContractWithBlacklist":
		n, err := NewTokCtrtWithoutSplitV2Blacklist(ctrtInfo.CtrtId.Str(), chain)
		if err != nil {
			return nil, fmt.Errorf("GetCtrtFromCtrtId: %w", err)
		}
		return n, nil
	default:
		return nil, fmt.Errorf("contract type unexpected: %s", ctrtInfo.Type)
	}
}
