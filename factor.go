package euler

import (
	. "math"
	//"log"
)

// Returns factors of n, lo, hi,2^n.
// cf Algorithm C p. 387 TAoCP v2 3e
func FermatFactor(n uint64) (uint64, uint64, uint64) {
	if n < 1 {
		return 0, 0, 1
	}

	n2 := uint64(1)
	for n&1 == 0 { // take out factors of 2
		n2 <<= 1
		n >>= 1
	}

	if n == 1 {
		return 1, 1, n2
	}

	nf := int64(Floor(Sqrt(float64(n))))
	a := 2*nf + 1
	b := int64(1)
	r := int64(uint64(nf*nf) - n)
	for r != 0 {
		r += a
		a += 2
		r -= b
		b += 2
		for r > 0 {
			r -= b
			b += 2
		}
	}

	return uint64((a - b) / 2), uint64((a + b - 2) / 2), n2
}

func MaxFactor(n uint64) uint64 {
	lo, hi, to := FermatFactor(n)
	if lo == 1 {
		if hi == 1 {
			return to
		} else {
			return hi
		}
	} else {
		a := MaxFactor(lo)
		b := MaxFactor(hi)
		if a > b {
			return a
		} else {
			return b
		}
	}
	return 0
}

func Divisors(n uint64, pv []int, cv []bool) []uint64 {
	if n < 3 {
		return []uint64{}
	}

	lcv := uint64(len(cv))
	rn := int(Sqrt(float64(n)))
	pdiv := make([]uint64, 1, 64)
	pdiv[0] = 1
	for _, pi32 := range pv {
		p := uint64(pi32)
		pnd := 1
		if n == 1 {
			break
		} else if pi32 > rn || (n < lcv && !cv[n]) {
			// using the composite flag significantly speeds things up.
			// ~ 5x faster when we precompute sqrt(MAX(N)) primes.
			p = n
			pnd = 2
			n = 1
		} else {
			for n%p == 0 {
				pnd++
				n /= p
			}
		}
		if pnd > 1 {
			lp := len(pdiv)
			for k, pp := 0, p; k < pnd-1; k, pp = k+1, pp*p {
				for i := 0; i < lp; i++ {
					pdiv = append(pdiv, pdiv[i]*pp)
				}
			}
		}
	}

	return pdiv
}

func NDivisors(n uint64, pv []int, cv []bool) int {
	if n < 3 {
		return int(n)
	}

	lcv := uint64(len(cv))
	rn := int(Sqrt(float64(n)))
	nd := 1

	for _, pi32 := range pv {
		if n == 1 {
			break
		} else if pi32 > rn || (n < lcv && !cv[n]) {
			// using the composite flag significantly speeds things up.
			// ~ 5x faster when we precompute sqrt(MAX(N)) primes.
			nd *= 2
			break
		}
		p := uint64(pi32)
		pnd := 1
		for n%p == 0 {
			pnd++
			n /= p
		}
		nd *= pnd
	}

	//log.Print(on, nd)
	return nd
}
