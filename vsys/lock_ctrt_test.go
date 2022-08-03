package vsys

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type lockCtrtTest struct{}

var lcT *lockCtrtTest

func (l *lockCtrtTest) TOK_MAX() float64 {
	return 100
}

func (l *lockCtrtTest) TOK_UNIT() uint64 {
	return 1
}

func (l *lockCtrtTest) newTokCtrt(t *testing.T) *TokCtrtWithoutSplit {
	tc, err := RegisterTokCtrtWithoutSplit(testAcnt0, l.TOK_MAX(), l.TOK_UNIT(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	_, err = tc.Issue(testAcnt0, l.TOK_MAX(), "")
	waitForBlock()

	return tc
}

func (l *lockCtrtTest) newCtrt(t *testing.T, tc *TokCtrtWithoutSplit) *LockCtrt {
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}

	lc, err := RegisterLockCtrt(testAcnt0, tokId.B58Str().Str(), "")
	waitForBlock()
	return lc
}

func (l *lockCtrtTest) test_Register(t *testing.T, tc *TokCtrtWithoutSplit, lc *LockCtrt) {
	maker, err := lc.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt0.Addr, maker)

	lcTokId, err := lc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	tcTokId, err := tc.TokId()
	require.Equal(t, tcTokId, lcTokId)

	bal, err := lc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0.0, bal.Amount())

	lockTime, err := lc.GetCtrtLockTime(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(0), lockTime.UnixTs())
}

func Test_LockCtrt_Register(t *testing.T) {
	tc := lcT.newTokCtrt(t)
	lc := lcT.newCtrt(t, tc)

	lcT.test_Register(t, tc, lc)
}

func (l *lockCtrtTest) test_Lock(t *testing.T, tc *TokCtrtWithoutSplit, lc *LockCtrt) {
	resp, err := tc.Deposit(testAcnt0, lc.CtrtId.B58Str().Str(), lcT.TOK_MAX(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	bal, err := lc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, lcT.TOK_MAX(), bal.Amount())

	later := time.Now().Unix() + 3*int64(avgBlockDelay.Seconds())
	resp, err = lc.Lock(testAcnt0, later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	lockTime, err := lc.GetCtrtLockTime(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(later), lockTime.UnixTs())

	// withdraw before expiration should fail
	resp, _ = tc.Withdraw(testAcnt0, lc.CtrtId.B58Str().Str(), lcT.TOK_MAX(), "")
	waitForBlock()
	assertTxStatus(t, string(resp.Id), "Failed")
	bal, err = lc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, lcT.TOK_MAX(), bal.Amount())

	time.Sleep(time.Second*time.Duration(later-time.Now().Unix()) + avgBlockDelay)

	// withdraw after expiration should succeed
	resp, _ = tc.Withdraw(testAcnt0, lc.CtrtId.B58Str().Str(), lcT.TOK_MAX(), "")
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	bal, err = lc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0.0, bal.Amount())
}

func Test_LockCtrt_Lock(t *testing.T) {
	tc := lcT.newTokCtrt(t)
	lc := lcT.newCtrt(t, tc)

	lcT.test_Lock(t, tc, lc)

}

func Test_LockCtrt_AsWhole(t *testing.T) {
	tc := lcT.newTokCtrt(t)
	lc := lcT.newCtrt(t, tc)

	lcT.test_Register(t, tc, lc)
	lcT.test_Lock(t, tc, lc)
}
