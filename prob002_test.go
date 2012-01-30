package euler

import (
	"testing"
)

func prob002() int {
	n, np := 1, 1
	sum := 0
	for n < 4e6 {
		np, n = n, n+np
		if n&1 == 0 {
			// n&1 faster than n%2
			sum += n
		}
	}

	return sum
}

func TestProb002(t *testing.T) {
	out := prob002()
	Validate(t, 2, out)
}

func BenchmarkProb002(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob002()
	}
}
