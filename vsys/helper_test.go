package vsys

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	AVG_BLOCK_DELAY = os.Getenv("GO_VSYS_AVG_BLOCK_DELAY")
	HOST            = os.Getenv("GO_VSYS_HOST")
	SEED            = os.Getenv("GO_VSYS_SEED")
	SUPERNODE_ADDR  = os.Getenv("GO_VSYS_SUPERNODE_ADDR")
)

var (
	avgBlockDelay time.Duration
	testApi       = NewNodeAPI(HOST)
	testCh        = NewChain(testApi, TEST_NET)
	testWal       *Wallet
	testAcnt0     *Account
	testAcnt1     *Account
	testAcnt2     *Account
)

func init() {
	abdVal, _ := strconv.Atoi(AVG_BLOCK_DELAY)
	avgBlockDelay = time.Duration(abdVal) * time.Second
	testWal, _ = NewWalletFromSeedStr(SEED)
	testAcnt0, _ = testWal.GetAccount(testCh, 0)
	testAcnt1, _ = testWal.GetAccount(testCh, 1)
	testAcnt2, _ = testWal.GetAccount(testCh, 2)
}

func waitForBlock() {
	time.Sleep(avgBlockDelay)
}

func assertTxStatus(t *testing.T, txId, status string) {
	tx, err := testApi.GetTxInfo(txId)
	if err != nil {
		t.Log(err)
	}
	require.Equal(t, tx.GetTxGeneral().Status.Str(), status)
}

func assertTxSuccess(t *testing.T, txId string) {
	assertTxStatus(t, txId, "Success")
}
