package euler

import (
	. "math"
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