package vsys

import (
	"fmt"
)

type BalDetails struct {
	Addr       Str    `json:"address"`
	Regular    VSYS   `json:"regular"`
	MintingAvg VSYS   `json:"mintingAverage"`
	Available  VSYS   `json:"available"`
	Effective  VSYS   `json:"effective"`
	Height     Height `json:"height"`
}

func (na *NodeAPI) GetBalDetails(addr string) (*BalDetails, error) {
	res := &BalDetails{}
	resp, err := na.R().SetResult(res).Get("/addresses/balance/details/" + addr)
	if err != nil {
		return nil, fmt.Errorf("GetBalDetails: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetBalDetails: %s", resp.String())
	}
	return res, nil
}
