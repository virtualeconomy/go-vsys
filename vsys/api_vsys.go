package vsys

import (
	"fmt"
)

type BroadcastPaymentPayload struct {
	SenderPubKey Str           `json:"senderPublicKey"`
	Recipient    Str           `json:"recipient"`
	Amount       VSYS          `json:"amount"`
	Fee          VSYS          `json:"fee"`
	FeeScale     FeeScale      `json:"feeScale"`
	Timestamp    VSYSTimestamp `json:"timestamp"`
	Attachment   Str           `json:"attachment"`
	Signature    Str           `json:"signature"`
}

type BroadcastPaymentTxResp struct {
	TxBasic

	Recipient  Str  `json:"recipient"`
	Amount     VSYS `json:"amount"`
	Attachment Str  `json:"attachment"`
}

func (na *NodeAPI) BroadcastPayment(p *BroadcastPaymentPayload) (*BroadcastPaymentTxResp, error) {
	res := &BroadcastPaymentTxResp{}
	resp, err := na.R().
		SetBody(p).
		SetResult(res).
		Post("/vsys/broadcast/payment")

	if err != nil {
		return nil, fmt.Errorf("BroadcastPayment: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("BroadcastPayment: %s", resp.String())
	}
	return res, nil
}
