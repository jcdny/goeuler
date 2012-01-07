package euler

import (
	"testing"
	"log"
)

func prob010() uint64 {
	ps, _ := PrimeSieve(2000000)
	n := uint64(0)
	for _, p := range ps {
		n += uint64(p)
	}
	return n
}

func TestProb010(t *testing.T) {
	if prob010() != 142913828922 {
		log.Print(prob010())
		t.Error("Prob010 failed")
	}
}

func BenchmarkProb010(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob010()
	}
}
