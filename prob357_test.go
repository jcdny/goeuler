package euler

import (
	"log"
	"testing"
)

const N = 100000000

// This took 6253 seconds - ugh.  There are 5761455 primes-1
// reject 5667266 after 1 round of FermatFactor, leaving ~ 94k but
// FermatFactor on 5.7mm numbers is not made of win.
func prob357bad() int {
	pv, cv := PrimeSieve(N)
	reject := 0
	for _, p := range pv {
		n := uint64(p - 1)
		lo, hi, n2 := FermatFactor(n)
		if cv[lo+n/lo] || cv[hi+n/hi] ||
			cv[2+n/2] ||
			(lo != 1 && cv[2*lo+n/(2*lo)]) {
			reject++
		} else {
			log.Print("accept ", n, lo, hi, n2)
		}
	}

	log.Print(len(pv), reject)

	return 1
}

func prob357() uint64 {
	pv, cv := PrimeSieve(N + 1)
	// the seived candidates.  with the exception of 1, 2 they must
	// composite so initially copy cv
	sv := make([]bool, len(cv))
	copy(sv, cv)

	// now remove any that are not a prime - 1
	for i := range sv[:len(sv)-1] {
		sv[i] = sv[i] && !cv[i+1]
	}
	// as we sieve again check the candidate number
	// p+n is prime where n = N/p
	for _, p := range pv {
		pp := p
		for n := 1; pp+n < N; n, pp = n+1, pp+p {
			sv[pp] = sv[pp] && !cv[p+n]
		}
	}
	// here we have ~32k candidates, just factor them and
	// check...
	sum := uint64(3) // 1 and 2 are special
	for i, s := range sv[:len(sv)-1] {
		if s {
			us := uint64(i)
			allp := true
			for _, d := range Divisors(us, pv, cv) {
				if cv[d+us/d] {
					allp = false
					break
				}
			}
			if allp {
				sum += uint64(i)
			}
		}
	}

	return sum
}

func TestProb357(t *testing.T) {
	out := prob357()
	Validate(t, 357, out)
}

func BenchmarkProb357(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob357()
	}
}
