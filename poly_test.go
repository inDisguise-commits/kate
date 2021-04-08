package kate

import (
	"crypto/rand"
	bls12381 "github.com/kilic/bls12-381"
	"testing"
)

func RandPoly(l int) []bls12381.Fr {
	p := make([]bls12381.Fr, l)
	for i := 0; i < l; i++ {
		z, _ := new(bls12381.Fr).Rand(rand.Reader)
		p[i].Set(z)
	}
	return p
}

func TestOrderedPoly(t *testing.T) {
	p1 := RandPoly(5)
	p2 := RandPoly(3)
	p3 := RandPoly(7)
	a, b := OrderedPoly(p1, p2)
	if len(a) != len(p2) || len(b) != len(p1) {
		t.Error("p1 was bigger")
	}

	a, b = OrderedPoly(p1, p3)
	if len(a) != len(p1) || len(b) != len(p3) {
		t.Error("p3 was bigger")
	}

}

func TestAdditionProperties(t *testing.T) {
	zeroPoly := Zero()
	p := RandPoly(8)
	if !Equal(Add(zeroPoly, p), p) {
		t.Fatal("0+p != p")
	}
	q := RandPoly(11)
	r := RandPoly(7)
	if !Equal(Add(p, q), Add(q, p)) {
		t.Fatal("p+q != q+p")
	}
	if !Equal(Add(Add(p, q), r), Add(p, Add(q, r))) {
		t.Fatal("(p+q)+r != p+(q+r)")
	}

}
