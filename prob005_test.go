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
	t.Log("Problem 005 ", out)
	if out != 232792560 {
		t.Fail()
	}
}

func BenchmarkProb005(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob005()
	}
}
