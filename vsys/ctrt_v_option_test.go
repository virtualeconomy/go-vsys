package vsys

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

type vOptionTest struct {
}

func (vot *vOptionTest) MAX_ISSUE_AMOUNT() float64 {
	return 1000
}

func (vot *vOptionTest) MINT_AMOUNT() float64 {
	return 200
}

func (vot *vOptionTest) UNLOCK_AMOUNT() float64 {
	return 100
}

func (vot *vOptionTest) EXEC_TIME_DELTA() int64 {
	return 50
}

func (vot *vOptionTest) EXEC_DDL_DELTA() int64 {
	return 95
}

var voT *vOptionTest

func (vot *vOptionTest) newTokCtrt(t *testing.T, mu *sync.Mutex) *TokCtrtWithoutSplit {
	mu.Lock()
	tc, err := RegisterTokCtrtWithoutSplit(testAcnt0, 1000, 1, "", "")
	if err != nil {
		t.Fatal(err)
	}
	mu.Unlock()
	waitForBlock()
	mu.Lock()
	resp, err := tc.Issue(testAcnt0, 1000, "")
	if err != nil {
		t.Fatal(err)
	}
	mu.Unlock()
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return tc
}

func (vot *vOptionTest) newVOptionCtrt(t *testing.T) *VOptionCtrt {
	g := new(errgroup.Group)
	var mu sync.Mutex
	var (
		baseTc, targetTc, optionTc, proofTc *TokCtrtWithoutSplit
	)
	g.Go(func() (err error) {
		baseTc = vot.newTokCtrt(t, &mu)
		return
	})
	g.Go(func() (err error) {
		targetTc = vot.newTokCtrt(t, &mu)
		return
	})
	g.Go(func() (err error) {
		optionTc = vot.newTokCtrt(t, &mu)
		return
	})
	g.Go(func() (err error) {
		proofTc = vot.newTokCtrt(t, &mu)
		return
	})
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}

	baseTokId, err := baseTc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	targetTokId, err := targetTc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	optionTokId, err := optionTc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	proofTokId, err := proofTc.TokId()
	if err != nil {
		t.Fatal(err)
	}

	vo, err := RegisterVOptionCtrt(
		testAcnt0,
		baseTokId.B58Str().Str(),
		targetTokId.B58Str().Str(),
		optionTokId.B58Str().Str(),
		proofTokId.B58Str().Str(),
		time.Now().Unix()+vot.EXEC_TIME_DELTA(),
		time.Now().Unix()+vot.EXEC_DDL_DELTA(),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	baseTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")   //nolint:errcheck
	targetTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "") //nolint:errcheck
	optionTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "") //nolint:errcheck
	proofTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")  //nolint:errcheck
	waitForBlock()

	return vo
}

func (vot *vOptionTest) newVOptionCtrtActivated(t *testing.T) *VOptionCtrt {
	vo := vot.newVOptionCtrt(t)
	resp, err := vo.Activate(testAcnt0, vot.MAX_ISSUE_AMOUNT(), 10, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return vo
}

func (vot *vOptionTest) newVOptionCtrtActivatedAndMinted(t *testing.T) *VOptionCtrt {
	vo := vot.newVOptionCtrtActivated(t)

	resp, err := vo.Mint(testAcnt0, vot.MINT_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return vo
}

func (vot *vOptionTest) test_Register(t *testing.T, vo *VOptionCtrt) {
	maker, err := vo.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt0.Addr, maker)
}

func Test_VOptionCtrt_Register(t *testing.T) {
	vo := voT.newVOptionCtrt(t)
	voT.test_Register(t, vo)
}

func (vot *vOptionTest) test_Activate(t *testing.T, vo *VOptionCtrt) {
	maxIssueNum, err := vo.MaxIssueNum()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, vot.MAX_ISSUE_AMOUNT(), maxIssueNum.Amount())
}

func Test_VOptionCtrt_Activate(t *testing.T) {
	vo := voT.newVOptionCtrtActivated(t)
	voT.test_Activate(t, vo)
}

func (vot *vOptionTest) test_Mint(t *testing.T, vo *VOptionCtrt) {
	targetTokBal, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, vot.MAX_ISSUE_AMOUNT()-vot.MINT_AMOUNT(), targetTokBal.Amount())
}

func Test_VOptionCtrt_Mint(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Mint(t, vo)
}

func (vot *vOptionTest) test_Unlock(t *testing.T, vo *VOptionCtrt) {
	resp, err := vo.Unlock(testAcnt0, vot.UNLOCK_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	targetTokBal, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, vot.MAX_ISSUE_AMOUNT()-vot.MINT_AMOUNT()+vot.UNLOCK_AMOUNT(), targetTokBal.Amount())
}

func Test_VOptionCtrt_Unlock(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Unlock(t, vo)
}

func (vot *vOptionTest) test_ExecuteAndCollect(t *testing.T, vo *VOptionCtrt) {
	exec_amount := 10.0
	targetTokBalOld, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(avgBlockDelay * 6)

	exeTx, err := vo.Execute(testAcnt0, exec_amount, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, exeTx.Id.Str())

	targetTokBalExec, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, exec_amount, targetTokBalExec.Amount()-targetTokBalOld.Amount())

	time.Sleep(avgBlockDelay * 5)
	colTx, err := vo.Collect(testAcnt0, 10.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, colTx.Id.Str())

	targetTokBalCol, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 9.0, targetTokBalCol.Amount()-targetTokBalExec.Amount())
}

func Test_VOptionCtrt_ExecuteAndCollect(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_ExecuteAndCollect(t, vo)
}

func (vot *vOptionTest) test_Supersede(t *testing.T, vo *VOptionCtrt) {
	resp, err := vo.Supersede(testAcnt0, testAcnt1.Addr.B58Str().Str(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	maker, err := vo.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt1.Addr, maker)
}

func Test_VOptionCtrt_Supersede(t *testing.T) {
	vo := voT.newVOptionCtrt(t)
	voT.test_Supersede(t, vo)
}

func Test_VOptionCtrt_AsWhole(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Register(t, vo)
	voT.test_Activate(t, vo)
	voT.test_Mint(t, vo)
	voT.test_Unlock(t, vo)
	voT.test_ExecuteAndCollect(t, vo)
	voT.test_Supersede(t, vo)
}
