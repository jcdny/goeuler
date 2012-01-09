package euler

import (
	"testing"
)

func prob028(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum += (2+4*i)*(2+4*i) - 12*i
	}

	return sum
}

func TestProb028(t *testing.T) {
	sum := prob028(500)
	if sum != 669171001 {
		t.Error("Prob028 failed ", sum)
	}
}

func BenchmarkProb028(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob028(500)
	}
}
