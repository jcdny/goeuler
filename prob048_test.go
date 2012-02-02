package euler

import (
	"math/big"
	"testing"
)

const M = 1000

func prob048() string {
	zm := big.NewInt(1e10)
	sum := big.NewInt(1)
	for a := int64(2); a <= M; a++ {
		t := big.NewInt(int64(a))
		t.Exp(t, t, zm)
		sum.Add(sum, t)
	}

	sum.Mod(sum, zm)
	return sum.String()
}

func TestProb048(t *testing.T) {
	out := prob048()
	Validate(t, 48, out)
}

func BenchmarkProb048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob048()
	}
}
