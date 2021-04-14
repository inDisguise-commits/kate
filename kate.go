package kate

import (
	"errors"
	"fmt"
	bls12381 "github.com/kilic/bls12-381"
)

type katePubKey struct {
	G1s []bls12381.PointG1
	G2s []bls12381.PointG2
}

func (pk *katePubKey) NewKatePubKey(g1s []bls12381.PointG1, g2s []bls12381.PointG2) (*katePubKey, error) {
	if len(g1s) != len(g2s) {
		return nil, errors.New("lengths are different")
	}
	pk.G1s = g1s
	pk.G2s = g2s
	return pk, nil
}

func (pk *katePubKey) Commit(p Poly) (*bls12381.PointG1, error) {

	g1 := bls12381.NewG1()
	if len(p) != len(pk.G1s) {
		return nil, errors.New("lengths are different")
	}

	r := g1.New()

	tempP := make([]*bls12381.Fr, len(p), len(p))
	tempPkG1 := make([]*bls12381.PointG1, len(p), len(p))

	for i := 0; i < len(p); i++ {
		tempP[i] = &p[i]
		tempPkG1[i] = &pk.G1s[i]
	}

	_, _ = g1.MultiExp(r, tempPkG1, tempP)
	return r, nil
}

func (pk *katePubKey) CreateWitness(i bls12381.Fr, p Poly) (bls12381.PointG1, bls12381.Fr) {
	g1 := bls12381.NewG1()
	r := g1.New()
	var temp bls12381.Fr

	q := NewPolyWithLenD(2)
	temp.Neg(&i)
	q[0].Set(&temp)
	q[1] = *(new(bls12381.Fr).One())

	//bls12381.Fr.One(&q[1])

	pCopy := NewPolyWithLenD(len(p))
	pCopy.Set(&p)

	ati := Eval(p, i)
	temp.Sub(&pCopy[0], &ati)
	pCopy[0].Set(&temp)

	psi, _, _ := LongDivision(pCopy, q)

	tempPsi := make([]*bls12381.Fr, len(psi), len(psi))
	tempPkG1 := make([]*bls12381.PointG1, len(psi), len(psi))

	for i := 0; i < len(psi); i++ {
		tempPsi[i] = &psi[i]
		tempPkG1[i] = &pk.G1s[i]
		fmt.Println(i)
	}
	_, _ = g1.MultiExp(r, tempPkG1, tempPsi)

	return *r, ati
}

func (pk *katePubKey) VerifyEval(c bls12381.PointG1, i bls12381.Fr, ati bls12381.Fr, wi bls12381.PointG1) bool {
	bls := bls12381.NewEngine()
	gt := bls.GT()
	g2 := bls.G2
	var temp bls12381.PointG2

	eleft := bls.AddPair(&c, &(pk.G2s)[0]).Result()

	g2.MulScalar(&temp, &pk.G2s[0], &i)
	g2.Sub(&temp, &pk.G2s[1], &temp)

	eright1 := bls.AddPair(&wi, &temp).Result()

	temp2 := bls.AddPair(&(pk.G1s)[0], &(pk.G2s)[0]).Result()
	gt.Exp(temp2, temp2, ati.ToBig())

	gt.Mul(temp2, temp2, eright1)

	return eleft.Equal(temp2)

}
