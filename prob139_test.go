package euler

import (
	"testing"
)

func prob139() int {
	P := 100000000
	// Find the Pythagorean Primitive Triples with perimeter <= P
	ptv := PythagoreanTriples(100000, func(x PTriple) bool { return x.A+x.B+x.B > P })
	n := 0
	for _, pt := range ptv {
		hole := (pt.C*pt.C - 2*pt.A*pt.B)
		// Its tileable iff 
		if pt.C*pt.C%hole == 0 {
			n += P / (pt.A + pt.B + pt.C)
			//log.Print(kpt, hole, P/(pt.A+pt.B+pt.C), pt)
		}
	}

	return n
}

func TestProb139(t *testing.T) {
	out := prob139()
	Validate(t, 139, out)
}

func BenchmarkProb139(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob139()
	}
}
