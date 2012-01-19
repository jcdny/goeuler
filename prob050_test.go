package euler

import (
	"testing"
)

func prob050() int {
	pv, cv := PrimeSieve(1000000)
	end := len(pv) / 20
	sum := make([]int, end)
	pmax := 0
	copy(sum, pv[:end])
	for mp := 1; mp < len(pv); mp++ {
		for add := 0; add < end-mp; add++ {
			//log.Print(mp, add, sum[add], pv[add+mp])
			sum[add] += pv[add+mp]
			if sum[add] >= len(cv) {
				end = add + mp + 1
				break
			} else if !cv[sum[add]] {
				pmax = sum[add]
			}
		}
	}

	return pmax
}

func TestProb050(t *testing.T) {
	out := prob050()
	t.Log("Problem 050 ", out)
	if out != 997651 {
		t.Fail()
	}
}

func BenchmarkProb050(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob050()
	}
}
