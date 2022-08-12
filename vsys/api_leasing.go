package vsys

import "fmt"

type BroadcastLeasingPayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	Recipient    Str           `json:"recipient"`
	Amount       VSYS          `json:"amount"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Signature    Str           `json:"signature"`
}

type BroadcastLeaseTxResp struct {
	TxBasic

	Supernode Str  `json:"recipient"`
	Amount    VSYS `json:"amount"`
}

func (na *NodeAPI) BroadcastLease(p *BroadcastLeasingPayload) (*BroadcastLeaseTxResp, error) {
	res := &BroadcastLeaseTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/leasing/broadcast/lease")

	if err != nil {
		return nil, fmt.Errorf("BroadcastLease: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastLease: %s", resp.String())
	}
	return res, nil
}

func (b *BroadcastLeaseTxResp) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}

type BroadcastCancelLeasePayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	TxId         Str           `json:"txId"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Signature    Str           `json:"signature"`
}

type BroadcastCancelLeaseTxResp struct {
	TxBasic

	LeaseId Str `json:"leaseId"`
}

func (na *NodeAPI) BroadcastCancelLease(p *BroadcastCancelLeasePayload) (*BroadcastCancelLeaseTxResp, error) {
	res := &BroadcastCancelLeaseTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/leasing/broadcast/cancel")

	if err != nil {
		return nil, fmt.Errorf("BroadcastCancelLease: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastCancelLease: %s", resp.String())
	}
	return res, nil
}

func (b *BroadcastCancelLeasePayload) String() string {
	return fmt.Sprintf("%T(%+v)", b, *b)
}
