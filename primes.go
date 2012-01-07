package euler

import (
	. "math"
)

func NPrimesApprox(n int) int {
	x := float64(n)
	// P. Dusart, "The kth prime is greater than k(ln k+ln ln k-1) for k> 2," 
	// Math. Comp., 68:225 (January 1999)
	// conservative bound on # of primes < n
	return int((x / Log(x)) * (1 + 1.2762/Log(x)))
}

// Return a list of primes less than n and a bool array of composite flags.
func PrimeSieve(n int) ([]int, []bool) {
	// composite flag since that way we can use default initialization
	c := make([]bool, n, n)
	end := int(Ceil(Sqrt(float64(n))))
	for i := 2; i < end; i++ {
		if !c[i] {
			j := i << 1
			for j < n {
				c[j] = true
				j += i
			}
		}
	}

	p := make([]int, 0, NPrimesApprox(n))
	for i := 2; i < len(c); i++ {
		if !c[i] {
			p = append(p, i)
		}
	}
	return p, c
}