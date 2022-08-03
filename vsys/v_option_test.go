package vsys

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

type vOptionTest struct {
}

func (v *vOptionTest) MAX_ISSUE_AMOUNT() float64 {
	return 1000
}

func (v *vOptionTest) MINT_AMOUNT() float64 {
	return 200
}

func (v *vOptionTest) UNLOCK_AMOUNT() float64 {
	return 100
}

func (v *vOptionTest) EXEC_TIME_DELTA() int64 {
	return 50
}

func (v *vOptionTest) EXEC_DDL_DELTA() int64 {
	return 95
}

var voT *vOptionTest

func (v *vOptionTest) newTokCtrt(t *testing.T) (*TokCtrtWithoutSplit, error) {
	tc, err := RegisterTokCtrtWithoutSplit(testAcnt0, 1000, 1, "", "")
	if err != nil {
		return nil, fmt.Errorf("newTokCtrt: %w", err)
	}
	waitForBlock()
	resp, err := tc.Issue(testAcnt0, 1000, "")
	if err != nil {
		return nil, fmt.Errorf("newTokCtrt: %w", err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return tc, nil
}

func (v *vOptionTest) newVOptionCtrt(t *testing.T) *VOptionCtrt {
	g := new(errgroup.Group)
	var (
		baseTc, targetTc, optionTc, proofTc *TokCtrtWithoutSplit
	)
	g.Go(func() (err error) {
		baseTc, err = v.newTokCtrt(t)
		return
	})
	g.Go(func() (err error) {
		targetTc, err = v.newTokCtrt(t)
		return
	})
	g.Go(func() (err error) {
		optionTc, err = v.newTokCtrt(t)
		return
	})
	g.Go(func() (err error) {
		proofTc, err = v.newTokCtrt(t)
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
		time.Now().Unix()+v.EXEC_TIME_DELTA(),
		time.Now().Unix()+v.EXEC_DDL_DELTA(),
		"",
	)
	waitForBlock()

	baseTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")
	targetTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")
	optionTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")
	proofTc.Deposit(testAcnt0, vo.CtrtId.B58Str().Str(), 1000, "")
	waitForBlock()

	return vo
}

func (v *vOptionTest) newVOptionCtrtActivated(t *testing.T) *VOptionCtrt {
	vo := v.newVOptionCtrt(t)
	resp, err := vo.Activate(testAcnt0, v.MAX_ISSUE_AMOUNT(), 10, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return vo
}

func (v *vOptionTest) newVOptionCtrtActivatedAndMinted(t *testing.T) *VOptionCtrt {
	vo := voT.newVOptionCtrtActivated(t)

	resp, err := vo.Mint(testAcnt0, v.MINT_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())
	return vo
}

func (v *vOptionTest) test_Register(t *testing.T, vo *VOptionCtrt) {
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

func (v *vOptionTest) test_Activate(t *testing.T, vo *VOptionCtrt) {
	maxIssueNum, err := vo.MaxIssueNum()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, voT.MAX_ISSUE_AMOUNT(), maxIssueNum.Amount())
}

func Test_VOptionCtrt_Activate(t *testing.T) {
	vo := voT.newVOptionCtrtActivated(t)
	voT.test_Activate(t, vo)
}

func (v *vOptionTest) test_Mint(t *testing.T, vo *VOptionCtrt) {
	targetTokBal, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, voT.MAX_ISSUE_AMOUNT()-voT.MINT_AMOUNT(), targetTokBal.Amount())
}

func Test_VOptionCtrt_Mint(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Mint(t, vo)
}

func (v *vOptionTest) test_Unlock(t *testing.T, vo *VOptionCtrt) {
	resp, err := vo.Unlock(testAcnt0, voT.UNLOCK_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	targetTokBal, err := vo.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, voT.MAX_ISSUE_AMOUNT()-voT.MINT_AMOUNT()+voT.UNLOCK_AMOUNT(), targetTokBal.Amount())
}

func Test_VOptionCtrt_Unlock(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Unlock(t, vo)
}

func (v *vOptionTest) test_ExecuteAndCollect(t *testing.T, vo *VOptionCtrt) {
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

func Test_VOptionCtrt_AsWhole(t *testing.T) {
	vo := voT.newVOptionCtrtActivatedAndMinted(t)
	voT.test_Register(t, vo)
	voT.test_Activate(t, vo)
	voT.test_Mint(t, vo)
	voT.test_Unlock(t, vo)
	voT.test_ExecuteAndCollect(t, vo)
}
