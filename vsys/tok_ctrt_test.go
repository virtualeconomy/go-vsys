package vsys

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type iTokCtrtTest interface {
	ctrtId() *CtrtId
	chain() *Chain
	Issuer() (*Addr, error)
	Maker() (*Addr, error)
	Issue(*Account, float64, string) (*BroadcastExecuteTxResp, error)
	Send(*Account, string, float64, string) (*BroadcastExecuteTxResp, error)
	Transfer(*Account, string, string, float64, string) (*BroadcastExecuteTxResp, error)
	Deposit(*Account, string, float64, string) (*BroadcastExecuteTxResp, error)
	Withdraw(*Account, string, float64, string) (*BroadcastExecuteTxResp, error)
}

type tokCtrtTest struct {
}

var tcT tokCtrtTest

func (tct *tokCtrtTest) newTokCtrtWithoutSplit(t *testing.T, by *Account) *TokCtrtWithoutSplit {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return tc
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitWithTok(t *testing.T, by *Account) *TokCtrtWithoutSplit {
	tc := tct.newTokCtrtWithoutSplit(t, by)
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc
}

func (tct *tokCtrtTest) test_TokCtrt_Register(t *testing.T, by *Account, tc iTokCtrtTest) {
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

	// if tokCtrtV2 need to check regulator
	if tcv2, ok := tc.(iTokCtrtV2Test); ok {
		regulator, err := tcv2.Regulator()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, by.Addr, regulator)
	}

}

func (tct *tokCtrtTest) test_TokCtrt_Issue(t *testing.T, by *Account, tc iTokCtrtTest) {
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokId, err := tc.ctrtId().GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tokBal, err := by.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 100, int(tokBal.Balance))
}

func (tct *tokCtrtTest) test_TokCtrt_Send(t *testing.T, sender, receiver *Account, tc iTokCtrtTest) {
	tokId, err := tc.ctrtId().GetTokId(0)
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

func (tct *tokCtrtTest) test_TokCtrt_Transfer(t *testing.T, sender, receiver *Account, tc iTokCtrtTest) {
	tokId, err := tc.ctrtId().GetTokId(0)
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

func (tct *tokCtrtTest) test_TokCtrt_DepositWithdraw(t *testing.T, by *Account, tc iTokCtrtTest) {
	tokId, err := tc.ctrtId().GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}

	ac, err := RegisterAtomicSwapCtrt(by, string(tokId.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}

	tok_bal, err := tc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal.Balance))

	// Need to add contract to whitelist before depositing.
	if tcv2w, ok := tc.(*TokCtrtWithoutSplitV2Whitelist); ok {
		resp, err := tcv2w.UpdateListCtrt(by, ac.CtrtId.B58Str().Str(), true, "")
		if err != nil {
			t.Fatal(err)
		}
		waitForBlock()
		assertTxSuccess(t, resp.Id.Str())
	}

	resp, err := tc.Deposit(by, string(ac.CtrtId.B58Str()), 5.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = tc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 95, int(tok_bal.Balance))

	depositedTokBal, err := ac.GetCtrtBal(string(by.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 5.0, depositedTokBal.Amount())

	resp, err = tc.Withdraw(by, string(ac.CtrtId.B58Str()), 5.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = tc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 100, int(tok_bal.Balance))

	depositedTokBal, _ = ac.GetCtrtBal(string(by.Addr.B58Str()))
	require.Equal(t, 0.0, depositedTokBal.Amount())
}

type iTokCtrtV1 interface {
	Supersede(*Account, string, string) (*BroadcastExecuteTxResp, error)
	Issuer() (*Addr, error)
}

func (tct *tokCtrtTest) test_TokCtrt_Supersede(t *testing.T, by, newIssuer *Account, tc iTokCtrtV1) {
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

func Test_TokCtrtWithoutSplit_Register(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplit(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplit_Issue(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplit(t, testAcnt0)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplit_Send(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplit_Transfer(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Transfer(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplit_DepositWithdraw(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplit_Supersede(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplit(t, testAcnt0)
	tcT.test_TokCtrt_Supersede(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplit_AsWhole(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplit(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)

	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
	tcT.test_TokCtrt_Transfer(t, testAcnt1, testAcnt0, tc)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
	tcT.test_TokCtrt_Supersede(t, testAcnt0, testAcnt1, tc)
}

func (tct *tokCtrtTest) newTokCtrtWithSplit(t *testing.T, by *Account) *TokCtrtWithSplit {
	tc, err := RegisterTokCtrtWithSplit(by, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return tc
}

func (tct *tokCtrtTest) newTokCtrtWithSplitWithTok(t *testing.T, by *Account) *TokCtrtWithSplit {
	tc := tct.newTokCtrtWithSplit(t, by)
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc
}

func Test_TokCtrtWithSplit_Register(t *testing.T) {
	tc := tcT.newTokCtrtWithSplit(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
}

func Test_TokCtrtWithSplit_Issue(t *testing.T) {
	tc := tcT.newTokCtrtWithSplit(t, testAcnt0)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)
}

func Test_TokCtrtWithSplit_Send(t *testing.T) {
	tc := tcT.newTokCtrtWithSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithSplit_Transfer(t *testing.T) {
	tc := tcT.newTokCtrtWithSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Transfer(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithSplit_DepositWithdraw(t *testing.T) {
	tc := tcT.newTokCtrtWithSplitWithTok(t, testAcnt0)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
}

func Test_TokCtrtWithSplit_Supersede(t *testing.T) {
	tc := tcT.newTokCtrtWithSplit(t, testAcnt0)
	tcT.test_TokCtrt_Supersede(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithSplit_AsWhole(t *testing.T) {
	tc := tcT.newTokCtrtWithSplit(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)

	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
	tcT.test_TokCtrt_Transfer(t, testAcnt1, testAcnt0, tc)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
	tcT.test_TokCtrt_Supersede(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithSplit_Split(t *testing.T) {
	tc := tcT.newTokCtrtWithSplit(t, testAcnt0)

	resp, err := tc.Split(testAcnt0, 12, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	newUnit, err := tc.Unit()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, Unit(12), newUnit)
}

type iTokCtrtV2Test interface {
	Issuer() (*Addr, error)
	Regulator() (*Addr, error)
	Supersede(*Account, string, string, string) (*BroadcastExecuteTxResp, error)
	IsUserInList(string) (bool, error)
	UpdateListUser(*Account, string, bool, string) (*BroadcastExecuteTxResp, error)
	IsCtrtInList(string) (bool, error)
	UpdateListCtrt(*Account, string, bool, string) (*BroadcastExecuteTxResp, error)
}

func (tct *tokCtrtTest) test_TokCtrtV2_Supersede(t *testing.T, tc iTokCtrtV2Test, by *Account, newIssuer, newRegulator *Account) {
	issuer, err := tc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, issuer)
	regulator, err := tc.Regulator()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, regulator)

	resp, err := tc.Supersede(by, string(newIssuer.Addr.B58Str()), string(newRegulator.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	issuer, err = tc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newIssuer.Addr, issuer)
	regulator, err = tc.Regulator()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newRegulator.Addr, regulator)
}

func (tct *tokCtrtTest) test_TokCtrtV2_UpdateListUser(t *testing.T, tc iTokCtrtV2Test, by, newUser *Account) {
	inList, err := tc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)

	resp, err := tc.UpdateListUser(
		by,
		string(newUser.Addr.B58Str()),
		true,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	inList, err = tc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, inList)

	resp, err = tc.UpdateListUser(
		by,
		string(newUser.Addr.B58Str()),
		false,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	inList, err = tc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)
}

func (tct *tokCtrtTest) test_TokCtrtV2_UpdateListCtrt(t *testing.T, tc iTokCtrtV2Test, by *Account, ctrtId string) {
	inList, err := tc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)

	resp, err := tc.UpdateListCtrt(
		by,
		ctrtId,
		true,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	inList, err = tc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, inList)

	resp, err = tc.UpdateListCtrt(
		by,
		ctrtId,
		false,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	inList, err = tc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitV2Whitelist(t *testing.T, by *Account) *TokCtrtWithoutSplitV2Whitelist {
	tc, err := RegisterTokCtrtWithoutSplitV2Whitelist(by, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	_, err = tc.UpdateListUser(testAcnt0, testAcnt0.Addr.B58Str().Str(), true, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.UpdateListUser(testAcnt0, testAcnt1.Addr.B58Str().Str(), true, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return tc
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitV2WhitelistWithTok(t *testing.T, by *Account) *TokCtrtWithoutSplitV2Whitelist {
	tc := tct.newTokCtrtWithoutSplitV2Whitelist(t, by)
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc
}

func Test_TokCtrtWithoutSplitV2Whitelist_Register(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Whitelist(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Whitelist_Supersede(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	tcT.test_TokCtrtV2_Supersede(t, tc, testAcnt0, testAcnt1, testAcnt1)
}

func Test_TokCtrtWithoutSplitV2Whitelist_UpdateListUser(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	tcT.test_TokCtrtV2_UpdateListUser(t, tc, testAcnt0, testAcnt2)
}

func Test_TokCtrtWithoutSplitV2Whitelist_UpdateListCtrt(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	// arbitraryCtrtId from nft ctrt test file.
	tcT.test_TokCtrtV2_UpdateListCtrt(t, tc, testAcnt0, arbitraryCtrtId(t))
}

func Test_TokCtrtWithoutSplitV2Whitelist_Issue(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Whitelist(t, testAcnt0)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Whitelist_Send(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplitV2Whitelist_Transfer(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Transfer(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplitV2Whitelist_DepositWithdraw(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2WhitelistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Whitelist_AsWhole(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Whitelist(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)

	tcT.test_TokCtrtV2_UpdateListUser(t, tc, testAcnt0, testAcnt2)
	tcT.test_TokCtrtV2_UpdateListCtrt(t, tc, testAcnt0, arbitraryCtrtId(t))

	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
	tcT.test_TokCtrt_Transfer(t, testAcnt1, testAcnt0, tc)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
	tcT.test_TokCtrtV2_Supersede(t, tc, testAcnt0, testAcnt1, testAcnt1)
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitV2Blacklist(t *testing.T, by *Account) *TokCtrtWithoutSplitV2Blacklist {
	tc, err := RegisterTokCtrtWithoutSplitV2Blacklist(by, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return tc
}

func (tct *tokCtrtTest) newTokCtrtWithoutSplitV2BlacklistWithTok(t *testing.T, by *Account) *TokCtrtWithoutSplitV2Blacklist {
	tc := tct.newTokCtrtWithoutSplitV2Blacklist(t, by)
	resp, err := tc.Issue(by, 100, "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc
}

func Test_TokCtrtWithoutSplitV2Blacklist_Register(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Blacklist(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Blacklist_Supersede(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	tcT.test_TokCtrtV2_Supersede(t, tc, testAcnt0, testAcnt1, testAcnt1)
}

func Test_TokCtrtWithoutSplitV2Blacklist_UpdateListUser(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	tcT.test_TokCtrtV2_UpdateListUser(t, tc, testAcnt0, testAcnt2)
}

func Test_TokCtrtWithoutSplitV2Blacklist_UpdateListCtrt(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	// arbitraryCtrtId from nft ctrt test file.
	tcT.test_TokCtrtV2_UpdateListCtrt(t, tc, testAcnt0, arbitraryCtrtId(t))
}

func Test_TokCtrtWithoutSplitV2Blacklist_Issue(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Blacklist(t, testAcnt0)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Blacklist_Send(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplitV2Blacklist_Transfer(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_Transfer(t, testAcnt0, testAcnt1, tc)
}

func Test_TokCtrtWithoutSplitV2Blacklist_DepositWithdraw(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2BlacklistWithTok(t, testAcnt0)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
}

func Test_TokCtrtWithoutSplitV2Blacklist_AsWhole(t *testing.T) {
	tc := tcT.newTokCtrtWithoutSplitV2Blacklist(t, testAcnt0)
	tcT.test_TokCtrt_Register(t, testAcnt0, tc)
	tcT.test_TokCtrt_Issue(t, testAcnt0, tc)

	tcT.test_TokCtrtV2_UpdateListUser(t, tc, testAcnt0, testAcnt2)
	tcT.test_TokCtrtV2_UpdateListCtrt(t, tc, testAcnt0, arbitraryCtrtId(t))

	tcT.test_TokCtrt_Send(t, testAcnt0, testAcnt1, tc)
	tcT.test_TokCtrt_Transfer(t, testAcnt1, testAcnt0, tc)
	tcT.test_TokCtrt_DepositWithdraw(t, testAcnt0, tc)
	tcT.test_TokCtrtV2_Supersede(t, tc, testAcnt0, testAcnt1, testAcnt1)
}
