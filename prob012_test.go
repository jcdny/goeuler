package euler

import (
	"testing"
)

func prob012(nd int) uint64 {
	pv, cv := PrimeSieve(100000)
	n := uint64(1)
	for i := uint64(2); NDivisors(n, pv, cv) < nd; n, i = n+i, i+1 {
	}
	return n
}

func TestProb012(t *testing.T) {
	out := prob012(500)
	Validate(t, 12, out)
}

func BenchmarkProb012(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob012(500)
	}
}
