package euler

import (
	"testing"
)

func prob033() int {
	dp := 1
	np := 1
	for d := 12; d < 100; d++ {
		d1 := d / 10
		d2 := d % 10
		if d2 == 0 {
			continue
		}
		for n := 11; n < d; n++ {
			n1 := n / 10
			n2 := n % 10

			if (d1 == n1 && n2*d/d2 == n && n2*d%d2 == 0) ||
				(d1 == n2 && n1*d/d2 == n && n1*d%d2 == 0) ||
				(d2 == n1 && n2*d/d1 == n && n2*d%d1 == 0) ||
				(d2 == n2 && n1*d/d1 == n && n1*d%d1 == 0) {
				dp *= d
				np *= n
			}
		}
	}

	return dp / Gcd(np, dp)
}

func TestProb033(t *testing.T) {
	out := prob033()
	Validate(t, 33, out)
}

func BenchmarkProb033(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob033()
	}
}
