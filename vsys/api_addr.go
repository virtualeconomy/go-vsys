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

// GetBalDetails gets the balance details of the given address.
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

func (b *BalDetails) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}

type AddrResp struct {
	Address Str `json:"address"`
}

// GetAddr gets the address from the public key.
func (na *NodeAPI) GetAddr(pubKey string) (*AddrResp, error) {
	res := &AddrResp{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/addresses/publicKey/%s", pubKey))
	if err != nil {
		return nil, fmt.Errorf("GetAddr: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetAddr: %s", resp.String())
	}
	return res, nil
}

func (a *AddrResp) String() string {
	return fmt.Sprintf("%T(%+v)", a, *a)
}

type BalResp struct {
	Address       Str  `json:"address"`
	Confirmations int  `json:"confirmations"`
	Balance       VSYS `json:"balance"`
}

// GetBal gets the ledger(regular) balance of the given address.
func (na *NodeAPI) GetBal(addr string) (*BalResp, error) {
	res := &BalResp{}
	resp, err := na.R().SetResult(res).Get(fmt.Sprintf("/addresses/balance/%s", addr))
	if err != nil {
		return nil, fmt.Errorf("GetBal: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("GetBal: %s", resp.String())
	}
	return res, nil
}

func (b *BalResp) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}
