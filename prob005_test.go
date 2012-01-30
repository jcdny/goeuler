package euler

import (
	"testing"
)

func prob005() int {
	lcm := 2
	for i := 3; i < 21; i++ {
		lcm = Lcm(lcm, i)
	}

	return lcm
}

func TestProb005(t *testing.T) {
	out := prob005()
	Validate(t, 5, out)
}

func BenchmarkProb005(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob005()
	}
}
