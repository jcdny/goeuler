package euler

import (
	"testing"
)

func prob037() int {
	pv, cv := PrimeSieve(1000000)
	cv[1] = true
	sum := 0
	for _, op := range pv[4:] {
		p := op
		m := 10
		for p /= 10; p > 0 && !cv[p] && !cv[op%m]; p /= 10 {
			m *= 10
		}
		if p > 0 {
			continue
		}
		sum += op
	}
	return sum
}

func TestProb037(t *testing.T) {
	out := prob037()
	Validate(t, 37, out)
}

func BenchmarkProb037(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob037()
	}
}
