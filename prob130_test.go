package euler

import (
	"testing"
)

// Find the sum of the first forty prime factors of R(10^9).
func prob130() int {
	sum := 0
	i := 0
	_, cv := PrimeSieve(20000)
	for n, c := range cv {
		if c && n > 10 {
			ap := RepunitDivisor(n)
			if ap > 0 && (n-1)%ap == 0 {
				i++
				sum += n
			}
			if i == 25 {
				break
			}
		}
	}

	return sum
}

func TestProb130(t *testing.T) {
	out := prob130()
	Validate(t, 130, out)
}

func BenchmarkProb130(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob130()
	}
}
