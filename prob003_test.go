package euler

import (
	"testing"
)

func prob003() uint64 {
	i := uint64(600851475143)
	return MaxFactor(i)
}

func TestProb003(t *testing.T) {
	out := prob003()
	Validate(t, 3, out)
}

func BenchmarkProb003(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob003()
	}
}
