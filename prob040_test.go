package euler

import (
	"testing"
)

func prob040() int {
	prod := 1

	nd := 0 // number of digits
	wd := 1 // digit wanted	
	n := 1  // digits in the current number
	i := 0
	m := 10
	for {
		i++
		if i == m {
			m *= 10
			n++
		}
		nd += n
		if nd >= wd {
			ii := i
			for d := nd; d != wd; d-- {
				ii /= 10
			}
			prod *= ii % 10
			wd *= 10
			if wd > 1000000 {
				break
			}
		}
	}

	return prod
}

func TestProb040(t *testing.T) {
	out := prob040()
	t.Log("Problem 040 ", out)
	if out != 210 {
		t.Fail()
	}
}

func BenchmarkProb040(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob040()
	}
}
