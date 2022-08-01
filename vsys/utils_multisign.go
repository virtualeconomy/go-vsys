package vsys

import (
	pkgsha512 "crypto/sha512"
	"fmt"
	"math/big"
)

var (
	C252 = big.NewInt(0).Lsh(
		big.NewInt(1),
		252,
	)
	C255 = big.NewInt(0).Lsh(
		big.NewInt(1),
		255,
	)
	C25519 = func() *big.Int {
		return big.NewInt(0).Sub(
			C255,
			big.NewInt(19),
		)
	}()
	BASE_FIELD_Z_P = C25519
)

func modpInv(x *big.Int) *big.Int {
	return big.NewInt(0).Exp(
		x,
		big.NewInt(0).Sub(BASE_FIELD_Z_P, big.NewInt(2)),
		BASE_FIELD_Z_P,
	)
}

var (
	CURVE_CONST_D = func() *big.Int {
		return big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				big.NewInt(-121665),
				modpInv(big.NewInt(121666)),
			),
			BASE_FIELD_Z_P,
		)
	}()

	GROUP_ORDER_Q = func() *big.Int {
		t, _ := big.NewInt(0).SetString("27742317777372353535851937790883648493", 10)
		return big.NewInt(0).Add(
			C252,
			t,
		)
	}()
)

// reverseBytes reverse the given byte slice and returns the slice.
func reverseBytes(b []byte, inPlace bool) []byte {
	s := b

	if !inPlace {
		s = make([]byte, len(b))
		copy(s, b)
	}

	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

func sha512(s []byte) [64]byte {
	return pkgsha512.Sum512(s)
}

func sha512Modq(s []byte) *big.Int {
	h := sha512(s)
	return big.NewInt(0).Mod(
		big.NewInt(0).SetBytes(
			reverseBytes(h[:], true),
		),
		GROUP_ORDER_Q,
	)
}

func sha512Modp(s []byte) *big.Int {
	h := sha512(s)
	return big.NewInt(0).Mod(
		big.NewInt(0).SetBytes(
			reverseBytes(h[:], true),
		),
		BASE_FIELD_Z_P,
	)
}

type MultiSignPoint struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
	T *big.Int
}

func (p *MultiSignPoint) Add(other *MultiSignPoint) *MultiSignPoint {
	pYXSub := big.NewInt(0).Sub(p.Y, p.X)
	otherYXSub := big.NewInt(0).Sub(other.Y, other.X)

	A := big.NewInt(0).Mod(
		big.NewInt(0).Mul(pYXSub, otherYXSub),
		BASE_FIELD_Z_P,
	)

	pYXAdd := big.NewInt(0).Add(p.Y, p.X)
	otherYXAdd := big.NewInt(0).Add(other.Y, other.X)

	B := big.NewInt(0).Mod(
		big.NewInt(0).Mul(pYXAdd, otherYXAdd),
		BASE_FIELD_Z_P,
	)

	C := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0).Mul(
				big.NewInt(0).Mul(
					big.NewInt(2),
					p.T,
				),
				other.T,
			),
			CURVE_CONST_D,
		),
		BASE_FIELD_Z_P,
	)

	D := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0).Mul(
				big.NewInt(2),
				p.Z,
			),
			other.Z,
		),
		BASE_FIELD_Z_P,
	)

	E := big.NewInt(0).Sub(B, A)
	F := big.NewInt(0).Sub(D, C)
	G := big.NewInt(0).Add(D, C)
	H := big.NewInt(0).Add(B, A)

	return &MultiSignPoint{
		X: big.NewInt(0).Mul(E, F),
		Y: big.NewInt(0).Mul(G, H),
		Z: big.NewInt(0).Mul(F, G),
		T: big.NewInt(0).Mul(E, H),
	}
}

func (p *MultiSignPoint) Mul(s *big.Int) *MultiSignPoint {
	P := p
	Q := &MultiSignPoint{
		X: big.NewInt(0),
		Y: big.NewInt(1),
		Z: big.NewInt(1),
		T: big.NewInt(0),
	}

	n := big.NewInt(0).Set(s)

	for n.Cmp(big.NewInt(0)) > 0 {
		sAnd := big.NewInt(0).And(n, big.NewInt(1))
		if sAnd.Cmp(big.NewInt(0)) > 0 {
			Q = Q.Add(P)
		}
		P = P.Add(P)
		n.Rsh(n, 1)
	}
	return Q
}

func (p *MultiSignPoint) Compress() []byte {
	zinv := modpInv(p.Z)
	x := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			p.X,
			zinv,
		),
		BASE_FIELD_Z_P,
	)
	y := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			p.Y,
			zinv,
		),
		BASE_FIELD_Z_P,
	)

	b := big.NewInt(0).Or(
		y,
		big.NewInt(0).Lsh(
			big.NewInt(0).And(
				x,
				big.NewInt(1),
			),
			255,
		),
	).Bytes()

	return reverseBytes(b, true)
}

func (p *MultiSignPoint) PubKey() []byte {
	zinv := modpInv(p.Y)
	x := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0),
			zinv,
		),
		BASE_FIELD_Z_P,
	)

	y := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			p.Y,
			zinv,
		),
		BASE_FIELD_Z_P,
	)

	b := big.NewInt(0).Or(
		y,
		big.NewInt(0).Lsh(
			big.NewInt(0).And(x, big.NewInt(1)),
			255,
		),
	).Bytes()

	return reverseBytes(b, true)
}

func NewMultiSignPointFromBytes(b []byte) (*MultiSignPoint, error) {
	if len(b) != 32 {
		return nil, fmt.Errorf("Invalid input length")
	}

	y := big.NewInt(0).SetBytes(reverseBytes(b, false))
	sign := big.NewInt(0).Rsh(y, 255)
	y.And(
		y,
		big.NewInt(0).Sub(
			C255,
			big.NewInt(1),
		),
	)

	x, err := PointRecoverX(y, sign)
	if err != nil {
		return nil, err
	}

	return &MultiSignPoint{
		X: x,
		Y: y,
		Z: big.NewInt(1),
		T: big.NewInt(0).Mod(
			big.NewInt(0).Mul(x, y),
			BASE_FIELD_Z_P,
		),
	}, nil
}

func PointRecoverX(y, sign *big.Int) (*big.Int, error) {
	if y.Cmp(BASE_FIELD_Z_P) >= 0 {
		return nil, fmt.Errorf("Invalid y")
	}

	x2 := big.NewInt(0).Mul(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(y, y),
			big.NewInt(1),
		),
		modpInv(
			big.NewInt(0).Add(
				big.NewInt(0).Mul(
					big.NewInt(0).Mul(CURVE_CONST_D, y),
					y,
				),
				big.NewInt(1),
			),
		),
	)

	if x2.Cmp(big.NewInt(0)) == 0 {
		if sign.Cmp(big.NewInt(0)) > 0 {
			return nil, fmt.Errorf("Invalid x2 & sign")
		}
		return big.NewInt(0), nil
	}

	modpSqrtM1 := big.NewInt(0).Exp(
		big.NewInt(2),
		big.NewInt(0).Div(
			big.NewInt(0).Sub(
				BASE_FIELD_Z_P,
				big.NewInt(1),
			),
			big.NewInt(4),
		),
		BASE_FIELD_Z_P,
	)

	x := big.NewInt(0).Exp(
		x2,
		big.NewInt(0).Div(
			big.NewInt(0).Add(
				BASE_FIELD_Z_P,
				big.NewInt(3),
			),
			big.NewInt(8),
		),
		BASE_FIELD_Z_P,
	)

	checked := false

	for big.NewInt(0).Mod(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(x, x),
			x2,
		),
		BASE_FIELD_Z_P,
	).Cmp(big.NewInt(0)) != 0 {
		if checked {
			return nil, fmt.Errorf("Invalid x")
		}
		x = big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				x,
				modpSqrtM1,
			),
			BASE_FIELD_Z_P,
		)
		checked = true
	}

	if big.NewInt(0).And(x, big.NewInt(1)).Cmp(sign) != 0 {
		x = big.NewInt(0).Sub(BASE_FIELD_Z_P, x)
	}

	return x, nil
}

func NewMultiSignPointG() (*MultiSignPoint, error) {
	y := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(4),
			modpInv(big.NewInt(5)),
		),
		BASE_FIELD_Z_P,
	)

	x, err := PointRecoverX(y, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	return &MultiSignPoint{
		X: x,
		Y: y,
		Z: big.NewInt(1),
		T: big.NewInt(0).Mod(
			big.NewInt(0).Mul(x, y),
			BASE_FIELD_Z_P,
		),
	}, nil
}

var G, _ = NewMultiSignPointG()

type MultiSignPriKey struct {
	PriKey []byte
	Vara   *big.Int
	VarA   []byte
	PubKey []byte
}

func NewMultiSignPriKey(b []byte) (*MultiSignPriKey, error) {
	p := &MultiSignPriKey{
		PriKey: b,
	}

	p.Vara = p.Geta()
	p.VarA = p.GetA()

	var err error
	p.PubKey, err = p.GetPubKey()
	if err != nil {
		return nil, fmt.Errorf("NewMultiSignPriKey: %w", err)
	}

	return p, nil
}

func (p *MultiSignPriKey) Geta() *big.Int {
	return big.NewInt(0).SetBytes(reverseBytes(p.PriKey, false))
}

func (p *MultiSignPriKey) GetA() []byte {
	return G.Mul(p.Vara).Compress()
}

func (p *MultiSignPriKey) GetPubKey() ([]byte, error) {
	if len(p.PriKey) != 32 {
		return nil, fmt.Errorf("Bad size of private key")
	}

	h := sha512(p.PriKey)
	a := big.NewInt(0).SetBytes(reverseBytes(h[:32], false))
	a.And(
		a,
		big.NewInt(0).Sub(
			big.NewInt(0).Lsh(
				big.NewInt(1),
				254,
			),
			big.NewInt(8),
		),
	)
	a.Or(
		a,
		big.NewInt(0).Lsh(
			big.NewInt(1),
			254,
		),
	)
	return G.Mul(a).Compress(), nil
}

func (p *MultiSignPriKey) Getr(msg, rand []byte) *big.Int {
	prefix := big.NewInt(0xFE)
	for i := 0; i < 31; i++ {
		prefix.Mul(prefix, big.NewInt(256))
		prefix.Add(prefix, big.NewInt(0xFF))
	}
	prefixBytes := prefix.Bytes()

	l := len(prefixBytes) + len(p.PriKey) + len(msg) + len(rand)
	b := make([]byte, 0, l)
	b = append(b, prefixBytes...)
	b = append(b, p.PriKey...)
	b = append(b, msg...)
	b = append(b, rand...)

	return sha512Modq(b)
}

func (p *MultiSignPriKey) GetR(msg, rand []byte) *MultiSignPoint {
	r := p.Getr(msg, rand)
	return G.Mul(r)
}

func (p *MultiSignPriKey) Getx(allAs [][]byte) *big.Int {
	if len(allAs) == 1 {
		return big.NewInt(1)
	}

	prefix := big.NewInt(0xFD)
	for i := 0; i < 31; i++ {
		prefix.Mul(prefix, big.NewInt(256))
		prefix.Add(prefix, big.NewInt(0xFF))
	}
	prefixBytes := prefix.Bytes()

	l := len(p.VarA) + len(prefixBytes)
	for _, Ai := range allAs {
		l += len(Ai)
	}

	b := make([]byte, 0, l)
	b = append(b, prefixBytes...)
	b = append(b, p.VarA...)
	for _, Ai := range allAs {
		b = append(b, Ai...)
	}

	return sha512Modq(b)
}

func (p *MultiSignPriKey) GetbpA(allAs [][]byte) *MultiSignPoint {
	x := p.Getx(allAs)
	n := big.NewInt(0).Mod(
		big.NewInt(0).Mul(x, p.Vara),
		GROUP_ORDER_Q,
	)
	return G.Mul(n)
}

func (p *MultiSignPriKey) GetxA(allAs [][]byte) *MultiSignPoint {
	x := p.Getx(allAs)
	n := big.NewInt(0).Mul(x, p.Vara)
	return G.Mul(n)
}

func (p *MultiSignPriKey) Sign(msg, rand, unionA []byte, unionR *MultiSignPoint, allAs [][]byte) *big.Int {
	r := p.Getr(msg, rand)
	x := p.Getx(allAs)
	a := p.Vara

	unionRBytes := unionR.Compress()
	l := len(unionRBytes) + len(unionA) + len(msg)
	b := make([]byte, 0, l)
	b = append(b, unionRBytes...)
	b = append(b, unionA...)
	b = append(b, msg...)

	h := sha512Modq(b)

	return big.NewInt(0).Mod(
		big.NewInt(0).Mod(
			big.NewInt(0).Add(
				r,
				big.NewInt(0).Mul(
					big.NewInt(0).Mul(h, x),
					a,
				),
			),
			GROUP_ORDER_Q,
		),
		GROUP_ORDER_Q,
	)
}

func MultiSignGetUnionA(xAs []*MultiSignPoint) []byte {
	p := xAs[0]
	for i := 1; i < len(xAs); i++ {
		p = p.Add(xAs[i])
	}
	return p.Compress()
}

func MultiSignGetUnionR(Rs []*MultiSignPoint) *MultiSignPoint {
	p := Rs[0]
	for i := 1; i < len(Rs); i++ {
		p = p.Add(Rs[i])
	}
	return p
}

func MultiSignTransferSig(sig *big.Int, A []byte) []byte {
	b := reverseBytes(sig.Bytes(), false)
	b[31] = (b[31] & 0x7F) | (A[31] & 0x80)
	return b
}

func MultiSignGetSig(unionA []byte, unionR *MultiSignPoint, sigs []*big.Int) []byte {
	s := big.NewInt(0)
	for _, sig := range sigs {
		s.Add(s, sig)
	}
	s.Mod(s, GROUP_ORDER_Q)

	unionRBytes := unionR.Compress()
	transSig := MultiSignTransferSig(s, unionA)

	l := len(unionRBytes) + len(transSig)
	b := make([]byte, 0, l)

	b = append(b, unionRBytes...)
	b = append(b, transSig...)
	return b
}

func MultiSignGetPub(bpAs []*MultiSignPoint) []byte {
	p := bpAs[0]

	for i := 1; i < len(bpAs); i++ {
		p = p.Add(bpAs[i])
	}

	zinv := modpInv(p.Z)
	y := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			p.Y,
			zinv,
		),
		BASE_FIELD_Z_P,
	)

	n := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0).Add(y, big.NewInt(1)),
			modpInv(
				big.NewInt(0).Sub(
					big.NewInt(1),
					y,
				),
			),
		),
		BASE_FIELD_Z_P,
	)

	return reverseBytes(n.Bytes(), true)
}
