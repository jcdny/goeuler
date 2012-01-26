package euler

import (
	"testing"
)

func prob044() int {
	N := 4000 // check first 4k

	pn := make([]int, 0, N+1)
	pf := make([]bool, N*N*3/2)
	for p, i := 1, 4; p < len(pf); p, i = p+i, i+3 {
		pf[p] = true
		pn = append(pn, p)
	}
	for i := 1; i < len(pn); i++ {
		for j := i + 1; j < len(pn); j++ {
			if pn[i]+pn[j] < len(pf) {
				if pf[pn[i]+pn[j]] && pf[pn[j]-pn[i]] {
					return pn[j] - pn[i]
				}
			}
		}
	}

	return 0
}

func TestProb044(t *testing.T) {
	out := prob044()
	t.Log("Problem 044 ", out)
	if out != 5482660 {
		t.Fail()
	}
}

func BenchmarkProb044(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob044()
	}
}
