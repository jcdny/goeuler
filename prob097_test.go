package euler

import (
	"math/big"
	"testing"
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
	Validate(t, 97, out)
}

func BenchmarkProb097(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob097()
	}
}
