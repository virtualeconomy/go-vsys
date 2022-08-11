package vsys

import (
	"fmt"
)

// SysCtrt is the struct representing VSYS System contract.
type SysCtrt struct {
	*Ctrt
	tokId *TokenId
}

const (
	MAINNET_CTRT_ID = "CCL1QGBqPAaFjYiA8NMGVhzkd3nJkGeKYBq"
	TESTNET_CTRT_ID = "CF9Nd9wvQ8qVsGk8jYHbj6sf8TK7MJ2GYgt"
)

// NewSysCtrtForMainnet returns the SysCtrt instance for mainnet.
func NewSysCtrtForMainnet(ch *Chain) *SysCtrt {
	// No need to test for error since we set b58string
	ctrtIdMd, _ := NewCtrtIdFromB58Str(MAINNET_CTRT_ID)

	return &SysCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  ch,
		},
	}
}

// NewSysCtrtForTestnet returns the SysCtrt instance for testnet.
func NewSysCtrtForTestnet(ch *Chain) *SysCtrt {
	// No need to test for error since we set b58string
	ctrtIdMd, _ := NewCtrtIdFromB58Str(TESTNET_CTRT_ID)

	return &SysCtrt{
		Ctrt: &Ctrt{
			CtrtId: ctrtIdMd,
			Chain:  ch,
		},
	}
}

// TokId returns the token ID of the contract.
func (s *SysCtrt) TokId() (*TokenId, error) {
	if s.tokId == nil {
		var err error
		s.tokId, err = s.CtrtId.GetTokId(0)
		if err != nil {
			return nil, fmt.Errorf("TokId: %w", err)
		}
	}
	return s.tokId, nil
}

// Unit returns the unit of the token contract.
func (s *SysCtrt) Unit() (Unit, error) {
	return Unit(VSYS_UNIT), nil
}

// Send sends VSYS coins to another account
func (s *SysCtrt) Send(by *Account, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	rcptMd, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}
	rcptMd.MustOn(by.Chain)

	deAmount, err := NewDeAmountForTokAmount(amount, VSYS_UNIT)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		s.CtrtId,
		FUNC_IDX_SYS_CTRT_SEND,
		DataStack{NewDeAddr(rcptMd), deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}
	return resp, nil
}

// Transfer transfers tokens from sender to recipient.
func (s *SysCtrt) Transfer(by *Account, sender, recipient string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	senderMd, err := NewAddrFromB58Str(sender)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	rcptMd, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	senderMd.MustOn(by.Chain)
	rcptMd.MustOn(by.Chain)

	deAmount, err := NewDeAmountForTokAmount(amount, VSYS_UNIT)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		s.CtrtId,
		FUNC_IDX_SYS_CTRT_TRANSFER,
		DataStack{NewDeAddr(senderMd), NewDeAddr(rcptMd), deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Transfer: %w", err)
	}
	return resp, nil
}

// Deposit deposits the tokens into the contract.
func (s *SysCtrt) Deposit(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrt, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}
	senderMd, err := NewAddr(by.Addr.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, VSYS_UNIT)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		s.CtrtId,
		FUNC_IDX_SYS_CTRT_DEPOSIT,
		DataStack{NewDeAddr(senderMd), NewDeCtrtAddrFromCtrtId(ctrt), deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Deposit: %w", err)
	}
	return resp, nil
}

// Withdraw withdraws tokens from another contract.
func (s *SysCtrt) Withdraw(by *Account, ctrtId string, amount float64, attachment string) (*BroadcastExecuteTxResp, error) {
	ctrt, err := NewCtrtIdFromB58Str(ctrtId)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}
	rcptMd, err := NewAddr(by.Addr.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	deAmount, err := NewDeAmountForTokAmount(amount, VSYS_UNIT)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}

	txReq := NewExecCtrtFuncTxReq(
		s.CtrtId,
		FUNC_IDX_SYS_CTRT_WITHDRAW,
		DataStack{NewDeCtrtAddrFromCtrtId(ctrt), NewDeAddr(rcptMd), deAmount},
		NewVSYSTimestampForNow(),
		Str(attachment),
		FEE_EXEC_CTRT,
	)
	resp, err := by.ExecuteCtrt(txReq)
	if err != nil {
		return nil, fmt.Errorf("Withdraw: %w", err)
	}
	return resp, nil
}
