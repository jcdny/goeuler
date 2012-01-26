package euler

import (
	"testing"
)

func prob069() int {
	// I think you can prove this has to be correct but in any case it
	// is.
	return 2 * 3 * 5 * 7 * 11 * 13 * 17
}

func TestProb069(t *testing.T) {
	out := prob069()
	t.Log("Problem 069 ", out)
	if out != 510510 {
		t.Fail()
	}
}

func BenchmarkProb069(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob069()
	}
}
