package euler

import (
	"log"
	"math/big"
	"testing"
)

// This is basically long division but all we care about is
// having 0 remainder so we just loop adding digits until we
// end up with 0 remainder.
func RepunitDivisor(n int) int {
	if n&1 == 0 || n%5 == 0 {
		return 0
	}
	u := 1
	ru := 1

	for u != 0 {
		for n > u {
			u = u*10 + 1
			ru++
		}
		u = u % n
	}
	return ru
}

// The fast version takes an N and computes the first repunit it can evenly divide.
// Also we know A(n) < n so start at 1mm.
func prob129() int {
	for n := 1000000; n < 10000171; n++ {
		if RepunitDivisor(n) > 1000000 {
			return n
		}
	}
	return 1
}

// Very brute force version takes about 10m, worked though.
func prob129brute() int {
	//pv, cv := PrimeSieve(1200000)
	ru := big.NewInt(0)
	div := big.NewInt(0)
	rem := big.NewInt(0)
	quo := big.NewInt(0)
	one := big.NewInt(1)
	ten := big.NewInt(10)

	mnru := int64(1)
	for n := int64(1000000); ; n++ {
		// Initially assumed it would be a prime > 1m and got 10000171
		// which was wrong...

		//for _, p := range pv {
		//n := int64(p)
		if n%2 == 0 || n%5 == 0 {
			continue
		}
		nru := int64(1)
		ru.SetInt64(1)
		for nru < n {
			div.SetInt64(n)
			quo.QuoRem(ru, div, rem)
			if rem.BitLen() == 0 {
				if true || nru > mnru {
					log.Print(n, nru, n-int64(nru))
					if nru > 999999 {
						return int(nru)
					}
					mnru = nru
				}
				break
			}
			ru.Mul(ru, ten)
			ru.Add(ru, one)
			nru++
		}
	}

	return 1
}

// Just looking at the early A(n) to see if there is anything interesting.
func prob129explore() int {
NEXT:
	for n := int64(3); n < 1000; n++ {
		if n%2 == 0 || n%5 == 0 {
			continue
		}
		nru := 2
		ru := int64(11)
		for ru < 1111111111111111111 {
			if ru%n == 0 {
				log.Print(n, nru)
				continue NEXT
			}
			ru = ru*10 + 1
			nru++
		}
		log.Print(n, "missed")
	}
	return 1
}

func TestProb129(t *testing.T) {
	out := prob129()
	Validate(t, 129, out)
}

func BenchmarkProb129(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob129()
	}
}
