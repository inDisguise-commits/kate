package kate

import (
	bls12381 "github.com/kilic/bls12-381"
)

type Poly []bls12381.Fr

func NewPolyWithDegreeD(d int) Poly {
	p := make([]bls12381.Fr, d, d)
	return p
}

func (p *Poly) Set(p2 *Poly) *Poly {
	for i := 0; i < len(*p2); i++ {
		(*p)[i] = (*p2)[i]
	}
	return p
}

func (p *Poly) Zero() *Poly {
	for i := 0; i < len(*p); i++ {
		(*p)[i].Zero()
	}
	return p
}

func OrderedPoly(a, b Poly) (p1, p2 Poly) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

func Equal(a, b Poly) bool {
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

func Add(a, b Poly) Poly {

	p1, p2 := OrderedPoly(a, b)
	c := NewPolyWithDegreeD(len(p2))
	for i := 0; i < len(p1); i++ {
		c[i].Add(&p1[i], &p2[i])
	}
	for i := len(p1); i < len(p2); i++ {
		c[i] = p2[i]
	}
	return c
}
