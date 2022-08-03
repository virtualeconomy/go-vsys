package vsys

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type vStableSwapTest struct {
}

var vssT *vStableSwapTest

func (v *vStableSwapTest) test_Register(t *testing.T, acnt *Account, newctrt *VStableSwapCtrt) *VStableSwapCtrt {
	maker, err := newctrt.Maker()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, acnt.Addr, maker)
	return newctrt
}

func (v *vStableSwapTest) newCtrtWithTok(by *Account) (*TokCtrtWithoutSplit, error) {
	tc, err := RegisterTokCtrtWithoutSplit(by, 1000, 1, "", "")
	if err != nil {
		return nil, err
	}
	waitForBlock()

	_, err = tc.Issue(by, 1000, "")
	if err != nil {
		return nil, err
	}
	waitForBlock()
	return tc, nil
}

func (v *vStableSwapTest) newStableCtrt(t *testing.T, by *Account) *VStableSwapCtrt {
	baseTc, err := vssT.newCtrtWithTok(by)
	if err != nil {
		t.Fatal(err)
	}
	targetTc, err := vssT.newCtrtWithTok(by)
	if err != nil {
		t.Fatal(err)
	}

	baseTokId, _ := baseTc.TokId()
	targetTokId, _ := targetTc.TokId()

	ac, err := RegisterVStableSwapCtrt(by, baseTokId.B58Str().Str(), targetTokId.B58Str().Str(), 5, 1, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	_, err = baseTc.Deposit(by, ac.CtrtId.B58Str().Str(), 1000, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = targetTc.Deposit(by, ac.CtrtId.B58Str().Str(), 1000, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	return ac
}

func Test_VStableSwapCtrt_Register(t *testing.T) {
	vss := vssT.newStableCtrt(t, testAcnt0)
	vssT.test_Register(t, testAcnt0, vss)
}

func (v *vStableSwapTest) test_SetAndUpdateOrder(t *testing.T, vss *VStableSwapCtrt) string {
	resp, err := vss.SetOrder(testAcnt0, 1, 1, 0, 100, 0, 100, 2, 1, 500, 500, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	order_id := string(resp.Id)
	assertTxSuccess(t, order_id)

	baseTokBal, err := vss.GetBaseTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	targetTokBal, err := vss.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	priceBase1, err := vss.GetPriceBase(order_id)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, 500.0, baseTokBal.Amount())
	require.Equal(t, 500.0, targetTokBal.Amount())
	require.Equal(t, 2.0, priceBase1.Amount())

	resp, err = vss.UpdateOrder(testAcnt0, order_id, 1, 1, 0, 100, 0, 100, 1, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	priceBase2, err := vss.GetPriceBase(order_id)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 1.0, priceBase2.Amount())
	return order_id
}

func Test_VStableSwapCtrt_SetAndUpdateOrder(t *testing.T) {
	vss := vssT.newStableCtrt(t, testAcnt0)

	vssT.test_SetAndUpdateOrder(t, vss)
}

func (v *vStableSwapTest) newStableSwapCtrtWithOrder(t *testing.T) (*VStableSwapCtrt, string) {
	vss := vssT.newStableCtrt(t, testAcnt0)
	resp, err := vss.SetOrder(testAcnt0, 1, 1, 0, 100, 0, 100, 1, 1, 500, 500, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	return vss, string(resp.Id)
}

func (v *vStableSwapTest) test_DepositAndWithdraw(t *testing.T, vss *VStableSwapCtrt, orderId string) {
	resp, err := vss.OrderDeposit(testAcnt0, orderId, 200, 100, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	baseTokBal, err := vss.GetBaseTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	targetTokBal, err := vss.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 300.0, baseTokBal.Amount())
	require.Equal(t, 400.0, targetTokBal.Amount())

	resp, err = vss.OrderWithdraw(testAcnt0, orderId, 200, 100, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	baseTokBal, err = vss.GetBaseTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	targetTokBal, err = vss.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 500.0, baseTokBal.Amount())
	require.Equal(t, 500.0, targetTokBal.Amount())

}

func Test_VStableSwapCtrt_DepositAndWithdraw(t *testing.T) {
	vss, order_id := vssT.newStableSwapCtrtWithOrder(t)
	vssT.test_DepositAndWithdraw(t, vss, order_id)
}

func (v *vStableSwapTest) test_Swap(t *testing.T, vss *VStableSwapCtrt, orderId string) {
	deadline := time.Now().Unix() + 1500
	resp, err := vss.SwapBaseToTarget(testAcnt0, orderId, 10, 1, 1, deadline, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	baseTokBal, err := vss.GetBaseTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	targetTokBal, err := vss.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 490.0, baseTokBal.Amount())
	require.Equal(t, 509.0, targetTokBal.Amount())

	resp, err = vss.SwapTargetToBase(testAcnt0, orderId, 10, 1, 1, deadline, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	baseTokBal, err = vss.GetBaseTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	targetTokBal, err = vss.GetTargetTokBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 499.0, baseTokBal.Amount())
	require.Equal(t, 499.0, targetTokBal.Amount())
}

func Test_VStableSwapCtrt_Swap(t *testing.T) {
	vss, order_id := vssT.newStableSwapCtrtWithOrder(t)

	vssT.test_Swap(t, vss, order_id)
}

func (v *vStableSwapTest) test_CloseOrder(t *testing.T, vss *VStableSwapCtrt, orderId string) {
	status, err := vss.GetOrderStatus(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, status)

	resp, err := vss.CloseOrder(testAcnt0, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	status, err = vss.GetOrderStatus(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.False(t, status)
}

func Test_VStableSwapCtrt_CloseOrder(t *testing.T) {
	vss, order_id := vssT.newStableSwapCtrtWithOrder(t)
	vssT.test_CloseOrder(t, vss, order_id)
}

func Test_VStableSwapCtrt_AsWhole(t *testing.T) {
	vss := vssT.newStableCtrt(t, testAcnt0)
	vss = vssT.test_Register(t, testAcnt0, vss)
	orderId := vssT.test_SetAndUpdateOrder(t, vss)

	vssT.test_DepositAndWithdraw(t, vss, orderId)
	vssT.test_Swap(t, vss, orderId)
	vssT.test_CloseOrder(t, vss, orderId)
}
