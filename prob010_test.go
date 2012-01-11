package euler

import (
	"testing"
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
	out := prob010()
	t.Log("Problem 010 ", out)
	if out != 142913828922 {
		t.Fail()
	}
}

func BenchmarkProb010(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob010()
	}
}
