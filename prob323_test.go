package euler

import (
	"testing"
	. "math"
)

// This is basically markov chain calculation.
func prob323() float64 {
	// compute the binomial dist
	binom := make([][]float64, 33)
	binom[0] = []float64{1}
	for n := 1; n < 33; n++ {
		w := make([]float64, 0, n+1)
		t := 0.0
		for _, e := range binom[n-1] {
			w = append(w, e+t)
			t = e
		}
		w = append(w, t)
		binom[n] = w
	}

	// normalize
	for n := range binom {
		d := float64(int64(1) << uint(n))
		for k := range binom[n] {
			binom[n][k] /= d
		}
	}

	// start state n=0, all in state 32 zeros
	xn := make([]float64, 33)
	xn[32] = 1.0
	n1 := make([]float64, 0, 100)

	// evolve state until convergence
	for n, r := 1.0, 1.0; r*n > 1e-11; n++ {
		w := make([]float64, 33)

		for k := 1; k < len(xn); k++ {
			for i, e := range binom[k] {
				w[i] += xn[k] * e
			}
		}

		r -= w[0]
		n1 = append(n1, w[0])
		copy(xn, w)
	}

	avg := 0.0
	for j, w := range n1 {
		avg += float64(j+1) * w
	}

	return avg
}

func TestProb323(t *testing.T) {
	out := prob323()
	t.Log("Problem 323 ", out)
	if Fabs(out-6.3551758451) > 1e-10 {
		t.Fail()
	}
}

func BenchmarkProb323(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob323()
	}
}
