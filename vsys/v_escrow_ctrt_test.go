package vsys

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func test_VEscrowCtrt_Register(t *testing.T, acnt *Account, tc *TokCtrtWithoutSplit) *VEscrowCtrt {
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	vec, err := RegisterVEscrowCtrt(acnt, string(tokId.B58Str()), 30, 60, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	maker, err := vec.Maker()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, acnt.Addr, maker)
	judge, err := vec.Judge()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, acnt.Addr, judge)
	tokIdFromCtrt, err := vec.TokId()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, tokId, tokIdFromCtrt)

	duration, err := vec.Duration()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, int64(30), duration.UnixTs())

	judge_duration, err := vec.JudgeDuration()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, int64(60), judge_duration.UnixTs())

	unit, _ := tc.Unit()
	vec_unit, _ := vec.Unit()
	require.Equal(t, unit, vec_unit)

	return vec
}

func Test_VEscrowCtrt_Register(t *testing.T) {
	tc, err := newTokCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatalf("Cannot get new token ctrt: %s\n", err.Error())
	}
	test_VEscrowCtrt_Register(t, testAcnt0, tc)
}

func newVEscrowCtrt_forTest(t *testing.T, judge, maker, recipient *Account) *VEscrowCtrt {
	tc, err := newTokCtrtWithTok(t, judge)
	if err != nil {
		t.Fatalf("Cannot get new token ctrt: %s\n", err.Error())
	}
	// testing send is out of scope of this function
	_, err = tc.Send(judge, string(maker.Addr.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Send(judge, string(recipient.Addr.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	ac, err := RegisterVEscrowCtrt(testAcnt0, string(tokId.B58Str()), 30, 60, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	_, err = tc.Deposit(judge, string(ac.CtrtId.B58Str()), 50, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(maker, string(ac.CtrtId.B58Str()), 50, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(recipient, string(ac.CtrtId.B58Str()), 50, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	return ac
}

func Test_VEscrowCtrt_Supersede(t *testing.T) {
	vc := newVEscrowCtrt_forTest(t, testAcnt0, testAcnt1, testAcnt2)

	judge, err := vc.Judge()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, testAcnt0.Addr, judge)

	resp, err := vc.Supersede(testAcnt0, string(testAcnt1.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	judge, err = vc.Judge()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, testAcnt1.Addr, judge)
}

func Test_VEscrowCtrt_Create(t *testing.T) {
	vc := newVEscrowCtrt_forTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + 45
	resp, err := vc.Create(testAcnt1, string(testAcnt2.Addr.B58Str()), 10, 2, 3, 4, 5, later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	orderId := string(resp.Id)

	payer, err := vc.GetOrderPayer(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt1.Addr, payer)
	rcpt, err := vc.GetOrderRecipient(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt2.Addr, rcpt)
	orderAmount, err := vc.GetOrderAmount(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 10.0, orderAmount.Amount())
	rcptDepositAmount, err := vc.GetOrderRecipientDeposit(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 2.0, rcptDepositAmount.Amount())
	judgeDepositAmount, err := vc.GetOrderJudgeDeposit(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 3.0, judgeDepositAmount.Amount())
	orderFee, err := vc.GetOrderFee(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 4.0, orderFee.Amount())
	orderRcptAmount, err := vc.GetOrderRecipientAmount(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 10.0-4.0, orderRcptAmount.Amount())

	totalInOrder := 10.0 + 3.0 + 2.0
	orderRcptRefund, err := vc.GetOrderRecipientRefund(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, totalInOrder-5.0, orderRcptRefund.Amount())
	expTime, _ := vc.GetOrderExpirationTime(orderId)
	require.Equal(t, later, expTime.UnixTs())
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)
	orderRcptDepStatus, _ := vc.GetOrderRecipientDepositStatus(orderId)
	require.Equal(t, false, orderRcptDepStatus)
	orderJudgeDepStatus, _ := vc.GetOrderJudgeDepositStatus(orderId)
	require.Equal(t, false, orderJudgeDepStatus)
	submitStatus, _ := vc.GetOrderSubmitStatus(orderId)
	require.Equal(t, false, submitStatus)
	judgeStatus, _ := vc.GetOrderJudgeStatus(orderId)
	require.Equal(t, false, judgeStatus)
	rcptLockedAmount, _ := vc.GetOrderRecipientLockedAmount(orderId)
	require.Equal(t, 0.0, rcptLockedAmount.Amount())
	judgeLockedAmount, _ := vc.GetOrderRecipientLockedAmount(orderId)
	require.Equal(t, 0.0, judgeLockedAmount.Amount())
}
