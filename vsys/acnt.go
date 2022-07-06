package vsys

import (
	"fmt"
)

type Wallet struct {
	Seed *Seed
}

func NewWallet(s *Seed) *Wallet {
	return &Wallet{Seed: s}
}

func NewWalletFromSeedStr(s string) (*Wallet, error) {
	seed, err := NewSeed(s)
	if err != nil {
		return nil, fmt.Errorf("NewWalletFromSeedStr: %w", err)
	}
	return NewWallet(seed), nil
}

func GenWallet() (*Wallet, error) {
	seed, err := NewRandSeed()
	if err != nil {
		return nil, err
	}
	return NewWallet(seed), nil
}

func (w *Wallet) GetAccount(c *Chain, n Nonce) (*Account, error) {
	acntSeedHash := GetAccountSeedHash(w.Seed, n)

	priKey, err := NewPriKeyFromRand(acntSeedHash)

	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}

	ac, err := NewAccount(c, priKey)
	if err != nil {
		return nil, fmt.Errorf("GetAccount: %w", err)
	}

	return ac, nil
}

func GetAccountSeedHash(s *Seed, n Nonce) Bytes {
	b := Sha256Hash(
		Keccak256Hash(
			Blake2bHash([]byte(fmt.Sprintf("%d%s", n, s.Str.Str()))),
		),
	)

	return Bytes(b)
}

func (w *Wallet) String() string {
	return fmt.Sprintf("%T(%+v)", w, *w)
}

type Account struct {
	Chain  *Chain
	PriKey *PriKey
	PubKey *PubKey
	Addr   *Addr
}

func NewAccount(c *Chain, pri *PriKey) (*Account, error) {
	pub, err := NewPubKeyFromPriKey(pri)
	if err != nil {
		return nil, fmt.Errorf("NewAccount: %w", err)
	}

	addr, err := NewAddrFromPubKey(pub, c.ChainID)
	if err != nil {
		return nil, fmt.Errorf("NewAccount: %w", err)
	}

	return &Account{
		Chain:  c,
		PriKey: pri,
		PubKey: pub,
		Addr:   addr,
	}, nil
}

func (a *Account) API() *NodeAPI {
	return a.Chain.NodeAPI
}

func (a *Account) Bal() (VSYS, error) {
	res, err := a.API().GetBalDetails(a.Addr.B58Str().Str())
	if err != nil {
		return 0, fmt.Errorf("Bal: %w", err)
	}
	return res.Regular, nil
}

func (a *Account) Pay(
	recipient string,
	amount float64,
	attachment string,
) (*BroadcastPaymentTxResp, error) {
	rcptMd, err := NewAddrFromB58Str(recipient)
	if err != nil {
		return nil, fmt.Errorf("Pay: %w", err)
	}

	amntMd, err := NewVSYSForAmount(amount)
	if err != nil {
		return nil, fmt.Errorf("Pay: %w", err)
	}

	tsMd := NewVSYSTimestampForNow()

	txReq := NewPaymentTxReq(
		rcptMd,
		amntMd,
		tsMd,
		Str(attachment),
		FEE_PAYMENT,
	)

	payload, err := txReq.BroadcastPaymentPayload(a.PriKey, a.PubKey)
	if err != nil {
		return nil, fmt.Errorf("Pay: %w", err)
	}

	resp, err := a.API().BroadcastPayment(payload)

	if err != nil {
		return nil, fmt.Errorf("Pay: %w", err)
	}
	return resp, nil
}

func (a *Account) RegisterCtrt(txReq *RegCtrtTxReq) (*BroadcastRegisterTxResp, error) {
	payload, err := txReq.BroadcastRegisterPayload(a.PriKey, a.PubKey)
	if err != nil {
		return nil, fmt.Errorf("RegisterCtrt: %w", err)
	}

	resp, err := a.API().BroadcastRegister(payload)

	if err != nil {
		return nil, fmt.Errorf("RegisterCtrt: %w", err)
	}
	return resp, err
}

func (a *Account) ExecuteCtrt(txReq *ExecCtrtFuncTxReq) (*BroadcastExecuteTxResp, error) {
	payload, err := txReq.BroadcastExecutePayload(a.PriKey, a.PubKey)
	if err != nil {
		return nil, fmt.Errorf("ExecuteCtrt: %w", err)
	}

	resp, err := a.API().BroadcastExecute(payload)

	if err != nil {
		return nil, fmt.Errorf("ExecuteCtrt: %w", err)
	}
	return resp, err
}

func (a *Account) String() string {
	return fmt.Sprintf("%T(%+v)", a, *a)
}
