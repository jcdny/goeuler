package euler

import (
	"testing"
	"log"
)

func prob012(nd int) int64 {
	pv, cv := PrimeSieve(100000)
	n := int64(1)
	for i := int64(2); NDivisors(n, pv, cv) < nd; n, i = n+i, i+1 {
	}
	return n
}

func TestProb012(t *testing.T) {
	out := prob012(500)
	log.Print("Prob012 ", out)
	if out != 76576500 {
		t.Error("Prob012 failed ", out)
	}
}

func BenchmarkProb012(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob012(500)
	}
}
