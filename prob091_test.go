package euler

import (
	"testing"
)

// Find the number of right angle triangles in the quadrant.
func prob091() int {
	var tot int
	N := 50
	tot = 0
	// generate all the pairs of points in the 50x50 grid and check if
	// they are right triangles.  Double counts which could easily be
	// fixed because it's symmetric across diagonal but this is plenty
	// fast.
	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {
			for c := 0; c <= N; c++ {
				for d := 0; d <= N; d++ {
					if (a == c && d == b) ||
						(a == b && b == 0) ||
						(c == d && d == 0) {
						continue
					}
					s1 := (a*a + b*b)
					s2 := (c*c + d*d)
					s3 := (c-a)*(c-a) + (d-b)*(d-b)
					if s1 == s2+s3 || s2 == s1+s3 || s3 == s1+s2 {
						tot++
					}
				}
			}
		}
	}
	tot /= 2 // correct for double counting

	return tot
}

func TestProb091(t *testing.T) {
	out := prob091()
	Validate(t, 91, out)
}

func BenchmarkProb091(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob091()
	}
}
