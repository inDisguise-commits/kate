package kate

import (
	"errors"
	bls12381 "github.com/kilic/bls12-381"
)

type Poly []bls12381.Fr

func NewPolyWithLenD(d int) Poly {
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

func (p *Poly) isZero() bool {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].IsZero() == false {
			return false
		}
	}
	return true
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
	c := NewPolyWithLenD(len(p2))
	for i := 0; i < len(p1); i++ {
		c[i].Add(&p1[i], &p2[i])
	}
	for i := len(p1); i < len(p2); i++ {
		c[i] = p2[i]
	}
	return c
}

func NegPoly(p Poly) Poly {
	r := NewPolyWithLenD(len(p))
	r.Set(&p)

	for i := 0; i < len(p); i++ {
		r[i].Neg(&r[i])
	}
	return r
}

func Eval(p Poly, x bls12381.Fr) bls12381.Fr {
	pow := *(new(bls12381.Fr).One())
	r := *(new(bls12381.Fr).Zero())
	temp := *(new(bls12381.Fr).Zero())

	for i := 0; i < len(p); i++ {

		temp.Mul(&p[i], &pow)
		r.Add(&r, &temp)
		pow.Mul(&pow, &x)
	}
	return r
}

func LongDivision(n, d Poly) (Poly, Poly, error) {
	if len(d) > len(n) {
		return nil, nil, errors.New("lengths are different")
	}

	degn := len(n) - 1
	degd := len(d) - 1
	q := NewPolyWithLenD(degn - degd + 1)
	q.Zero()
	r := NewPolyWithLenD(degn + 1)
	r.Set(&n)

	var coef, temp2 bls12381.Fr

	diff := degn - degd
	for diff >= 0 {
		coef.Inverse(&d[degd])
		coef.Mul(&coef, &r[degn])
		q[diff].Set(&coef)

		for i := degd; i >= 0; i-- {
			temp2.Mul(&coef, &d[i])
			temp2.Sub(&r[i+diff], &temp2)
			r[diff+i].Set(&temp2)
		}
		degn -= 1
		diff -= 1
	}

	return q, r, nil
}
