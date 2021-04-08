package kate

import (
	bls12381 "github.com/kilic/bls12-381"
)

type poly []bls12381.Fr

func NewPoly() *poly {
	return &poly{}
}

//func (p *poly) Set (p2 *poly) *poly {
//	for i := 0; i < len(p2) ; i++ {
//		p[i] = p2[i]
//	}
//	return p
//}

//func (z *poly) Zero() *poly{
//	z= make([]bls12381.Fr, 1, 1)
//
//}

func Zero() []bls12381.Fr {
	z := make([]bls12381.Fr, 1, 1)
	z[0].Zero()
	return z
}

func OrderedPoly(a, b []bls12381.Fr) (p1, p2 []bls12381.Fr) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

func Equal(a, b []bls12381.Fr) bool {
	p1, p2 := OrderedPoly(a, b)
	for i := 0; i < len(p1); i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	for i := len(p1); i < len(p2); i++ {
		if !p2[i].IsZero() {
			return false
		}
	}
	return true
}

func Add(a []bls12381.Fr, b []bls12381.Fr) []bls12381.Fr {

	p1, p2 := OrderedPoly(a, b)
	c := make([]bls12381.Fr, len(p2), len(p2))

	for i := 0; i < len(p1); i++ {
		c[i].Add(&p1[i], &p2[i])
	}
	for i := len(p1); i < len(p2); i++ {
		c[i] = p2[i]
	}
	return c
}
