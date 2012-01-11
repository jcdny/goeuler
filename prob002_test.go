package euler

import (
	"testing"
)

func prob002() int {
	n, np := 1, 1
	sum := 0
	for n < 4e6 {
		np, n = n, n+np
		if n&1 == 0 { // n&1 faster than n%2, doh 6g
			sum += n
		}
	}

	return sum
}

func TestProb002(t *testing.T) {
	out := prob002()
	t.Log("Problem 002 ", out)
	if out != 4613732 {
		t.Fail()
	}
}

func BenchmarkProb002(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob002()
	}
}
