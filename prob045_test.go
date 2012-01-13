package euler

import (
	"testing"
)

func prob045() uint64 {
	np := uint64(166)
	nh := uint64(143)
	p := (3*np*np - np) / 2
	h := (2*nh*nh - nh) / 2
	for p != h {
		for p < h {
			np++
			p = (3*np*np - np) / 2
		}
		for h < p {
			nh++
			h = nh * (2*nh - 1)
		}
		//log.Print(np, nh)
	}

	return p
}

func TestProb045(t *testing.T) {
	out := prob045()
	t.Log("Problem 045 ", out)
	if out != 1533776805 {
		t.Fail()
	}
}

func BenchmarkProb045(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob045()
	}
}
