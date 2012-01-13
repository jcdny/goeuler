package euler

import (
	"testing"
)

func prob034() int {
	var fact = [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	sum := 0
	end := 6*fact[9] + 3
	for i := 3; i < end; i++ {
		j := i
		fs := 0
		for j > 0 {
			fs += fact[j%10]
			j /= 10
		}
		if fs == i {
			sum += i
		}
	}
	return sum
}

func TestProb034(t *testing.T) {
	out := prob034()
	t.Log("Problem 034 ", out)
	if out != 40730 {
		t.Fail()
	}
}

func BenchmarkProb034(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob034()
	}
}
