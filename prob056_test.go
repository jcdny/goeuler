package euler

import (
	"math/big"
	"testing"
)

func prob056() int {
	M := int64(100)
	za := big.NewInt(0)
	zb := big.NewInt(0)
	msum := 0
	for a := int64(2); a < M; a++ {
		za.SetInt64(a)
		for b := int64(2); b < M; b++ {
			zb.SetInt64(b)
			zb.Exp(za, zb, nil)
			nb := SumDigits(zb)
			if nb > msum {
				//log.Print(a, b, nb)
				msum = nb
			}
		}
	}

	return msum
}

func TestProb056(t *testing.T) {
	out := prob056()
	Validate(t, 56, out)
}

func BenchmarkProb056(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob056()
	}
}
