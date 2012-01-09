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

// Return a list of primes less than n and a bool array of composite flags and the prime factors.
func PrimeSieveFactor(n int) ([]int, []bool, [][]int) {
	// composite flag since that way we can use default initialization
	c := make([]bool, n, n)
	f := make([][]int, n, n)
	// don't use end here since we need factors not just primes.
	// end := int(Ceil(Sqrt(float64(n))))
	pn := 0
	for i := 2; i <= n/2; i++ {
		if !c[i] {
			j := i << 1
			for j < n {
				//log.Print(i, j, pn)
				f[j] = append(f[j], pn)
				c[j] = true
				j += i
			}
			pn++
		}
	}

	p := make([]int, 0, NPrimesApprox(n))

	for i := 2; i < len(c); i++ {
		if !c[i] {
			p = append(p, i)
		}
	}
	return p, c, f
}

// Return a list of primes less than n, list of divisors, and sum of proper divisors
func PrimeSieveDiv(n int) ([]int, [][]int, []int) {
	// composite flag since that way we can use default initialization
	c := make([]bool, n, n)
	f := make([][]int, n, n)
	fp := make([][]int, n, n)
	d := make([][]int, n, n)
	dsum := make([]int, n, n)
	// don't use end here since we need factors not just primes.
	// end := int(Ceil(Sqrt(float64(n))))
	pn := 0
	for i := 2; i <= n/2; i++ {
		if !c[i] {
			j := i << 1
			m := 2
			for j < n {
				f[j] = append(f[j], pn)
				fp[j] = append(f[j], m)
				c[j] = true
				j += i
				m++
			}
			pn++
		}
	}

	p := make([]int, 0, NPrimesApprox(n))
	for i := 2; i < len(c); i++ {
		if !c[i] {
			p = append(p, i)
		}
	}

	// now compute proper divisors and their sum
	for i := range f {
		if len(f) == 0 {
			dsum[i] = i + 1
			d[i] = append(d[i], 1, i)
		} else {
			e := PermuteList(fp[i])
			d[i] = make([]int, len(e))
		}

	}
	return p, d, dsum
}
