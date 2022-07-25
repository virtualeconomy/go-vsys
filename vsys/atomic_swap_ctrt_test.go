package vsys

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func newTokCtrtWithTok(t *testing.T, by *Account) (*TokCtrtWithoutSplit, error) {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		return nil, fmt.Errorf("newMakerTokCtrtWithTok: %w", err)
	}
	waitForBlock()

	resp, err := tc.Issue(by, 1000, "")
	if err != nil {
		return nil, fmt.Errorf("newMakerTokCtrtWithTok: %w", err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return tc, nil
}

func newAtomicSwap(t *testing.T, by *Account) (*AtomicSwapCtrt, error) {
	token, err := newTokCtrtWithTok(t, by)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	tokId, _ := token.TokId()
	ac, err := RegisterAtomicSwapCtrt(by, string(tokId.B58Str()), "")
	if err != nil {
		return nil, err
	}
	waitForBlock()

	resp, err := token.Deposit(by, string(ac.CtrtId.B58Str()), 1000, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	return ac, nil
}

func test_AtomicSwapCtrt_Register(t *testing.T, acnt *Account, tc *TokCtrtWithoutSplit) *AtomicSwapCtrt {
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	ac, err := RegisterAtomicSwapCtrt(acnt, string(tokId.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	maker, err := ac.Maker()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, acnt.Addr, maker)
	tokIdFromCtrt, err := ac.TokId()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, tokId, tokIdFromCtrt)
	return ac
}

func Test_AtomicSwapCtrt_Register(t *testing.T) {
	tc, err := newTokCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatalf("Cannot get new maker token ctrt: %s\n", err.Error())
	}
	test_AtomicSwapCtrt_Register(t, testAcnt0, tc)
}

func test_AtomicSwapCtrt_Lock(t *testing.T, maker, taker *Account, makerCtrt, takerCtrt *AtomicSwapCtrt) {
	makerBalInit, _ := makerCtrt.GetCtrtBal(string(maker.Addr.B58Str()))
	takerBalInit, _ := takerCtrt.GetCtrtBal(string(taker.Addr.B58Str()))

	makerLockAmount := 10.0
	makerLockTimestamp := time.Now().Unix() + 1800
	makerPuzzlePlain := "abc"
	puzzleBytes := Sha256Hash([]byte(makerPuzzlePlain))

	makerLockTxInfo, err := makerCtrt.Lock(maker, makerLockAmount, string(taker.Addr.B58Str()), puzzleBytes, makerLockTimestamp, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	makerLockTxId := string(makerLockTxInfo.Id)
	assertTxSuccess(t, makerLockTxId)

	makerSwapOwner, _ := makerCtrt.GetSwapOwner(makerLockTxId)
	require.Equal(t, maker.Addr, makerSwapOwner)
	makerSwapRecipient, _ := makerCtrt.GetSwapRecipient(makerLockTxId)
	require.Equal(t, taker.Addr, makerSwapRecipient)
	makerSwapAmount, _ := makerCtrt.GetSwapAmount(makerLockTxId)
	require.Equal(t, makerLockAmount, makerSwapAmount.Amount())
	makerSwapTimestamp, _ := makerCtrt.GetSwapExpiredTime(makerLockTxId)
	require.Equal(t, makerLockTimestamp, makerSwapTimestamp.UnixTs())
	makerSwapStatus, _ := makerCtrt.GetSwapStatus(makerLockTxId)
	require.Equal(t, true, makerSwapStatus)
	makerPuzzle, _ := makerCtrt.GetSwapPuzzle(makerLockTxId)
	decoded, err := B58Decode(string(makerPuzzle))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, Sha256Hash([]byte(makerPuzzlePlain)), decoded)
	makerBalAfterLock, _ := makerCtrt.GetCtrtBal(string(maker.Addr.B58Str()))
	require.Equal(t, makerBalInit.Amount()-makerLockAmount, makerBalAfterLock.Amount())

	// Taker lock here
	takerLockAmount := 5.0
	takerLockTimestamp := time.Now().Unix() + 1500
	takerLockTxInfo, err := takerCtrt.Lock(
		taker,
		takerLockAmount,
		string(maker.Addr.B58Str()),
		decoded, // Got decoded puzzle bytes from previous checks
		takerLockTimestamp,
		"",
	)
	if err != nil {
		t.Fatal(t)
	}
	waitForBlock()
	assertTxSuccess(t, string(takerLockTxInfo.Id))
	takerBalAfterLock, _ := takerCtrt.GetCtrtBal(string(taker.Addr.B58Str()))
	require.Equal(t, takerBalInit.Amount()-takerLockAmount, takerBalAfterLock.Amount())
}

func Test_AtomicSwapCtrt_Lock(t *testing.T) {
	g := new(errgroup.Group)
	var makerCtrt, takerCtrt *AtomicSwapCtrt
	g.Go(func() error {
		var err error
		makerCtrt, err = newAtomicSwap(t, testAcnt0)
		return err
	})
	g.Go(func() error {
		var err error
		takerCtrt, err = newAtomicSwap(t, testAcnt1)
		return err
	})
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
	test_AtomicSwapCtrt_Lock(t, testAcnt0, testAcnt1, makerCtrt, takerCtrt)
}

func test_AtomicSwapCtrt_Solve(t *testing.T, maker, taker *Account, makerCtrt, takerCtrt *AtomicSwapCtrt) {
	// Maker lock
	makerLockAmount := 10.0
	makerLockTimestamp := time.Now().Unix() + 1800
	makerPuzzlePlain := "abc"
	puzzleBytes := Sha256Hash([]byte(makerPuzzlePlain))

	makerLockTxInfo, err := makerCtrt.Lock(maker, makerLockAmount, string(taker.Addr.B58Str()), puzzleBytes, makerLockTimestamp, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	makerLockTxId := string(makerLockTxInfo.Id)
	assertTxSuccess(t, makerLockTxId)

	makerPuzzle, _ := makerCtrt.GetSwapPuzzle(makerLockTxId)
	puzzleBytes2, err := B58Decode(string(makerPuzzle))
	if err != nil {
		t.Fatal(err)
	}

	// Taker lock here
	takerLockAmount := 5.0
	takerLockTimestamp := time.Now().Unix() + 1500
	takerLockTxInfo, err := takerCtrt.Lock(
		taker,
		takerLockAmount,
		string(maker.Addr.B58Str()),
		puzzleBytes2,
		takerLockTimestamp,
		"",
	)
	waitForBlock()
	assertTxSuccess(t, string(takerLockTxInfo.Id))

	// maker solve
	makerSolveTxInfo, err := takerCtrt.Solve(maker, string(takerLockTxInfo.Id), makerPuzzlePlain, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(makerSolveTxInfo.Id))

	dict_data, err := maker.Chain.NodeAPI.GetTxInfo(string(makerSolveTxInfo.Id))
	if err != nil {
		t.Fatal(err)
	}
	funcData := dict_data.(*ExecCtrtFuncTxInfoResp).FuncData
	decoded, err := B58Decode(string(funcData))
	if err != nil {
		t.Fatal(err)
	}
	ds, err := NewDataStackFromBytes(decoded)
	if err != nil {
		t.Fatal(err)
	}
	revealedSecret := string(ds[1].DataBytes())
	require.Equal(t, makerPuzzlePlain, revealedSecret)

	// taker solve
	takerSolveTxInfo, err := makerCtrt.Solve(taker, makerLockTxId, revealedSecret, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(takerSolveTxInfo.Id))
}

func Test_AtomicSwapCtrt_Solve(t *testing.T) {
	g := new(errgroup.Group)
	var makerCtrt, takerCtrt *AtomicSwapCtrt
	g.Go(func() error {
		var err error
		makerCtrt, err = newAtomicSwap(t, testAcnt0)
		return err
	})
	g.Go(func() error {
		var err error
		takerCtrt, err = newAtomicSwap(t, testAcnt1)
		return err
	})
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
	test_AtomicSwapCtrt_Solve(t, testAcnt0, testAcnt1, makerCtrt, takerCtrt)
}

func test_AtomicSwapCtrt_ExpWithdraw(t *testing.T, acnt0, acnt1 *Account, makerCtrt *AtomicSwapCtrt) {
	makerLockAmount := 10.0
	makerLockTimestamp := time.Now().Unix() + 8
	makerPuzzlePlain := "abc"
	puzzleBytes := Sha256Hash([]byte(makerPuzzlePlain))

	// Maker lock
	makerLockTxInfo, err := makerCtrt.Lock(acnt0, makerLockAmount, string(acnt1.Addr.B58Str()), puzzleBytes, makerLockTimestamp, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	makerLockTxId := string(makerLockTxInfo.Id)
	assertTxSuccess(t, makerLockTxId)

	bal_old, err := makerCtrt.GetCtrtBal(string(acnt0.Addr.B58Str()))

	time.Sleep(6 * time.Second) // wait until lock expires

	expWithdrawTxInfo, err := makerCtrt.ExpWithdraw(acnt0, makerLockTxId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(expWithdrawTxInfo.Id))

	bal, _ := makerCtrt.GetCtrtBal(string(acnt0.Addr.B58Str()))
	require.Equal(t, bal_old.Amount()+makerLockAmount, bal.Amount())
}

func Test_AtomicSwapCtrt_ExpWithdraw(t *testing.T) {
	makerCtrt, err := newAtomicSwap(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	test_AtomicSwapCtrt_ExpWithdraw(t, testAcnt0, testAcnt1, makerCtrt)
}

func Test_AtomicSwapCtrt_AsWhole(t *testing.T) {
	maker_tc, err := newTokCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatal(err)
	}
	makerCtrt := test_AtomicSwapCtrt_Register(t, testAcnt0, maker_tc)
	resp, err := maker_tc.Deposit(testAcnt0, string(makerCtrt.CtrtId.B58Str()), 1000, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	takerCtrt, err := newAtomicSwap(t, testAcnt1)
	if err != nil {
		t.Fatal(err)
	}
	test_AtomicSwapCtrt_Solve(t, testAcnt0, testAcnt1, makerCtrt, takerCtrt)
	test_AtomicSwapCtrt_ExpWithdraw(t, testAcnt0, testAcnt1, makerCtrt)
}
