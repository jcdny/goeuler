package euler

import (
	"testing"
)

func prob027() int {
	pv, cv := PrimeSieve(100000)
	mn := 1
	mp := 0
	for _, b := range pv {
		if b > 999 {
			break
		}
		for a := -999; a < 1000; a++ {
			if a == -1 || a+b <= 1 || cv[a+b+1] {
				continue
			}
			n := 0
			for n*n+n*a+b > 0 && !cv[n*n+n*a+b] {
				n++
			}
			if n > mn {
				mn = n
				mp = a * b
			}
		}
	}

	return mp
}

func TestProb027(t *testing.T) {
	out := prob027()
	Validate(t, 27, out)
}

func BenchmarkProb027(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob027()
	}
}
