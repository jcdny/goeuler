package euler

import (
	"testing"
)

func prob039brute() int {
	P := 1000
	ns := make([]int, P+1)

	// Mostly brute force though with some thought about bounds on a,b,c
	for c := 5; c < P/2; c++ {
		c2 := c * c
		for a := 3; a < P/2; a++ {
			c2a2 := c2 - a*a
			for b := a + 1; a+b+c <= P && b*b <= c2a2; b++ {
				if c2a2-b*b == 0 {
					p := a + b + c
					ns[p]++
				}
			}
		}
	}

	maxs := 0
	maxp := 0
	for p, n := range ns {
		if n > maxs {
			maxs = n
			maxp = p
		}
	}

	return maxp
}

func prob039() int {
	P := 1000
	ns := make([]int, P+1)
	// Find the Pythagorean Primitive Triples with perimeter <= 1000
	ptv := PythagoreanTriples(1000, func(x PTriple) bool { return x.A+x.B+x.B > P })
	for _, pt := range ptv {
		p := pt.A + pt.B + pt.C
		// now count primitive + non-primitive triangles by perimeter
		for n := 1; n*p <= P; n++ {
			ns[n*p]++
		}
	}

	maxs := 0
	maxp := 0
	for p, n := range ns {
		if n > maxs {
			maxs = n
			maxp = p
		}
	}

	return maxp
}

func TestProb039(t *testing.T) {
	out := prob039()
	Validate(t, 39, out)
}

func BenchmarkProb039(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob039()
	}
}

func BenchmarkProb039brute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob039brute()
	}
}
