package vsys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		print("gainActual: ", testAcnt1GainActual)
		testAcnt1GainExpected := amountMd
		print("gainExpected: ", testAcnt1GainExpected)
		assert.Equal(t, testAcnt1GainActual, testAcnt1GainExpected)
	})
}
