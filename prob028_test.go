package euler

import (
	"testing"
)

// There is also a simple closed form solution
// given that the #s along each diagonal are
//   1+9+25+49+...
//   1+7+21+43+...
//   1+5+17+37+...
//   1+3+13+31+...
//
// Alternatively:
//   4*(1+(9-3)+(25-6)+(49-9)+...)-3
//
// (the final -3 is to account for quadruple counting the 1)
// but we know the sum of the odd numbers squared is n(4n^2-1)/3
// and the sum of 0, -3, -6, -9,... is -3*n(n-1)/2
// we can reduce the above to:
//   4n(4n^2-1)/3 - 6n(n-1) - 3
// so for a 1001x1001 spiral, n=501 and we get the expected answer
//
// see https://oeis.org/A000447 for where all this comes from.

func prob028(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum += (2+4*i)*(2+4*i) - 12*i
	}

	return sum
}

func TestProb028(t *testing.T) {
	out := prob028(500)
	Validate(t, 28, out)
}

func BenchmarkProb028(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob028(500)
	}
}
