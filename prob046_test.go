package euler

import (
	"testing"
)

func prob046() int {
	// just sieve the composites and check for the first one 
	// which is unmatched. ~ 500usec
	N := 100
	pv, cv := PrimeSieve(2 * N * N)
	ncv := len(cv)
	for _, p := range pv {
		for s := 1; s < N; s++ {
			ps2 := p + s*s*2
			if ps2 > ncv {
				break
			}
			cv[ps2] = false
		}
	}
	for i := 33; i < len(cv); i += 2 {
		if cv[i] {
			return i
		}
	}

	return 0
}

func TestProb046(t *testing.T) {
	out := prob046()
	t.Log("Problem 046 ", out)
	if out != 5777 {
		t.Fail()
	}
}

func BenchmarkProb046(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob046()
	}
}
