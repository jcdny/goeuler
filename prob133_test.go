package euler

import (
	"testing"
)

// Find the sum of all the primes below one-hundred thousand that will never be a factor of R(10^n).
func prob133() int {
	sum := 0
	pv, _ := PrimeSieve(100001)

	for _, p := range pv {
		d := RepunitDivisor(p)
		// the trick here is that for p to divide R(10^n) 
		// A(p) has to be of form 2^m*5^k since (p-1) % A(p) == 0
		// eg: 11: 2, 17: 16, 41: 5, 73: 8
		if d > 0 {
			for d&1 == 0 {
				d /= 2
			}
			for d%5 == 0 {
				d /= 5
			}
		}
		if d != 1 {
			sum += p
		} else {
			//log.Print(p)
		}

	}
	return sum
}

func TestProb133(t *testing.T) {
	out := prob133()
	Validate(t, 133, out)
}

func BenchmarkProb133(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob133()
	}
}
