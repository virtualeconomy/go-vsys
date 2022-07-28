package vsys

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type vSwapTest struct {
}

func (v *vSwapTest) TOK_MAX() float64 {
	return 1_000_000_000
}
func (v *vSwapTest) TOK_UNIT() uint64 {
	return 1_000
}
func (v *vSwapTest) MIN_LIQ() int {
	return 10
}
func (v *vSwapTest) INIT_AMOUNT() float64 {
	return 10_000
}

var vsT *vSwapTest

func (v *vSwapTest) newCtrt(t *testing.T, acnt0, acnt1 *Account) *VSwapCtrt {
	tca, tcb, tcl := v.newTokCtrts(t, acnt0)

	resp1, err := tca.Send(acnt0, string(acnt1.Addr.B58Str()), v.TOK_MAX()/2, "")
	if err != nil {
		t.Fatal(err)
	}
	resp2, err := tcb.Send(acnt0, string(acnt1.Addr.B58Str()), v.TOK_MAX()/2, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp1.Id))
	assertTxSuccess(t, string(resp2.Id))

	tokIdA, _ := tca.TokId()
	tokIdB, _ := tcb.TokId()
	tokIdL, _ := tcl.TokId()

	vs, err := RegisterVSwapCtrt(acnt0, tokIdA.B58Str().Str(), tokIdB.B58Str().Str(), tokIdL.B58Str().Str(), v.MIN_LIQ(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	respa0, _ := tca.Deposit(acnt0, string(vs.CtrtId.B58Str()), v.TOK_MAX()/2, "")
	respb0, _ := tcb.Deposit(acnt0, string(vs.CtrtId.B58Str()), v.TOK_MAX()/2, "")
	respa1, _ := tca.Deposit(acnt1, string(vs.CtrtId.B58Str()), v.TOK_MAX()/2, "")
	respb1, _ := tcb.Deposit(acnt1, string(vs.CtrtId.B58Str()), v.TOK_MAX()/2, "")
	respl, _ := tcl.Deposit(acnt0, string(vs.CtrtId.B58Str()), v.TOK_MAX(), "")
	waitForBlock()
	assertTxSuccess(t, respa0.Id.Str())
	assertTxSuccess(t, respb0.Id.Str())
	assertTxSuccess(t, respa1.Id.Str())
	assertTxSuccess(t, respb1.Id.Str())
	assertTxSuccess(t, respl.Id.Str())

	return vs
}

func (v *vSwapTest) newTokCtrts(t *testing.T, acnt *Account) (a, b, c *TokCtrtWithoutSplit) {
	var toks [3]*TokCtrtWithoutSplit
	var err error
	toks[0], err = RegisterTokCtrtWithoutSplit(acnt, v.TOK_MAX(), v.TOK_UNIT(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	toks[1], err = RegisterTokCtrtWithoutSplit(acnt, v.TOK_MAX(), v.TOK_UNIT(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	toks[2], err = RegisterTokCtrtWithoutSplit(acnt, v.TOK_MAX(), v.TOK_UNIT(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	var resps [3]*BroadcastExecuteTxResp
	var errs [3]error
	for i, tc := range toks {
		resps[i], errs[i] = tc.Issue(acnt, v.TOK_MAX(), "")
	}
	for i, err := range errs {
		if err != nil {
			t.Fatalf("Token %d issue failed: %s", i, err.Error())
		}
	}
	waitForBlock()
	for _, r := range resps {
		assertTxSuccess(t, string(r.Id))
	}
	return toks[0], toks[1], toks[2]
}

func (v *vSwapTest) newCtrtWithPool(t *testing.T) *VSwapCtrt {
	vs := vsT.newCtrt(t, testAcnt0, testAcnt1)

	resp, err := vs.SetSwap(testAcnt0, vsT.INIT_AMOUNT(), vsT.INIT_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	status, err := vs.IsSwapActive()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, status)

	return vs
}

func (v *vSwapTest) test_Supersede(t *testing.T, vs *VSwapCtrt) {
	maker, err := vs.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt0.Addr, maker)

	resp, err := vs.Supersede(testAcnt0, string(testAcnt1.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	maker, err = vs.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt1.Addr, maker)
}
func Test_VSwapCtrt_Supersede(t *testing.T) {
	vs := vsT.newCtrt(t, testAcnt0, testAcnt1)

	vsT.test_Supersede(t, vs)
}

func (v *vSwapTest) test_SetSwap(t *testing.T, vs *VSwapCtrt) {
	status, err := vs.IsSwapActive()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, false, status)

	resp, err := vs.SetSwap(testAcnt0, vsT.INIT_AMOUNT(), vsT.INIT_AMOUNT(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	status, err = vs.IsSwapActive()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, true, status)
}

func Test_VSwapCtrt_SetSwap(t *testing.T) {
	vs := vsT.newCtrt(t, testAcnt0, testAcnt1)

	vsT.test_SetSwap(t, vs)
}

func (v *vSwapTest) test_AddLiquidity(t *testing.T, vs *VSwapCtrt) {
	const (
		DELTA     = 10_000
		DELTA_MIN = 9_000
	)

	tokAReservedOld, err := vs.TokAReserved()
	if err != nil {
		t.Fatal(err)
	}
	tokBReservedOld, err := vs.TokBReserved()
	if err != nil {
		t.Fatal(err)
	}
	liqTokLeftOld, err := vs.LiqTokLeft()
	if err != nil {
		t.Fatal(err)
	}
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.AddLiquidity(testAcnt0,
		DELTA,
		DELTA,
		DELTA_MIN,
		DELTA_MIN,
		ten_sec_later,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokAReserved, err := vs.TokAReserved()
	if err != nil {
		t.Fatal(err)
	}
	tokBReserved, err := vs.TokBReserved()
	if err != nil {
		t.Fatal(err)
	}
	liqTokLeft, err := vs.LiqTokLeft()
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, tokAReservedOld.Amount()+DELTA, tokAReserved.Amount())
	require.Equal(t, tokBReservedOld.Amount()+DELTA, tokBReserved.Amount())
	require.Equal(t, liqTokLeftOld.Amount()-DELTA, liqTokLeft.Amount())
}

func Test_VSwapCtrt_AddLiquidity(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_AddLiquidity(t, vs)
}

func (v *vSwapTest) test_RemoveLiquidity(t *testing.T, vs *VSwapCtrt) {
	const (
		DELTA = 1_000
	)

	tokAReservedOld, err := vs.TokAReserved()
	if err != nil {
		t.Fatal(err)
	}
	tokBReservedOld, err := vs.TokBReserved()
	if err != nil {
		t.Fatal(err)
	}
	liqTokLeftOld, err := vs.LiqTokLeft()
	if err != nil {
		t.Fatal(err)
	}
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.RemoveLiquidity(
		testAcnt0,
		DELTA,
		DELTA,
		DELTA,
		ten_sec_later,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	tokAReserved, err := vs.TokAReserved()
	if err != nil {
		t.Fatal(err)
	}
	tokBReserved, err := vs.TokBReserved()
	if err != nil {
		t.Fatal(err)
	}
	liqTokLeft, err := vs.LiqTokLeft()
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, liqTokLeftOld.Amount()+DELTA, liqTokLeft.Amount())
	require.GreaterOrEqual(t, tokAReservedOld.Amount()-tokAReserved.Amount(), float64(DELTA))
	require.GreaterOrEqual(t, tokBReservedOld.Amount()-tokBReserved.Amount(), float64(DELTA))
}

func Test_VSwapCtrt_RemoveLiquidity(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_RemoveLiquidity(t, vs)
}

func (v *vSwapTest) test_SwapBForExactA(t *testing.T, vs *VSwapCtrt) {
	balAOld, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balBOld, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	amountA := 10.0
	amountBMax := 20.0
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.SwapBForExactA(testAcnt1, amountA, amountBMax, ten_sec_later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	balA, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balB, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, balAOld.Amount()+amountA, balA.Amount())
	require.GreaterOrEqual(t, amountBMax, balBOld.Amount()-balB.Amount())
}

func Test_VSwapCtrt_SwapBForExactA(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_SwapBForExactA(t, vs)
}

func (v *vSwapTest) test_SwapExactBForA(t *testing.T, vs *VSwapCtrt) {
	balAOld, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balBOld, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	amountAmin := 10.0
	amountB := 20.0
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.SwapExactBForA(testAcnt1, amountAmin, amountB, ten_sec_later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	balA, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balB, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	require.LessOrEqual(t, amountAmin, balA.Amount()-balAOld.Amount())
	require.Equal(t, balBOld.Amount()-amountB, balB.Amount())
}

func Test_VSwapCtrt_SwapExactBForA(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_SwapExactBForA(t, vs)
}

func (v *vSwapTest) test_SwapAForExactB(t *testing.T, vs *VSwapCtrt) {
	balAOld, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balBOld, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	amountB := 10.0
	amountAMax := 20.0
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.SwapAForExactB(testAcnt1, amountB, amountAMax, ten_sec_later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	balA, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balB, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, balBOld.Amount()+amountB, balB.Amount())
	require.GreaterOrEqual(t, amountAMax, balAOld.Amount()-balA.Amount())
}

func Test_VSwapCtrt_SwapAForExactB(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_SwapAForExactB(t, vs)
}

func (v *vSwapTest) test_SwapExactAForB(t *testing.T, vs *VSwapCtrt) {
	balAOld, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balBOld, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	amountBmin := 10.0
	amountA := 20.0
	ten_sec_later := time.Now().Unix() + 10

	resp, err := vs.SwapExactAForB(testAcnt1, amountBmin, amountA, ten_sec_later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	balA, err := vs.GetTokABal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	balB, err := vs.GetTokBBal(string(testAcnt1.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, balAOld.Amount()-amountA, balA.Amount())
	require.LessOrEqual(t, amountBmin, balB.Amount()-balBOld.Amount())

}

func Test_VSwapCtrt_SwapExactAForB(t *testing.T) {
	vs := vsT.newCtrtWithPool(t)

	vsT.test_SwapExactAForB(t, vs)
}

func Test_VSwapCtrt_AsWhole(t *testing.T) {
	vs := vsT.newCtrt(t, testAcnt0, testAcnt1)
	vsT.test_SetSwap(t, vs)
	vsT.test_AddLiquidity(t, vs)
	vsT.test_RemoveLiquidity(t, vs)
	vsT.test_SwapBForExactA(t, vs)
	vsT.test_SwapExactBForA(t, vs)
	vsT.test_SwapAForExactB(t, vs)
	vsT.test_SwapExactAForB(t, vs)
	vsT.test_Supersede(t, vs)
}
