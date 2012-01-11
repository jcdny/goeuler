package euler

import (
	"testing"
)

func prob030() int {
	sum := 0
	// 295278 is 2^5 + 5*9^5 which is the largest possible number which
	// could be the sum of it's digits^5
	for i := 2; i < 295278; i++ {
		j := i
		n := 0
		for j > 0 {
			d := j % 10
			n += d * d * d * d * d
			j /= 10
		}
		if n == i {
			sum += i
		}
	}
	return sum
}

func TestProb030(t *testing.T) {
	out := prob030()
	t.Log("Problem 030 ", out)
	if out != 443839 {
		t.Fail()
	}
}

func BenchmarkProb030(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob030()
	}
}
