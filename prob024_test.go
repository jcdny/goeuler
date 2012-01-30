package euler

import (
	"testing"
)

func prob024() string {
	return PermuteString("0123456789", 1000000)
}

func TestProb024(t *testing.T) {
	out := prob024()
	Validate(t, 24, out)
}

func BenchmarkProb024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob024()
	}
}
