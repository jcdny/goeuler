package euler

import (
	"math/big"
	"testing"
)

func prob020() int {
	z := big.NewInt(0)
	z = z.MulRange(1, 100)
	return SumDigits(z)
}

func TestProb020(t *testing.T) {
	out := prob020()
	Validate(t, 20, out)
}

func BenchmarkProb020(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob020()
	}
}
