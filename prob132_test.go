package euler

import (
	"testing"
	"big"
)

// Find the sum of the first forty prime factors of R(10^9).
func prob132() int {
	// I did 133 before this one and reading the forum found
	// that the prime factors satisfy 10^(10^9) % p == 1
	sum := 0
	i := 0
	pv, _ := PrimeSieve(200001)
	x := big.NewInt(10)
	y := big.NewInt(1e9)
	z := big.NewInt(0)
	m := big.NewInt(0)
	for _, p := range pv[4:] {
		m := m.SetInt64(int64(p))
		z.Exp(x, y, m)
		if z.BitLen() == 1 {
			i++
			sum += p
		}
		if i == 40 {
			break
		}
	}
	return sum
}

func TestProb132(t *testing.T) {
	out := prob132()
	Validate(t, 132, out)
}

func BenchmarkProb132(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob132()
	}
}
