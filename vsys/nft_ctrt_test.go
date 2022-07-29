package vsys

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type iNFTtest interface {
	ctrtId() *CtrtId
	chain() *Chain
	Issuer() (*Addr, error)
	Maker() (*Addr, error)
	Issue(*Account, string, string) (*BroadcastExecuteTxResp, error)
	Send(*Account, string, int, string) (*BroadcastExecuteTxResp, error)
	Transfer(*Account, string, string, int, string) (*BroadcastExecuteTxResp, error)
	Deposit(*Account, string, int, string) (*BroadcastExecuteTxResp, error)
	Withdraw(*Account, string, int, string) (*BroadcastExecuteTxResp, error)
}

func arbitraryCtrtId(t *testing.T) string {
	as, _ := asT.newAtomicSwap(t, testAcnt0)
	return string(as.CtrtId.B58Str())
}

func test_NFTCtrt_Register(t *testing.T, nc iNFTtest, by *Account) iNFTtest {
	issuer, err := nc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, issuer)
	maker, err := nc.Maker()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, maker)
	// if contract is nftv2 need to check regulator
	ncv2, ok := nc.(iNFTv2Test)
	if ok {
		regulator, err := ncv2.Regulator()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, by.Addr, regulator)
	}
	return nc
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
	resp, err := nc.Issue(by, "description", "attachment")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return nc, nil
}

func test_NFTCtrt_Issue(t *testing.T, nc iNFTtest, by *Account) {
	resp, err := nc.Issue(by, "description", "attachment")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokId, err := nc.ctrtId().GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tokBal, err := by.Chain.NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, int(tokBal.Balance))
}

func test_NFTCtrt_Send(t *testing.T, nc iNFTtest, sender, receiver *Account) {
	tokId, err := nc.ctrtId().GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}
	tok_bal_sender, err := sender.Chain.NodeAPI.GetTokBal(string(sender.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal_sender.Balance))
	tok_bal_receiver, err := sender.Chain.NodeAPI.GetTokBal(string(receiver.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal_receiver.Balance))

	resp, err := nc.Send(sender, string(receiver.Addr.B58Str()), 0, "sending nft")
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
	require.Equal(t, 1, int(tok_bal_receiver.Balance))
}

func test_NFTCtrt_Transfer(t *testing.T, nc iNFTtest, sender, receiver *Account) {
	tokId, err := nc.ctrtId().GetTokId(0)
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

func test_NFTCtrt_DepositWithdraw(t *testing.T, nc iNFTtest, by *Account) {
	tokId, err := nc.ctrtId().GetTokId(0)
	if err != nil {
		t.Fatal(err)
	}

	ac, err := RegisterAtomicSwapCtrt(by, string(tokId.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}

	tok_bal, err := nc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal.Balance))

	// Contract need to be added to whitelist before proceeding.
	if ncwl, ok := nc.(*NFTCtrtV2Whitelist); ok {
		_, err := ncwl.UpdateListCtrt(by, string(ac.CtrtId.B58Str()), true, "")
		if err != nil {
			t.Fatal(err)
		}
		waitForBlock()
	}

	resp, err := nc.Deposit(by, string(ac.CtrtId.B58Str()), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = nc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, int(tok_bal.Balance))

	depositedTokBal, err := ac.GetCtrtBal(string(by.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1.0, depositedTokBal.Amount())

	resp, err = nc.Withdraw(by, string(ac.CtrtId.B58Str()), 0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tok_bal, err = nc.chain().NodeAPI.GetTokBal(string(by.Addr.B58Str()), string(tokId.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1, int(tok_bal.Balance))

	depositedTokBal, err = ac.GetCtrtBal(string(by.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0.0, depositedTokBal.Amount())
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

func Test_NFTCtrt_Register(t *testing.T) {
	nc, err := RegisterNFTCtrt(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	test_NFTCtrt_Register(t, nc, testAcnt0)
}

func Test_NFTCtrt_Issue(t *testing.T) {
	nc, err := newNFTCtrt(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Issue(t, nc, testAcnt0)
}

func Test_NFTCtrt_Send(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrt_Transfer(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Transfer(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrt_DepositWithdraw(t *testing.T) {
	nc, err := newNFTCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
}

func Test_NFTCtrt_Supersede(t *testing.T) {
	nc, err := newNFTCtrt(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Supersede(t, testAcnt0, testAcnt1, nc)
}

func Test_NFTCtrt_AsWhole(t *testing.T) {
	nc, err := RegisterNFTCtrt(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	nc = test_NFTCtrt_Register(t, nc, testAcnt0).(*NFTCtrt)
	test_NFTCtrt_Issue(t, nc, testAcnt0)

	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
	test_NFTCtrt_Transfer(t, nc, testAcnt1, testAcnt0)
	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
	test_NFTCtrt_Supersede(t, testAcnt0, testAcnt1, nc)
}

// NFT V2 Tests

type iNFTv2Test interface {
	Issuer() (*Addr, error)
	Regulator() (*Addr, error)
	Supersede(*Account, string, string, string) (*BroadcastExecuteTxResp, error)
	IsUserInList(string) (bool, error)
	UpdateListUser(*Account, string, bool, string) (*BroadcastExecuteTxResp, error)
	IsCtrtInList(string) (bool, error)
	UpdateListCtrt(*Account, string, bool, string) (*BroadcastExecuteTxResp, error)
}

func test_NFTCtrtV2_Supersede(t *testing.T, nc iNFTv2Test, by *Account, newIssuer, newRegulator *Account) {
	issuer, err := nc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, issuer)
	regulator, err := nc.Regulator()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, by.Addr, regulator)

	resp, err := nc.Supersede(by, string(newIssuer.Addr.B58Str()), string(newRegulator.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	issuer, err = nc.Issuer()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newIssuer.Addr, issuer)
	regulator, err = nc.Regulator()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, newRegulator.Addr, regulator)
}

func test_NFTCtrtV2_UpdateListUser(t *testing.T, nc iNFTv2Test, by, newUser *Account) {
	inList, err := nc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)

	resp, err := nc.UpdateListUser(
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

	inList, err = nc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, inList)

	resp, err = nc.UpdateListUser(
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

	inList, err = nc.IsUserInList(string(newUser.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)
}

func test_NFTCtrtV2_UpdateListCtrt(t *testing.T, nc iNFTv2Test, by *Account, ctrtId string) {
	inList, err := nc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)

	resp, err := nc.UpdateListCtrt(
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

	inList, err = nc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, inList)

	resp, err = nc.UpdateListCtrt(
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

	inList, err = nc.IsCtrtInList(ctrtId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, inList)
}

func newNFTCtrtV2Whitelist(by *Account) (*NFTCtrtV2Whitelist, error) {
	nc, err := RegisterNFTCtrtV2Whitelist(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	return nc, nil
}

func newNFTCtrtV2WhitelistWithTok(t *testing.T, by *Account) (*NFTCtrtV2Whitelist, error) {
	nc, err := RegisterNFTCtrtV2Whitelist(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	resp, err := nc.UpdateListUser(by, string(by.Addr.B58Str()), true, "")
	resp, err = nc.UpdateListUser(by, string(testAcnt1.Addr.B58Str()), true, "")
	if err != nil {
		return nil, err
	}
	resp1, err := nc.Issue(by, "description", "attachment")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	assertTxSuccess(t, string(resp1.Id))
	return nc, nil
}

func Test_NFTCtrtV2Whitelist_Register(t *testing.T) {
	nc, err := RegisterNFTCtrtV2Whitelist(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	test_NFTCtrt_Register(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Whitelist_Issue(t *testing.T) {
	nc, err := newNFTCtrtV2Whitelist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Issue(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Whitelist_Send(t *testing.T) {
	nc, err := newNFTCtrtV2WhitelistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Whitelist_Transfer(t *testing.T) {
	nc, err := newNFTCtrtV2WhitelistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Transfer(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Whitelist_DepositWithdraw(t *testing.T) {
	nc, err := newNFTCtrtV2WhitelistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Whitelist_Supersede(t *testing.T) {
	nc, err := newNFTCtrtV2Whitelist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_Supersede(t, nc, testAcnt0, testAcnt1, testAcnt1)
}

func Test_NFTCtrtV2Whitelist_UpdateListUser(t *testing.T) {
	nc, err := newNFTCtrtV2Whitelist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_UpdateListUser(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Whitelist_UpdateListCtrt(t *testing.T) {
	nc, err := newNFTCtrtV2Whitelist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_UpdateListCtrt(t, nc, testAcnt0, arbitraryCtrtId(t))
}

func Test_NFTCtrtV2Whitelist_AsWhole(t *testing.T) {
	nc, err := RegisterNFTCtrtV2Whitelist(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	test_NFTCtrt_Register(t, nc, testAcnt0)
	test_NFTCtrtV2_UpdateListUser(t, nc, testAcnt0, testAcnt2)
	test_NFTCtrtV2_UpdateListCtrt(t, nc, testAcnt0, arbitraryCtrtId(t))

	test_NFTCtrt_Issue(t, nc, testAcnt0)

	nc, err = newNFTCtrtV2WhitelistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
	test_NFTCtrt_Transfer(t, nc, testAcnt1, testAcnt0)

	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
	test_NFTCtrtV2_Supersede(t, nc, testAcnt0, testAcnt1, testAcnt1)
}

func newNFTCtrtV2Blacklist(by *Account) (*NFTCtrtV2Blacklist, error) {
	nc, err := RegisterNFTCtrtV2Blacklist(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	return nc, nil
}

func newNFTCtrtV2BlacklistWithTok(t *testing.T, by *Account) (*NFTCtrtV2Blacklist, error) {
	nc, err := RegisterNFTCtrtV2Blacklist(by, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()

	resp, err := nc.Issue(by, "description", "attachment")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return nc, nil
}

func Test_NFTCtrtV2Blacklist_Register(t *testing.T) {
	nc, err := RegisterNFTCtrtV2Blacklist(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	test_NFTCtrt_Register(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Blacklist_Issue(t *testing.T) {
	nc, err := newNFTCtrtV2Blacklist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Issue(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Blacklist_Send(t *testing.T) {
	nc, err := newNFTCtrtV2BlacklistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Blacklist_Transfer(t *testing.T) {
	nc, err := newNFTCtrtV2BlacklistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_Transfer(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Blacklist_DepositWithdraw(t *testing.T) {
	nc, err := newNFTCtrtV2BlacklistWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
}

func Test_NFTCtrtV2Blacklist_Supersede(t *testing.T) {
	nc, err := newNFTCtrtV2Blacklist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_Supersede(t, nc, testAcnt0, testAcnt1, testAcnt1)
}

func Test_NFTCtrtV2Blacklist_UpdateListUser(t *testing.T) {
	nc, err := newNFTCtrtV2Blacklist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_UpdateListUser(t, nc, testAcnt0, testAcnt1)
}

func Test_NFTCtrtV2Blacklist_UpdateListCtrt(t *testing.T) {
	nc, err := newNFTCtrtV2Blacklist(testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_NFTCtrtV2_UpdateListCtrt(t, nc, testAcnt0, arbitraryCtrtId(t))
}

func Test_NFTCtrtV2Blacklist_AsWhole(t *testing.T) {
	nc, err := RegisterNFTCtrtV2Blacklist(testAcnt0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	test_NFTCtrt_Register(t, nc, testAcnt0)
	test_NFTCtrtV2_UpdateListUser(t, nc, testAcnt0, testAcnt1)
	test_NFTCtrtV2_UpdateListCtrt(t, nc, testAcnt0, arbitraryCtrtId(t))

	test_NFTCtrt_Issue(t, nc, testAcnt0)
	test_NFTCtrt_Send(t, nc, testAcnt0, testAcnt1)
	test_NFTCtrt_Transfer(t, nc, testAcnt1, testAcnt0)

	test_NFTCtrt_DepositWithdraw(t, nc, testAcnt0)
	test_NFTCtrtV2_Supersede(t, nc, testAcnt0, testAcnt1, testAcnt1)
}
