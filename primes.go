package euler

import (
	. "math"
	"log"
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
func PrimeSieveFactors(n int) ([]int, []bool, [][]int) {
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

// Return a list of primes <= n, prime factors, divisors, and sum of divisors.
// len(d) is sigma_0(i) https://oeis.org/A000005
// dsum(i) is sigma_1(i) https://oeis.org/A000203
func PrimeSieveDivisors(n int) ([]int, [][]int, [][]int, []int) {
	// composite flag since that way we can use default initialization
	if n < 2 {
		log.Panic("n < 2")
	}
	n++
	c := make([]bool, n, n)
	f := make([][]int, n, n)
	d := make([][]int, n, n)
	dsum := make([]int, n, n)

	d[1] = append(d[1], 1)
	dsum[1] = 1

	pn := -1
	for i := 2; i <= n/2; i++ {
		if !c[i] {
			pn++
			for j, m := i<<1, 2; j < n; j, m = j+i, m+1 {
				// seive the multiples of a given prime and compute divisors.
				c[j] = true
				f[j] = append(f[j], pn)
				if len(d[j]) == 0 {
					// TODO size sensibly
					d[j] = make([]int, 1, int(Sqrt(float64(j)))+1)
					d[j][0] = 1
					dsum[j] = 1
				}

				ld := len(d[j])
				for ii := i; j%ii == 0; ii = ii * i {
					for l := 0; l < ld; l++ {
						//log.Print(i, j, ii, len(d[j]))
						d[j] = append(d[j], d[j][l]*ii)
						dsum[j] += d[j][l] * ii
					}
				}
			}
		}
	}

	p := make([]int, 0, NPrimesApprox(n))
	buf := make([]int, 0, 2*NPrimesApprox(n))
	for i := 2; i < len(c); i++ {
		if !c[i] {
			p = append(p, i)
			if dsum[i] == 0 {
				dsum[i] = i + 1
				buf = append(buf, 1, i)
				d[i] = buf[len(buf)-2:]
			}
		}
	}

	return p, f, d, dsum
}
