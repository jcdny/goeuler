package euler

import (
	"testing"
)

// What is the side length of the square spiral for which the ratio of
// primes along both diagonals first falls below 10%?

func prob058() int {
	pv, _ := PrimeSieve(100000)

	ntot := 1
	np := 1 // Primes - this is 1 because '3' below is sieved
	for n := 3; ; n += 2 {
		ntot += 4
	NOTPRIME:
		for k := 1; k < 4; k++ {
			diag := n*n - (n-1)*k
			// this is really slow but gave the right answer
			// d, _, _ := FermatFactor(uint64(diag))
			// turns out trial division is a lot faster...
			for _, p := range pv {
				if p > n {
					break
				}
				if diag%p == 0 {
					continue NOTPRIME
				}
			}
			np++
		}
		if np*10 < ntot {
			return n
		}
	}

	return 0
}

func TestProb058(t *testing.T) {
	out := prob058()
	Validate(t, 58, out)
}

func BenchmarkProb058(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob058()
	}
}
