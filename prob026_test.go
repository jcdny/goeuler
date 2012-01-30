package euler

import (
	"testing"
)

func prob026() int {
	n := 1000

	mc := 0 // max cycle length
	mi := 0 // i for max cycle

	rem := make([]int, n) // remainder
	nd := 1               // number of digits

	for d := 1; d < n; d++ {
		// originally I zeroed rem[] each time through but 
		// it's faster to simply track the number of digits 
		// handled from the start of a given number and not zero it.
		// credit goes to the Project Euler forums for this idea
		snd := nd + 1
		r := 1
		// Loop until we see a remainder we already saw or have a full
		// expansion (r=0)
		for r != 0 && rem[r] < snd {
			rem[r] = nd
			for r < d {
				r *= 10
				nd++
			}
			r %= d
		}

		// if we have a remainder we have found a cycle
		// record it's length if it's now the longest.
		if r != 0 && nd-rem[r] > mc {
			mc = nd - rem[r]
			mi = d
		}
	}

	return mi
}

func TestProb026(t *testing.T) {
	out := prob026()
	Validate(t, 26, out)
}

func BenchmarkProb026(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob026()
	}
}
