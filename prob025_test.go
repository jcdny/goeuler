package euler

import (
	"testing"
	"big"
)

// This is the brute force solution
// Fk-1/Fk -> (1+sqrt(5))/2 as k -> \inf
// so easy enough to compute closed form but that
// just uses paper...
// Here is a nice writeup of the details...
// http://echorand.me/2010/02/02/eigen-values-vectors-and-fibonacci-sequence/
func prob025() int {
	nf := 12
	t := big.NewInt(0)
	fm := big.NewInt(89)
	f := big.NewInt(144)

	for {
		// log(10^999)/log(2) = 3318.6 so no need to test before then
		if f.BitLen() > 3318 {
			if len(f.String()) > 999 {
				break
			}
		}
		t.Set(f)
		f.Add(f, fm)
		fm.Set(t)
		nf++
	}

	return nf
}

func TestProb025(t *testing.T) {
	out := prob025()
	Validate(t, 25, out)
}

func BenchmarkProb025(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob025()
	}
}
