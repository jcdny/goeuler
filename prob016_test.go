package euler

import (
	"math/big"
	"testing"
)

func prob016() int {
	z := big.NewInt(1)
	z.Lsh(z, 1000) // 2^1000
	return SumDigits(z)
}

func TestProb016(t *testing.T) {
	out := prob016()
	Validate(t, 16, out)
}

func BenchmarkProb016(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob016()
	}
}
