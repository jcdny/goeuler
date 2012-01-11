package euler

import (
	"testing"
)

func prob001() int {
	n := 0
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			n += i
		}
	}

	return n
}

func TestProb001(t *testing.T) {
	out := prob001()
	t.Log("Problem 001 ", out)
	if out != 233168 {
		t.Fail()
	}
}

func BenchmarkProb001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob001()
	}
}
