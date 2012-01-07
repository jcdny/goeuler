package euler

import (
	"testing"
	"log"
)

func prob005() int {
	lcm := 2
	for i := 3; i < 21; i++ {
		lcm = Lcm(lcm, i)
	}

	return lcm
}

func TestProb005(t *testing.T) {
	if prob005() != 232792560 {
		log.Print(prob005())
		t.Error("Prob005 failed")
	}
}

func BenchmarkProb005(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob005()
	}
}
