package euler

import (
	"testing"
)

func prob007() int {
	p, _ := PrimeSieve(105000)
	return p[10000]
}

func TestProb007(t *testing.T) {
	out := prob007()
	t.Log("Problem 007 ", out)
	if out != 104743 {
		t.Fail()
	}
}

func BenchmarkProb007(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob007()
	}
}
