package euler

import (
	"strconv"
	"testing"
)

// Find the smallest prime which, by replacing part of the number (not
// necessarily adjacent digits) with the same digit, is part of an
// eight prime value family.
func prob051() string {
	N := 1000000
	d := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	pv, _ := PrimeSieve(N - 1)
	ps := make(map[string]struct{}, len(pv))
	for _, p := range pv {
		ps[strconv.Itoa(p)] = struct{}{}
	}

	for _, np := range pv {
		p := strconv.Itoa(np)
		b := []byte(p)
		for _, ca := range p {
			n := 0
			for _, nd := range d {
				for j, cb := range p {
					if ca == cb {
						b[j] = nd
					}
				}
				if _, ok := ps[string(b)]; ok {
					n++
				}
			}
			if n == 8 {
				return p
			}
		}
	}

	return ""
}

func TestProb051(t *testing.T) {
	out := prob051()
	Validate(t, 51, out)
}

func BenchmarkProb051(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob051()
	}
}
