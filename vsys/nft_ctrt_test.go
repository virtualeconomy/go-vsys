package vsys

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func test_NFTCtrt_Register(t *testing.T, by *Account) *NFTCtrt {
	nc, err := RegisterNFTCtrt(by, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	issuer, err := nc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testAcnt0.Addr, issuer)
	maker, err := nc.Maker()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testAcnt0.Addr, maker)
	return nc
}

func Test_NFTCtrt_Register(t *testing.T) {
	test_NFTCtrt_Register(t, testAcnt0)
}

func newNFTCtrt(by *Account) (*NFTCtrt, error) {
	nc, err := RegisterNFTCtrt(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	return nc, nil
}

func newNFTCtrtWithTok(t *testing.T, by *Account) (*NFTCtrt, error) {
	nc, err := RegisterNFTCtrt(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	resp, err := nc.Issue(testAcnt0, "description", "attachment")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return nc, nil
}

func Test_NFTCtrt_Issue(t *testing.T) {
	nc, err := newNFTCtrt(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Issue(t, testAcnt0, nc)
}

func test_NFTCtrt_Issue(t *testing.T, by *Account, nc *NFTCtrt) {
	resp, err := nc.Issue(by, "description", "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokId, err := nc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tokBal, err := testAcnt0.Chain.NodeAPI.GetTokBal(string(testAcnt0.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, int(tokBal.Balance))
}

func Test_NFTCtrt_Send(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Send(t, testAcnt0, testAcnt1, nc)
}

func test_NFTCtrt_Send(t *testing.T, sender, receiver *Account, nc *NFTCtrt) {
	tokId, err := nc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tok_bal_acnt0, err := sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal_acnt0.Balance))
	tok_bal_acnt1, err := testAcnt0.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_acnt1.Balance))

	resp, err := nc.Send(sender, string(receiver.Addr.B58Str()), 0, "sending nft")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal_acnt0, err = sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_acnt0.Balance))
	tok_bal_acnt1, err = sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal_acnt1.Balance))
}

func Test_NFTCtrt_Transfer(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Transfer(t, testAcnt0, testAcnt1, nc)
}

func test_NFTCtrt_Transfer(t *testing.T, sender, receiver *Account, nc *NFTCtrt) {
	tokId, err := nc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tok_bal_acnt0, err := sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal_acnt0.Balance))
	tok_bal_acnt1, err := sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_acnt1.Balance))

	resp, err := nc.Transfer(sender, string(sender.Addr.B58Str()), string(receiver.Addr.B58Str()), 0, "sending nft")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal_acnt0, err = sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_acnt0.Balance))
	tok_bal_acnt1, err = sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal_acnt1.Balance))
}

func Test_NFTCtrt_DepositWithdraw(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_DepositWithdraw(t, testAcnt0, nc)
}

func test_NFTCtrt_DepositWithdraw(t *testing.T, by *Account, nc *NFTCtrt) {
	tokId, err := nc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}

	ac, err := RegisterAtomicSwapCtrt(by, string(tokId.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}

	tok_bal, err := nc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal.Balance))

	resp, err := nc.Deposit(by, string(ac.CtrtId.B58Str()), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = nc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal.Balance))

	depositedTokBal, _ := ac.GetCtrtBal(string(by.Addr.B58Str()))
	require.Equal(t, 1.0, depositedTokBal.Amount())

	resp, err = nc.Withdraw(by, string(ac.CtrtId.B58Str()), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = nc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal.Balance))

	depositedTokBal, _ = ac.GetCtrtBal(string(by.Addr.B58Str()))
	require.Equal(t, 0.0, depositedTokBal.Amount())
}

func Test_NFTCtrt_Supersede(t *testing.T) {
	nc, err := newNFTCtrt(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Supersede(t, testAcnt0, testAcnt1, nc)
}

func test_NFTCtrt_Supersede(t *testing.T, by, newIssuer *Account, nc *NFTCtrt) {
	resp, err := nc.Supersede(by, string(newIssuer.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	issuer, err := nc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newIssuer.Addr, issuer)
}

func Test_NFTCtrt_AsWhole(t *testing.T) {
	nc := test_NFTCtrt_Register(t, testAcnt0)
	test_NFTCtrt_Issue(t, testAcnt0, nc)

	test_NFTCtrt_Send(t, testAcnt0, testAcnt1, nc)
	test_NFTCtrt_Transfer(t, testAcnt1, testAcnt0, nc)
	test_NFTCtrt_DepositWithdraw(t, testAcnt0, nc)
	test_NFTCtrt_Supersede(t, testAcnt0, testAcnt1, nc)
}
