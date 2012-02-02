package euler

import (
	"log"
	"math/big"
	"testing"
)

func prob138() string {
	p := big.NewInt(17)
	c := big.NewInt(305)
	sum := big.NewInt(322) // 17+305

	t := big.NewInt(0)
	one := big.NewInt(1)
	for i := 2; i < 12; i++ {
		t.Set(c)
		c.Mul(c, c)
		c.Div(c, p)
		c.Add(c, one)
		p.Set(t)
		sum.Add(sum, c)
	}

	return sum.String()
}

func prob138explore() {
	// this is the really brute force version, it overflows at the
	// seventh term but it's clear from the ratios of successive terms
	// that it converges to L_n+1 = L_n * L_n/L_n-1 which also turns
	// out to be 8*phi+5 or 17.9442719...
	n := int64(0)
	suml := n
	l := n
	rat := make([]float64, 8)
	for a, nl := int64(3), 0; nl < 7; a++ {
		a2b2 := 5*a*a - 4*a + 1
		a2b2p := 5*a*a + 4*a + 1

		for n*n < a2b2 {
			n++
		}
		if a2b2 != n*n {
			for n*n < a2b2p {
				n++
			}
		}
		if a2b2 == n*n || a2b2p == n*n {
			if nl > 0 {
				rat[nl] = float64(n) / float64(l)
			}
			nl++
			l = n
			suml += n
			log.Print(a, n)
		}
	}
	for i := 2; rat[i] > 0; i++ {
		log.Print(rat[i] - rat[i-1])
	}

	log.Print(rat)
}

func TestProb138(t *testing.T) {
	out := prob138()
	Validate(t, 138, out)
}

func BenchmarkProb138(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob138()
	}
}
