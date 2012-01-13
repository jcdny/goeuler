package euler

import (
	"testing"
)

func prob035() int {
	np := 4
	pv, cv := PrimeSieve(1000000)
SKIP:
	for _, p := range pv[4:] {
		pp := p
		m := 1
		n := -1
		for pp > 0 {
			m *= 10
			n++
			if (pp%10)&1 == 0 {
				continue SKIP
			}
			pp /= 10
		}
		m /= 10
		//log.Print(p, m, n)
		pp = p
		var i int
		for i = 0; i < n; i++ {
			pp = (pp%10)*m + pp/10
			if cv[pp] {
				break
			}
		}
		if i == n {
			np++
		}
	}

	return np
}

func TestProb035(t *testing.T) {
	out := prob035()
	t.Log("Problem 035 ", out)
	if out != 55 {
		t.Fail()
	}
}

func BenchmarkProb035(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob035()
	}
}
