package euler

import (
	"testing"
)

func prob021() int {
	// the max sum of divisors < 10000 is 25320 so no need to check above that.
	_, _, _, dsumv := PrimeSieveDivisors(25320)

	asum := 0
	for i, ds := range dsumv {
		if dsumv[ds-i] == ds && 2*i != ds {
			asum += i
		}
		if i > 10000 {
			break
		}
	}

	return asum
}

func TestProb021(t *testing.T) {
	out := prob021()
	t.Log("Problem 021 ", out)
	if out != 31626 {
		t.Fail()
	}
}

func BenchmarkProb021(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob021()
	}
}
