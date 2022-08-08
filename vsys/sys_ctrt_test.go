package vsys

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type sysCtrtTest struct {
}

var scT *sysCtrtTest

func (s *sysCtrtTest) newCtrt(t *testing.T) *SysCtrt {
	return NewSysCtrtForTestnet(testAcnt0.Chain)
}

func (s *sysCtrtTest) newPayChanCtrt(t *testing.T, sc *SysCtrt) *PayChanCtrt {
	tokId, err := sc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	pc, err := RegisterPayChanCtrt(testAcnt0, tokId.B58Str().Str(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return pc
}

func (s *sysCtrtTest) test_Send(t *testing.T, sc *SysCtrt) {
	acnt0Old, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	acnt1Old, err := testAcnt1.Bal()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sc.Send(testAcnt0, testAcnt1.Addr.B58Str().Str(), 1.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	acnt0Bal, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	acnt1Bal, err := testAcnt1.Bal()
	if err != nil {
		t.Fatal(err)
	}

	amount, _ := NewVSYSForAmount(1.0)

	require.Equal(t, amount+FEE_EXEC_CTRT, acnt0Old-acnt0Bal)
	require.Equal(t, acnt1Old+amount, acnt1Bal)
}

func Test_SysCtrt_Send(t *testing.T) {
	sc := scT.newCtrt(t)
	scT.test_Send(t, sc)
}

func (s *sysCtrtTest) test_Transfer(t *testing.T, sc *SysCtrt) {
	acnt0Old, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	acnt1Old, err := testAcnt1.Bal()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sc.Transfer(testAcnt0, testAcnt0.Addr.B58Str().Str(), testAcnt1.Addr.B58Str().Str(), 1.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	acnt0Bal, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	acnt1Bal, err := testAcnt1.Bal()
	if err != nil {
		t.Fatal(err)
	}

	amount, _ := NewVSYSForAmount(1.0)

	require.Equal(t, amount+FEE_EXEC_CTRT, acnt0Old-acnt0Bal)
	require.Equal(t, acnt1Old+amount, acnt1Bal)
}

func Test_SysCtrt_Transfer(t *testing.T) {
	sc := scT.newCtrt(t)
	scT.test_Transfer(t, sc)
}

func (s *sysCtrtTest) test_DepositAndWithdraw(t *testing.T, sc *SysCtrt, pc *PayChanCtrt) {
	balInit, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := sc.Deposit(testAcnt0, pc.CtrtId.B58Str().Str(), 1.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	balAfterDeposit, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	amount, _ := NewVSYSForAmount(1.0)
	require.Equal(t, balInit-amount-FEE_EXEC_CTRT, balAfterDeposit)

	resp, err = sc.Withdraw(testAcnt0, pc.CtrtId.B58Str().Str(), 1.0, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	balAfterWithdraw, err := testAcnt0.Bal()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, balAfterDeposit+amount-FEE_EXEC_CTRT, balAfterWithdraw)
}

func Test_SysCtrt_DepositAndWithdraw(t *testing.T) {
	sc := scT.newCtrt(t)
	pc := scT.newPayChanCtrt(t, sc)

	scT.test_DepositAndWithdraw(t, sc, pc)
}

func Test_SysCtrt_AsWhole(t *testing.T) {
	sc := scT.newCtrt(t)
	scT.test_Send(t, sc)
	scT.test_Transfer(t, sc)
	pc := scT.newPayChanCtrt(t, sc)
	scT.test_DepositAndWithdraw(t, sc, pc)
}
