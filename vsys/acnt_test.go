package vsys

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type accountTest struct {
}

var aT *accountTest

func (a *accountTest) PRI_KEY() *PriKey {
	key, _ := NewPriKeyFromB58Str("EV5stVcWZ1kEQhrS7qcfYQdHpMHM5jwkyRxi9n9kXteZ")
	return key
}

func (a *accountTest) PUB_KEY() *PubKey {
	key, _ := NewPubKeyFromB58Str("4EyuJtDzQH15qAfnTPgqa8QB4ZU1dzqihdCs13UYEiV4")
	return key
}

func (a *accountTest) ADDR() *Addr {
	addr, _ := NewAddrFromB58Str("ATuQXbkZV4dCKsoFtXSCH5eKw92dMXQdUYU")
	return addr
}

func Test_Account_PriOnlyCons(t *testing.T) {
	acnt, err := NewAccount(testCh, aT.PRI_KEY())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, aT.PRI_KEY(), acnt.PriKey)
	assert.Equal(t, aT.PUB_KEY(), acnt.PubKey)
	assert.Equal(t, aT.ADDR(), acnt.Addr)
}

func Test_Account_Pay(t *testing.T) {
	t.Run("", func(t *testing.T) {
		const PAY_AMOUNT = 1
		amountMd, _ := NewVSYSForAmount(PAY_AMOUNT)

		testAcnt0BalBefore, _ := testAcnt0.Bal()
		testAcnt1BalBefore, _ := testAcnt1.Bal()

		resp, err := testAcnt0.Pay(testAcnt1.Addr.B58Str().Str(), PAY_AMOUNT, "")
		if err != nil {
			t.Fatal(err)
		}
		waitForBlock()
		assertTxSuccess(t, resp.Id.Str())
		assert.Nil(t, err)

		testAcnt0BalAfter, _ := testAcnt0.Bal()
		testAcnt1BalAfter, _ := testAcnt1.Bal()

		testAcnt0CostActual := testAcnt0BalBefore - testAcnt0BalAfter
		testAcnt0CostExpected := amountMd + FEE_PAYMENT
		assert.Equal(t, testAcnt0CostActual, testAcnt0CostExpected)

		testAcnt1GainActual := testAcnt1BalAfter - testAcnt1BalBefore
		t.Log("gainActual:  ", testAcnt1GainActual)
		testAcnt1GainExpected := amountMd
		t.Log("gainExpected: ", testAcnt1GainExpected)
		assert.Equal(t, testAcnt1GainActual, testAcnt1GainExpected)
	})
}

func Test_Account_LeaseAndCancelLease(t *testing.T) {
	effBalInit, err := testAcnt0.EffBal()
	if err != nil {
		t.Fatal(err)
	}

	amount, _ := NewVSYSForAmount(5)
	resp, err := testAcnt0.Lease(SUPERNODE_ADDR, amount.Amount())
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	leaseTxId := resp.Id.Str()
	assertTxSuccess(t, leaseTxId)

	effBalLease, err := testAcnt0.EffBal()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, effBalInit-amount-FEE_LEASING, effBalLease)

	resp2, err := testAcnt0.CancelLease(leaseTxId)
	waitForBlock()
	assertTxSuccess(t, resp2.Id.Str())

	effBalCancel, err := testAcnt0.EffBal()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, effBalLease+amount-FEE_LEASING_CANCEL, effBalCancel)

}
