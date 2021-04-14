package kate

import (
	"crypto/rand"
	bls12381 "github.com/kilic/bls12-381"
	"testing"
)

func setup(s bls12381.Fr, n int) katePubKey {
	bls := bls12381.NewEngine()
	G2 := bls.G2
	G1 := bls.G1

	fr := new(bls12381.Fr)

	temp := make([]bls12381.PointG1, n, n)
	temp2 := make([]bls12381.PointG2, n, n)

	g1 := G1.One()
	g2 := G2.One()
	r, _ := new(bls12381.Fr).Rand(rand.Reader)
	G1.MulScalar(g1, g1, r)
	r, _ = new(bls12381.Fr).Rand(rand.Reader)
	G2.MulScalar(g2, g2, r)

	pow := fr.One()
	tmp := fr.One()

	for i := 0; i < n; i++ {
		G1.MulScalar(&temp[i], g1, pow)
		G2.MulScalar(&temp2[i], g2, pow)
		tmp.Mul(pow, &s)
		pow.Set(tmp)

	}

	var pk katePubKey
	pk.G1s = temp
	pk.G2s = temp2

	return pk
}

func TestKatePubKey_Commit(t *testing.T) {

	g1 := bls12381.NewG1()
	s, _ := new(bls12381.Fr).Rand(rand.Reader)
	pk := setup(*s, 5)

	temp := g1.One()
	g1.MulScalar(temp, &pk.G1s[1], s)
	p := RandPoly(5)

	c, _ := pk.Commit(p)

	i, _ := new(bls12381.Fr).Rand(rand.Reader)

	wi, ati := pk.CreateWitness(*i, p)

	if pk.VerifyEval(*c, *i, ati, wi) != true {
		t.Fatal("whats wrong")
	}
}
