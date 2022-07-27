package vsys

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type tokCtrtTest struct {
}

var tcT tokCtrtTest

func (tct *tokCtrtTest) newTokCtrtWithoutSplit(by *Account) (*TokCtrtWithoutSplit, error) {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	return tc, nil
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitWithTok(t *testing.T, by *Account) (*TokCtrtWithoutSplit, error) {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc, nil
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_Register(t *testing.T, by *Account) *TokCtrtWithoutSplit {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	issuer, err := tc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, issuer)
	maker, err := tc.Maker()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, maker)
	return tc
}

func Test_TokCtrtWithoutSplit_Register(t *testing.T) {
	tcT.test_TokCtrtWithoutSplit_Register(t, testAcnt0)
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_Issue(t *testing.T, by *Account, tc *TokCtrtWithoutSplit) {
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokId, err := tc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tokBal, err := by.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 100, int(tokBal.Balance))
}

func Test_TokCtrtWithoutSplit_Issue(t *testing.T) {
	tc, err := tcT.newTokCtrtWithoutSplit(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	tcT.test_TokCtrtWithoutSplit_Issue(t, testAcnt0, tc)
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_Send(t *testing.T, sender, receiver *Account, tc *TokCtrtWithoutSplit) {
	tokId, err := tc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tok_bal_acnt0, err := sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal_acnt0.Balance))
	tok_bal_acnt1, err := testAcnt0.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_acnt1.Balance))

	resp, err := tc.Send(sender, string(receiver.Addr.B58Str()), 100, "sending nft")
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
	require.Equal(t, 100, int(tok_bal_acnt1.Balance))
}

func Test_TokCtrtWithoutSplit_Send(t *testing.T) {
	tc, err := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	tcT.test_TokCtrtWithoutSplit_Send(t, testAcnt0, testAcnt1, tc)
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_Transfer(t *testing.T, sender, receiver *Account, tc *TokCtrtWithoutSplit) {
	tokId, err := tc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tok_bal_sender, err := sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal_sender.Balance))
	tok_bal_receiver, err := sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_receiver.Balance))

	resp, err := tc.Transfer(sender, string(sender.Addr.B58Str()), string(receiver.Addr.B58Str()), 100, "sending")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal_sender, err = sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_sender.Balance))
	tok_bal_receiver, err = sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal_receiver.Balance))
}

func Test_TokCtrtWithoutSplit_Transfer(t *testing.T) {
	tc, err := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	tcT.test_TokCtrtWithoutSplit_Transfer(t, testAcnt0, testAcnt1, tc)
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_DepositWithdraw(t *testing.T, by *Account, tc *TokCtrtWithoutSplit) {
	tokId, err := tc.CtrtId.GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}

	ac, err := RegisterAtomicSwapCtrt(by, string(tokId.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}

	tok_bal, err := tc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal.Balance))

	resp, err := tc.Deposit(by, string(ac.CtrtId.B58Str()), 5.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = tc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 95, int(tok_bal.Balance))

	depositedTokBal, _ := ac.GetCtrtBal(string(by.Addr.B58Str()))
	require.Equal(t, 5.0, depositedTokBal.Amount())

	resp, err = tc.Withdraw(by, string(ac.CtrtId.B58Str()), 5.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = tc.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal.Balance))

	depositedTokBal, _ = ac.GetCtrtBal(string(by.Addr.B58Str()))
	require.Equal(t, 0.0, depositedTokBal.Amount())
}

func Test_TokCtrtWithoutSplit_DepositWithdraw(t *testing.T) {
	tc, err := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	tcT.test_TokCtrtWithoutSplit_DepositWithdraw(t, testAcnt0, tc)
}

func (tct *tokCtrtTest) test_TokCtrtWithoutSplit_Supersede(t *testing.T, by, newIssuer *Account, tc *TokCtrtWithoutSplit) {
	resp, err := tc.Supersede(by, string(newIssuer.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	issuer, err := tc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newIssuer.Addr, issuer)
}

func Test_TokCtrtWithoutSplit_Supersede(t *testing.T) {
	tc, err := tcT.newTokCtrtWithoutSplit(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	tcT.test_TokCtrtWithoutSplit_Supersede(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplit_AsWhole(t *testing.T) {
	tc := test_NFTCtrt_Register(t, testAcnt0)
	test_NFTCtrt_Issue(t, testAcnt0, tc)

	test_NFTCtrt_Send(t, testAcnt0, testAcnt1, tc)
	test_NFTCtrt_Transfer(t, testAcnt1, testAcnt0, tc)
	test_NFTCtrt_DepositWithdraw(t, testAcnt0, tc)
	test_NFTCtrt_Supersede(t, testAcnt0, testAcnt1, tc)
}

// Do not test same functions for other Token Contracts since implementation is same.
// Test only unique functions for other token contracts.
