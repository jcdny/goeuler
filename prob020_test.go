package euler

import (
	"testing"
	"big"
)

func prob020() int {
	z := big.NewInt(0)
	z = z.MulRange(1, 100)
	return SumDigits(z)
}

func TestProb020(t *testing.T) {
	if prob020() != 648 {
		t.Error("Prob020 failed")
	}
}

func BenchmarkProb020(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob020()
	}
}
