package euler

import (
	"testing"
)

func prob075() int {
	P := 1500000
	np := make([]int, P+1)
	// Find the Pythagorean Primitive Triples with perimeter <= P
	ptv := PythagoreanTriples(1000000, func(x PTriple) bool { return x.A+x.B+x.B > P })
	for _, pt := range ptv {
		// Its tileable iff 
		p := (pt.A + pt.B + pt.C)
		for n := 1; n*p <= P; n++ {
			np[n*p]++
		}
	}

	count := 0
	for _, n := range np {
		if n == 1 {
			count++
		}
	}

	return count
}

func TestProb075(t *testing.T) {
	out := prob075()
	Validate(t, 75, out)
}

func BenchmarkProb075(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob075()
	}
}
