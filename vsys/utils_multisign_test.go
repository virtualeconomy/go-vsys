package vsys

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type multisignTest struct {
	PriKey1 *MultiSignPriKey
	PriKey2 *MultiSignPriKey
	Msg     []byte
	Rand    []byte
}

func getMultiSignPriKeyFromB58Str(s string) *MultiSignPriKey {
	b, _ := B58Decode(s)
	p, _ := NewMultiSignPriKey(b)
	return p
}

var mulT = multisignTest{
	PriKey1: getMultiSignPriKeyFromB58Str("EV9ADJzYKZpk4MjxEkXxDSfRRSzBFnA9LEQNbepKZRFc"),
	PriKey2: getMultiSignPriKeyFromB58Str("3hQRGJkqKFbks77cZ12ugHxDtbweH3EZjhfVzfr4RqPs"),
	Msg:     []byte("test"),
	Rand:    []byte{178, 152, 201, 209, 148, 127, 203, 131, 115, 162, 0, 35, 43, 111, 108, 121, 247, 133, 72, 20, 199, 78, 102, 97, 59, 46, 160, 200, 67, 189, 173, 80, 203, 60, 131, 115, 148, 47, 10, 138, 10, 170, 218, 140, 35, 140, 1, 18, 121, 85, 79, 116, 89, 186, 181, 96, 182, 235, 67, 153, 42, 175, 56, 134},
}

func Test_MultiSign_OneKey(t *testing.T) {
	varA := mulT.PriKey1.VarA
	allAs := [][]byte{varA}

	xA := mulT.PriKey1.GetxA(allAs)
	xAs := []*MultiSignPoint{xA}
	unionA := MultiSignGetUnionA(xAs)

	varR := mulT.PriKey1.GetR(mulT.Msg, mulT.Rand)
	Rs := []*MultiSignPoint{varR}
	unionR := MultiSignGetUnionR(Rs)

	subSig := mulT.PriKey1.Sign(mulT.Msg, mulT.Rand, unionA, unionR, allAs)
	sigs := []*big.Int{subSig}
	mulSig := MultiSignGetSig(unionA, unionR, sigs)

	bpA := mulT.PriKey1.GetbpA(allAs)
	bpAs := []*MultiSignPoint{bpA}
	mulPub := MultiSignGetPub(bpAs)

	rawSig := SignImpl(mulT.PriKey1.PriKey, mulT.Msg, mulT.Rand)
	assert.Equal(t, rawSig, mulSig)

	rawPub, _ := GenPubKey(mulT.PriKey1.PriKey)
	assert.Equal(t, rawPub, mulPub)

	valid := Verify(mulPub, mulT.Msg, mulSig)
	assert.True(t, valid)
}

func Test_MultiSign_TwoKeys(t *testing.T) {
	varA1 := mulT.PriKey1.VarA
	varA2 := mulT.PriKey2.VarA

	allAs := [][]byte{varA1, varA2}

	xA1 := mulT.PriKey1.GetxA(allAs)
	xA2 := mulT.PriKey2.GetxA(allAs)

	xAs := []*MultiSignPoint{xA1, xA2}
	unionA := MultiSignGetUnionA(xAs)

	varR1 := mulT.PriKey1.GetR(mulT.Msg, mulT.Rand)
	varR2 := mulT.PriKey2.GetR(mulT.Msg, mulT.Rand)
	Rs := []*MultiSignPoint{varR1, varR2}
	unionR := MultiSignGetUnionR(Rs)

	subSig1 := mulT.PriKey1.Sign(mulT.Msg, mulT.Rand, unionA, unionR, allAs)
	subSig2 := mulT.PriKey2.Sign(mulT.Msg, mulT.Rand, unionA, unionR, allAs)

	sigs := []*big.Int{subSig1, subSig2}
	mulSig := MultiSignGetSig(unionA, unionR, sigs)

	bpA1 := mulT.PriKey1.GetbpA(allAs)
	bpA2 := mulT.PriKey2.GetbpA(allAs)
	bpAs := []*MultiSignPoint{bpA1, bpA2}
	mulPub := MultiSignGetPub(bpAs)

	valid := Verify(mulPub, mulT.Msg, mulSig)
	assert.True(t, valid)
}
