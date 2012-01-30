package euler

import (
	"testing"
)

func prob047() int {
	N := 4
	pv, _, fv := PrimeSieveFactors(200000)
	for c := N; c < len(fv); c++ {
		var cp int
		for cp = 0; cp < N; cp++ {
			if len(fv[c-cp]) != N {
				break
			}
		}
		if cp == N {
			// have N numbers with N factors; check they are unique
			ff := make(map[int]struct{}, N*N)
			for cp = 0; cp < N; cp++ {
				t := c - cp
				for _, f := range fv[t] {
					pp := 1
					p := pv[f]
					for t%p == 0 {
						t /= p
						pp *= p
					}
					ff[pp] = struct{}{}
				}
			}
			if len(ff) == N*N {
				return (c - N + 1)
			}
		}
	}

	return 0
}

func TestProb047(t *testing.T) {
	out := prob047()
	Validate(t, 47, out)
}

func BenchmarkProb047(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob047()
	}
}
