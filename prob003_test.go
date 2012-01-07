package euler

import (
	"testing"
)

func prob003() uint64 {
	i := uint64(600851475143)
	return MaxFactor(i)
}

func TestProb003(t *testing.T) {
	if prob003() != 6857 {
		t.Error("Prob003 failed")
	}
}

func BenchmarkProb003(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob003()
	}
}
