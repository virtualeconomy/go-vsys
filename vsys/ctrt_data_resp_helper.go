package vsys

import (
	"fmt"
)

func ctrtDataRespToAddr(resp *CtrtDataResp) (*Addr, error) {
	switch val := resp.Val.(type) {
	case string:
		addr, err := NewAddrFromB58Str(val)
		if err != nil {
			return nil, fmt.Errorf("ctrtDataRespToAddr: %w", err)
		}
		return addr, nil
	default:
		return nil, fmt.Errorf("ctrtDataRespToAddr: Val is %T but string was expected", val)
	}
}

func ctrtDataRespToToken(resp *CtrtDataResp, unit Unit) (*Token, error) {
	switch val := resp.Val.(type) {
	case float64:
		return NewToken(Amount(val), unit), nil
	default:
		return nil, fmt.Errorf("ctrtDataRespToAddr: Val is %T but float64 was expected", val)
	}
}

func ctrtDataRespToTokenId(resp *CtrtDataResp) (*TokenId, error) {
	switch tokId := resp.Val.(type) {
	case string:
		tokIdMd, err := NewTokenIdFromB58Str(tokId)
		if err != nil {
			return nil, fmt.Errorf("ctrtDataRespToTokenId: %w", err)
		}
		return tokIdMd, nil
	default:
		return nil, fmt.Errorf("ctrtDataRespToTokenId: CtrtDataResp.Val is %T but string was expected", tokId)
	}
}
