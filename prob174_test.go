package euler

import (
	"testing"
)

func prob174() int {
	N := int64(1000000)
	// doing this with a map[int]int is 30x slower.
	L := make([]int, N+1)
	// biggest n we can tile has side length N/4 + 1
	max := N/4 + 1
	for n := int64(3); n <= max; n++ {
		n2 := n * n
		for m := n - 2; m > 0; m -= 2 {
			// t is number of tiles required
			t := n2 - m*m
			if t > N {
				break
			}
			L[t]++
		}
	}

	// Now just count up 0 < L(t) < 11
	sum := 0
	for _, l := range L {
		if l > 0 && l < 11 {
			sum++
		}
	}
	return sum
}

func TestProb174(t *testing.T) {
	out := prob174()
	t.Log("Problem 174 ", out)
	if out != 209566 {
		t.Fail()
	}
}

func BenchmarkProb174(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob174()
	}
}
