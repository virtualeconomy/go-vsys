package vsys

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type vEscrowTest struct {
}

var vecT *vEscrowTest

func (vect *vEscrowTest) DURATION() int64 {
	return 18
}

func (vect *vEscrowTest) ORDER_PERIOD() int64 {
	return 45
}

func (vect *vEscrowTest) newVEscrowCtrt_ForTest(t *testing.T, judge, maker, recipient *Account) *VEscrowCtrt {
	tc, err := asT.newTokCtrtWithTok(t, judge)
	if err != nil {
		t.Fatalf("Cannot get new token ctrt: %s\n", err.Error())
	}
	// testing send is out of scope of this file
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
	vec, err := RegisterVEscrowCtrt(testAcnt0, string(tokId.B58Str()), vect.DURATION(), vect.DURATION(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	_, err = tc.Deposit(judge, string(vec.CtrtId.B58Str()), 500, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(maker, string(vec.CtrtId.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(recipient, string(vec.CtrtId.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	return vec
}

func (vect *vEscrowTest) createOrder(t *testing.T, vc *VEscrowCtrt, payer, recipient *Account, expireAt int64) (orderId string) {
	resp, err := vc.Create(payer, string(recipient.Addr.B58Str()), 10, 2, 3, 4, 5, expireAt, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
	orderId = string(resp.Id)
	return
}

func (vect *vEscrowTest) depositToOrder(t *testing.T, vc *VEscrowCtrt, orderId string, recipient, judge *Account) {
	resp1, err := vc.JudgeDeposit(judge, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	resp2, err := vc.RecipientDeposit(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp1.Id))
	assertTxSuccess(t, string(resp2.Id))
}

func (vect *vEscrowTest) submitWork(t *testing.T, vc *VEscrowCtrt, orderId string, recipient *Account) {
	resp, err := vc.SubmitWork(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))
}

func (vect *vEscrowTest) test_Register(t *testing.T, acnt *Account, tc *TokCtrtWithoutSplit) *VEscrowCtrt {
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	vec, err := RegisterVEscrowCtrt(acnt, string(tokId.B58Str()), vect.DURATION(), vect.DURATION(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	maker, err := vec.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, acnt.Addr, maker)
	judge, err := vec.Judge()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, acnt.Addr, judge)
	tokIdFromCtrt, err := vec.TokId()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, tokId, tokIdFromCtrt)

	duration, err := vec.Duration()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, (&vEscrowTest{}).DURATION(), duration.UnixTs())

	judge_duration, err := vec.JudgeDuration()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, (&vEscrowTest{}).DURATION(), judge_duration.UnixTs())

	unit, _ := tc.Unit()
	vec_unit, _ := vec.Unit()
	require.Equal(t, unit, vec_unit)

	return vec
}

func Test_VEscrowCtrt_Register(t *testing.T) {
	tc, err := asT.newTokCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatalf("Cannot get new token ctrt: %s\n", err.Error())
	}
	vecT.test_Register(t, testAcnt0, tc)
}

func (vect *vEscrowTest) test_Supersede(t *testing.T, vc *VEscrowCtrt, newJudge, oldJudge *Account) {
	judge, err := vc.Judge()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, oldJudge.Addr, judge)

	resp, err := vc.Supersede(oldJudge, string(newJudge.Addr.B58Str()), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	judge, err = vc.Judge()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, newJudge.Addr, judge)
}

func Test_VEscrowCtrt_Supersede(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	vecT.test_Supersede(t, vc, testAcnt1, testAcnt0)
}

func (vect *vEscrowTest) test_Create(t *testing.T, vc *VEscrowCtrt, payer, recipient *Account) string {
	later := time.Now().Unix() + vect.ORDER_PERIOD()
	orderId := vect.createOrder(t, vc, payer, recipient, later)

	payerAddr, err := vc.GetOrderPayer(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, payer.Addr, payerAddr)
	rcpt, err := vc.GetOrderRecipient(orderId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, recipient.Addr, rcpt)
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

	return orderId
}

func Test_VEscrowCtrt_Create(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	vecT.test_Create(t, vc, testAcnt1, testAcnt2)
}

func (vect *vEscrowTest) test_RecipientDeposit(t *testing.T, vc *VEscrowCtrt, orderId string, recipient *Account) {
	orderRcptDepStatus, _ := vc.GetOrderRecipientDepositStatus(orderId)
	require.Equal(t, false, orderRcptDepStatus)
	rcptLockedAmount, _ := vc.GetOrderRecipientLockedAmount(orderId)
	require.Equal(t, 0.0, rcptLockedAmount.Amount())

	resp, err := vc.RecipientDeposit(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderRcptDepStatus, _ = vc.GetOrderRecipientDepositStatus(orderId)
	require.Equal(t, true, orderRcptDepStatus)
	rcptLockedAmount, _ = vc.GetOrderRecipientLockedAmount(orderId)
	require.Equal(t, 2.0, rcptLockedAmount.Amount())
}

func Test_VEscrowCtrt_RecipientDeposit(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.test_RecipientDeposit(t, vc, orderId, testAcnt2)
}

func (vect *vEscrowTest) test_JudgeDeposit(t *testing.T, vc *VEscrowCtrt, orderId string, judge *Account) {
	orderJudgeDepStatus, _ := vc.GetOrderJudgeDepositStatus(orderId)
	require.Equal(t, false, orderJudgeDepStatus)
	judgeLockedAmount, _ := vc.GetOrderJudgeLockedAmount(orderId)
	require.Equal(t, 0.0, judgeLockedAmount.Amount())

	resp, err := vc.JudgeDeposit(judge, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderJudgeDepStatus, _ = vc.GetOrderJudgeDepositStatus(orderId)
	require.Equal(t, true, orderJudgeDepStatus)
	judgeLockedAmount, _ = vc.GetOrderJudgeLockedAmount(orderId)
	require.Equal(t, 3.0, judgeLockedAmount.Amount())
}

func Test_VEscrowCtrt_JudgeDeposit(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.test_JudgeDeposit(t, vc, orderId, testAcnt0)
}

func (vect *vEscrowTest) test_PayerCancel(t *testing.T, vc *VEscrowCtrt, orderId string, payer *Account) {
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.PayerCancel(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}

	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)
}

func Test_VEscrowCtrt_PayerCancel(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.test_PayerCancel(t, vc, orderId, testAcnt1)
}

func (vect *vEscrowTest) test_RecipientCancel(t *testing.T, vc *VEscrowCtrt, orderId string, recipient *Account) {
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.RecipientCancel(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)
}

func Test_VEscrowCtrt_RecipientCancel(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.test_RecipientCancel(t, vc, orderId, testAcnt2)
}

func (vect *vEscrowTest) test_JudgeCancel(t *testing.T, vc *VEscrowCtrt, orderId string, judge *Account) {
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.JudgeCancel(judge, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)
}

func Test_VEscrowCtrt_JudgeCancel(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.test_JudgeCancel(t, vc, orderId, testAcnt0)
}

func (vect *vEscrowTest) test_SubmitWork(t *testing.T, vc *VEscrowCtrt, orderId string, recipient *Account) {
	orderStatus, _ := vc.GetOrderSubmitStatus(orderId)
	require.Equal(t, false, orderStatus)

	resp, err := vc.SubmitWork(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderSubmitStatus(orderId)
	require.Equal(t, true, orderStatus)
}

func Test_VEscrowCtrt_SubmitWork(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.test_SubmitWork(t, vc, orderId, testAcnt2)
}

func (vect *vEscrowTest) test_ApproveWork(t *testing.T, vc *VEscrowCtrt, orderId string, payer, recipient, judge *Account) {
	judgeBalOld, err := vc.GetCtrtBal(string(judge.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	rcptBalOld, err := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.ApproveWork(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	rcptAmt, _ := vc.GetOrderRecipientAmount(orderId)
	fee, _ := vc.GetOrderFee(orderId)
	rcptDep, _ := vc.GetOrderRecipientDeposit(orderId)
	judgeDep, _ := vc.GetOrderJudgeDeposit(orderId)
	rcptBal, _ := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	judgeBal, _ := vc.GetCtrtBal(string(judge.Addr.B58Str()))

	require.Equal(t, rcptBal.Amount()-rcptBalOld.Amount(), rcptAmt.Amount()+rcptDep.Amount())
	require.Equal(t, judgeBal.Amount()-judgeBalOld.Amount(), fee.Amount()+judgeDep.Amount())
}

func Test_VEscrowCtrt_ApproveWork(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)

	vecT.test_ApproveWork(t, vc, orderId, testAcnt1, testAcnt2, testAcnt0)
}

func (vect *vEscrowTest) test_ApplyToAndDoJudge(t *testing.T, vc *VEscrowCtrt, orderId string, payer, recipient, judge *Account) {
	payerBalOld, err := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	judgeBalOld, err := vc.GetCtrtBal(string(judge.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	rcptBalOld, err := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.ApplyToJudge(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	to_payer := 3.0
	to_rcpt := 5.0

	resp, err = vc.DoJudge(judge, orderId, to_payer, to_rcpt, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	fee, _ := vc.GetOrderFee(orderId)
	judgeDep, _ := vc.GetOrderJudgeDeposit(orderId)
	payerBal, _ := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	rcptBal, _ := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	judgeBal, _ := vc.GetCtrtBal(string(judge.Addr.B58Str()))

	require.Equal(t, to_payer, payerBal.Amount()-payerBalOld.Amount())
	require.Equal(t, to_rcpt, rcptBal.Amount()-rcptBalOld.Amount())
	require.Equal(t, fee.Amount()+judgeDep.Amount(), judgeBal.Amount()-judgeBalOld.Amount())
}

func Test_VEscrowCtrt_ApplyToAndDoJudge(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)

	vecT.test_ApplyToAndDoJudge(t, vc, orderId, testAcnt1, testAcnt2, testAcnt0)
}

func (vect *vEscrowTest) test_SubmitPenalty(t *testing.T, vc *VEscrowCtrt, orderId string, payer, judge *Account) {
	payerBalOld, err := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	judgeBalOld, err := vc.GetCtrtBal(string(judge.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	expireAt, err := vc.GetOrderExpirationTime(orderId)
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	// Ensure that the recipient submit work grace period has expired.
	time.Sleep(time.Duration(expireAt.UnixTs()-time.Now().Unix()+6) * time.Second)

	resp, err := vc.SubmitPenalty(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	rcptAmt, _ := vc.GetOrderRecipientAmount(orderId)
	rcptDep, _ := vc.GetOrderRecipientDeposit(orderId)
	fee, _ := vc.GetOrderFee(orderId)
	judgeDep, _ := vc.GetOrderJudgeDeposit(orderId)
	payerBal, _ := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	judgeBal, _ := vc.GetCtrtBal(string(judge.Addr.B58Str()))

	require.Equal(t, rcptAmt.Amount()+rcptDep.Amount(), payerBal.Amount()-payerBalOld.Amount())
	require.Equal(t, fee.Amount()+judgeDep.Amount(), judgeBal.Amount()-judgeBalOld.Amount())
}

func Test_VEscrowCtrt_SubmitPenalty(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + 5
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)

	vecT.test_SubmitPenalty(t, vc, orderId, testAcnt1, testAcnt0)
}

func (vect *vEscrowTest) test_PayerRefund(t *testing.T, vc *VEscrowCtrt, orderId string, payer, recipient *Account) {
	payerBalOld, err := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	rcptBalOld, err := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	expireAt, err := vc.GetOrderExpirationTime(orderId)
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.ApplyToJudge(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	// Ensure that the recipient submit work grace period has expired.
	time.Sleep(time.Duration(expireAt.UnixTs()-time.Now().Unix()+6) * time.Second)

	resp, err = vc.PayerRefund(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	payerRefund, _ := vc.GetOrderRefund(orderId)
	rcptRefund, _ := vc.GetOrderRecipientRefund(orderId)
	payerBal, _ := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	rcptBal, _ := vc.GetCtrtBal(string(recipient.Addr.B58Str()))

	require.Equal(t, payerRefund.Amount(), payerBal.Amount()-payerBalOld.Amount())
	require.Equal(t, rcptRefund.Amount(), rcptBal.Amount()-rcptBalOld.Amount())
}

func Test_VEscrowCtrt_PayerRefund(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)

	vecT.test_PayerRefund(t, vc, orderId, testAcnt1, testAcnt2)
}

func (vect *vEscrowTest) test_RecipientRefund(t *testing.T, vc *VEscrowCtrt, orderId string, payer, recipient *Account) {
	payerBalOld, err := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	rcptBalOld, err := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	expireAt, err := vc.GetOrderExpirationTime(orderId)
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	resp, err := vc.ApplyToJudge(payer, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	// Ensure that the recipient submit work grace period has expired.
	time.Sleep(time.Duration(expireAt.UnixTs()-time.Now().Unix()+6) * time.Second)

	resp, err = vc.RecipientRefund(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	payerRefund, _ := vc.GetOrderRefund(orderId)
	rcptRefund, _ := vc.GetOrderRecipientRefund(orderId)
	payerBal, _ := vc.GetCtrtBal(string(payer.Addr.B58Str()))
	rcptBal, _ := vc.GetCtrtBal(string(recipient.Addr.B58Str()))

	require.Equal(t, payerRefund.Amount(), payerBal.Amount()-payerBalOld.Amount())
	require.Equal(t, rcptRefund.Amount(), rcptBal.Amount()-rcptBalOld.Amount())
}

func Test_VEscrowCtrt_RecipientRefund(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)

	vecT.test_RecipientRefund(t, vc, orderId, testAcnt1, testAcnt2)
}

func (vect *vEscrowTest) test_Collect(t *testing.T, vc *VEscrowCtrt, orderId string, recipient *Account, judge *Account) {
	judgeBalOld, err := vc.GetCtrtBal(string(judge.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	rcptBalOld, err := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	if err != nil {
		t.Fatal(err)
	}
	expireAt, err := vc.GetOrderExpirationTime(orderId)
	if err != nil {
		t.Fatal(err)
	}
	orderStatus, _ := vc.GetOrderStatus(orderId)
	require.Equal(t, true, orderStatus)

	// Ensure that the recipient submit work grace period has expired.
	time.Sleep(time.Duration(expireAt.UnixTs()-time.Now().Unix()+6) * time.Second)

	resp, err := vc.Collect(recipient, orderId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, string(resp.Id))

	orderStatus, _ = vc.GetOrderStatus(orderId)
	require.Equal(t, false, orderStatus)

	rcptBal, _ := vc.GetCtrtBal(string(recipient.Addr.B58Str()))
	rcptAmt, _ := vc.GetOrderRecipientAmount(orderId)
	rcptDep, _ := vc.GetOrderRecipientDeposit(orderId)
	judgeBal, _ := vc.GetCtrtBal(string(judge.Addr.B58Str()))
	fee, _ := vc.GetOrderFee(orderId)
	judgeDep, _ := vc.GetOrderJudgeDeposit(orderId)

	require.Equal(t, rcptAmt.Amount()+rcptDep.Amount(), rcptBal.Amount()-rcptBalOld.Amount())
	require.Equal(t, fee.Amount()+judgeDep.Amount(), judgeBal.Amount()-judgeBalOld.Amount())
}

func Test_VEscrowCtrt_Collect(t *testing.T) {
	vc := vecT.newVEscrowCtrt_ForTest(t, testAcnt0, testAcnt1, testAcnt2)

	later := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId := vecT.createOrder(t, vc, testAcnt1, testAcnt2, later)

	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)
	vecT.test_Collect(t, vc, orderId, testAcnt2, testAcnt0)
}

func Test_VEscrowCtrt_AsWhole(t *testing.T) {
	tc, err := asT.newTokCtrtWithTok(t, testAcnt0)
	if err != nil {
		t.Fatalf("Cannot get new token ctrt: %s\n", err.Error())
	}
	_, err = tc.Send(testAcnt0, string(testAcnt1.Addr.B58Str()), 500, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Send(testAcnt0, string(testAcnt2.Addr.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	vc := vecT.test_Register(t, testAcnt0, tc)

	_, err = tc.Deposit(testAcnt0, string(vc.CtrtId.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(testAcnt1, string(vc.CtrtId.B58Str()), 500, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tc.Deposit(testAcnt2, string(vc.CtrtId.B58Str()), 200, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	orderId := vecT.test_Create(t, vc, testAcnt1, testAcnt2)
	vecT.test_RecipientDeposit(t, vc, orderId, testAcnt2)
	vecT.test_JudgeDeposit(t, vc, orderId, testAcnt0)
	vecT.test_SubmitWork(t, vc, orderId, testAcnt2)
	vecT.test_ApproveWork(t, vc, orderId, testAcnt1, testAcnt2, testAcnt0)

	// Test cancel
	expire_at := time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.test_PayerCancel(t, vc, orderId, testAcnt1)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.test_RecipientCancel(t, vc, orderId, testAcnt2)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.test_JudgeCancel(t, vc, orderId, testAcnt0)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)
	vecT.test_ApplyToAndDoJudge(t, vc, orderId, testAcnt1, testAcnt2, testAcnt0)

	expire_at = time.Now().Unix() + 10
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.test_SubmitPenalty(t, vc, orderId, testAcnt1, testAcnt0)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)
	vecT.test_PayerRefund(t, vc, orderId, testAcnt1, testAcnt2)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)
	vecT.test_RecipientRefund(t, vc, orderId, testAcnt1, testAcnt2)

	expire_at = time.Now().Unix() + vecT.ORDER_PERIOD()
	orderId = vecT.createOrder(t, vc, testAcnt1, testAcnt2, expire_at)
	vecT.depositToOrder(t, vc, orderId, testAcnt2, testAcnt0)
	vecT.submitWork(t, vc, orderId, testAcnt2)
	vecT.test_Collect(t, vc, orderId, testAcnt2, testAcnt0)

	vecT.test_Supersede(t, vc, testAcnt1, testAcnt0)
}
