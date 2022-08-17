package vsys

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type payChanTest struct {
}

func (p *payChanTest) TOK_MAX() float64 {
	return 100
}
func (p *payChanTest) TOK_UNIT() uint64 {
	return 1
}
func (p *payChanTest) INIT_LOAD() float64 {
	return p.TOK_MAX() / 2.0
}

var pcT *payChanTest

func (p *payChanTest) newTokCtrt(t *testing.T) *TokCtrtWithoutSplit {
	tc, err := RegisterTokCtrtWithoutSplit(testAcnt0, p.TOK_MAX(), p.TOK_UNIT(), "", "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	_, err = tc.Issue(testAcnt0, p.TOK_MAX(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return tc
}

func (p *payChanTest) newCtrt(t *testing.T, tc *TokCtrtWithoutSplit) *PayChanCtrt {
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}

	pc, err := RegisterPayChanCtrt(testAcnt0, tokId.B58Str().Str(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()

	resp, err := tc.Deposit(testAcnt0, pc.CtrtId.B58Str().Str(), p.TOK_MAX(), "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	return pc
}

func (p *payChanTest) test_Register(t *testing.T, tc *TokCtrtWithoutSplit, pc *PayChanCtrt) {
	maker, err := pc.Maker()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, testAcnt0.Addr, maker)
	tokId, err := tc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	pcTokId, err := pc.TokId()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, tokId, pcTokId)
	ctrtBal, err := pc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, p.TOK_MAX(), ctrtBal.Amount())
}

func Test_PayChanCtrt_Register(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc := pcT.newCtrt(t, tc)

	pcT.test_Register(t, tc, pc)
}

func (p *payChanTest) test_CreateAndLoad(t *testing.T, pc *PayChanCtrt) string {
	load_amount := p.INIT_LOAD()
	later := time.Now().Unix() + 60*10 // 10 min from now

	resp, err := pc.CreateAndLoad(testAcnt0, testAcnt1.Addr.B58Str().Str(), load_amount, later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	chanId := resp.Id.Str()

	chanCreator, err := pc.GetChanCreator(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testAcnt0.Addr, chanCreator)
	chanCreatorPubKey, err := pc.GetChanCreatorPubKey(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testAcnt0.PubKey, chanCreatorPubKey)
	chanAccumLoad, err := pc.GetChanAccumLoad(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, load_amount, chanAccumLoad.Amount())
	chanAccumPay, err := pc.GetChanAccumPay(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0.0, chanAccumPay.Amount())
	chanExpTime, err := pc.GetChanExpTime(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, later, chanExpTime.UnixTs())
	chanStatus, err := pc.GetChanStatus(chanId)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, chanStatus)
	return chanId
}

func (p *payChanTest) newCtrtWithChan(t *testing.T, tc *TokCtrtWithoutSplit) (*PayChanCtrt, string) {
	pc := p.newCtrt(t, tc)

	load_amount := p.INIT_LOAD()
	later := time.Now().Unix() + 60*10 // 10 min from now

	resp, err := pc.CreateAndLoad(testAcnt0, testAcnt1.Addr.B58Str().Str(), load_amount, later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	return pc, resp.Id.Str()
}

func Test_PayChanCtrt_CreateAndLoad(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc := pcT.newCtrt(t, tc)

	pcT.test_CreateAndLoad(t, pc)
}

func (p *payChanTest) test_ExtendExpTime(t *testing.T, pc *PayChanCtrt, chanId string) {
	chanExpTimeOld, err := pc.GetChanExpTime(chanId)
	if err != nil {
		t.Fatal(err)
	}

	newLater := chanExpTimeOld.UnixTs() + 300
	resp, err := pc.ExtendExpTime(testAcnt0, chanId, newLater, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	chanExpTime, err := pc.GetChanExpTime(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, newLater, chanExpTime.UnixTs())
}

func Test_PayChanCtrt_ExtendExpTime(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc, chanId := pcT.newCtrtWithChan(t, tc)

	pcT.test_ExtendExpTime(t, pc, chanId)
}

func (p *payChanTest) test_Load(t *testing.T, pc *PayChanCtrt, chanId string) {
	chanLoadOld, err := pc.GetChanAccumLoad(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, p.INIT_LOAD(), chanLoadOld.Amount())

	moreLoad := p.INIT_LOAD() / 2

	resp, err := pc.Load(testAcnt0, chanId, moreLoad, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	chanLoad, err := pc.GetChanAccumLoad(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, p.INIT_LOAD()+moreLoad, chanLoad.Amount())
}

func Test_PayChanCtrt_Load(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc, chanId := pcT.newCtrtWithChan(t, tc)

	pcT.test_Load(t, pc, chanId)
}

func (p *payChanTest) test_Abort(t *testing.T, pc *PayChanCtrt, chanId string) {
	chanStatus, err := pc.GetChanStatus(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, chanStatus)

	resp, err := pc.Abort(testAcnt0, chanId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	chanStatus, err = pc.GetChanStatus(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.False(t, chanStatus)
}

func Test_PayChanCtrt_Abort(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc, chanId := pcT.newCtrtWithChan(t, tc)

	pcT.test_Abort(t, pc, chanId)
}

func (p *payChanTest) test_Unload(t *testing.T, pc *PayChanCtrt) {
	loadAmount := p.TOK_MAX() / 10
	later := time.Now().Unix() + int64(avgBlockDelay.Seconds())*2
	resp, err := pc.CreateAndLoad(testAcnt0, testAcnt1.Addr.B58Str().Str(), loadAmount, later, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	chanId := resp.Id.Str()

	balOld, err := pc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(avgBlockDelay * 2)

	resp, err = pc.Unload(testAcnt0, chanId, "")
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	bal, err := pc.GetCtrtBal(testAcnt0.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, balOld.Amount()+loadAmount, bal.Amount())
}

func Test_PayChanCtrt_Unload(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc := pcT.newCtrt(t, tc)

	pcT.test_Unload(t, pc)
}

func (p *payChanTest) test_OffchainPayAndCollectPayment(t *testing.T, pc *PayChanCtrt, chanId string) {
	sig, err := pc.OffchainPay(
		testAcnt0.PriKey,
		chanId,
		p.INIT_LOAD(),
	)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := pc.CollectPayment(
		testAcnt1,
		chanId,
		p.INIT_LOAD(),
		sig,
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	waitForBlock()
	assertTxSuccess(t, resp.Id.Str())

	accumPay, err := pc.GetChanAccumPay(chanId)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, p.INIT_LOAD(), accumPay.Amount())
	acnt1Bal, err := pc.GetCtrtBal(testAcnt1.Addr.B58Str().Str())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, p.INIT_LOAD(), acnt1Bal.Amount())
}

func Test_PayChanCtrt_OffchainPayAndCollectPayment(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc, chanId := pcT.newCtrtWithChan(t, tc)

	pcT.test_OffchainPayAndCollectPayment(t, pc, chanId)
}

func Test_PayChanCtrt_AsWhole(t *testing.T) {
	tc := pcT.newTokCtrt(t)
	pc := pcT.newCtrt(t, tc)

	pcT.test_Register(t, tc, pc)
	chanId := pcT.test_CreateAndLoad(t, pc)

	pcT.test_ExtendExpTime(t, pc, chanId)
	pcT.test_Load(t, pc, chanId)
	pcT.test_OffchainPayAndCollectPayment(t, pc, chanId)
	pcT.test_Abort(t, pc, chanId)
	pcT.test_Unload(t, pc)
}
