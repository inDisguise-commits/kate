package kate

import (
	"errors"
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

func (pk *katePubKey) Commit(p []bls12381.Fr) (*bls12381.PointG1, error) {
	g := new(bls12381.G1)

	if len(p) != len(pk.G1s) {
		return nil, errors.New("lengths are different")
	}
	var r bls12381.PointG1

	tempP := make([]*bls12381.Fr, len(p), len(p))
	tempPkG1 := make([]*bls12381.PointG1, len(p), len(p))

	for i := 0; i < len(p); i++ {
		tempP[i] = &p[i]
		tempPkG1[i] = &pk.G1s[i]
	}
	_, _ = g.MultiExp(&r, tempPkG1, tempP)
	return &r, nil
}
