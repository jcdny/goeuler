package euler

import (
	"math/big"
	"testing"
)

func ksearch(n int64, t *big.Int) int64 {
	b := big.NewInt(0)
	i, k := int64(0), (n+1)/2
	for i < k {
		h := i + (k-i)/2
		if b.Binomial(n, h).Cmp(t) < 0 {
			i = h + 1
		} else {
			k = h
		}
	}

	return k
}

func prob053() int {
	z := big.NewInt(1000000)
	N := int64(0)
	for n := int64(23); n <= 100; n++ {
		// binary search for k such that C(n,k) > 1000000
		k := ksearch(n, z)
		// now since binomial function is symmetric we have an
		// expression for the number of terms > 1mm
		N += n - 2*k + 1
	}

	return int(N)
}

func TestProb053(t *testing.T) {
	out := prob053()
	Validate(t, 53, out)
}

func BenchmarkProb053(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob053()
	}
}
