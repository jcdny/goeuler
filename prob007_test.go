package euler

import (
	"testing"
	"log"
)

func prob007() int {
	p, _ := PrimeSieve(105000)
	return p[10000]
}

func TestProb007(t *testing.T) {
	if prob007() != 104743 {
		log.Print(prob007())
		t.Error("Prob007 failed")
	}
}

func BenchmarkProb007(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob007()
	}
}
