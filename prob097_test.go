package euler

import (
	"testing"
	"big"
)

func prob097() int64 {
	za := big.NewInt(2)
	zb := big.NewInt(7830457)
	zm := big.NewInt(10000000000)
	zb.Exp(za, zb, zm)
	zc := big.NewInt(28433)
	zc.Mul(zb, zc)
	res := zc.Mod(zc, zm).Int64()
	res++
	return res
}

func TestProb097(t *testing.T) {
	out := prob097()
	t.Log("Problem 097 ", out)
	if out != 8739992577 {
		t.Fail()
	}
}

func BenchmarkProb097(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob097()
	}
}
