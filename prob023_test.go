package euler

import (
	"testing"
)

func prob023() int {
	_, _, _, dsumv := PrimeSieveDivisors(28124)

	// find the abundant numbers and convert to sum of proper divisors
	an := make([]int, 0, 28123)
	for i, ds := range dsumv {
		if ds-i > i {
			//log.Print("abundant ", i)
			an = append(an, i)
			dsumv[i] -= i
		} else {
			dsumv[i] = 0
		}
	}
	sum := 0
	for i := 1; i < 28123; i++ {
		var a int
		for _, a = range an {
			if a > i/2 || dsumv[i-a] > 0 {
				break
			}
		}
		if a > i/2 {
			sum += i
		}
	}
	return sum
}

func TestProb023(t *testing.T) {
	out := prob023()
	t.Log("Problem 023 ", out)
	if out != 4179871 {
		t.Fail()
	}
}

func BenchmarkProb023(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob023()
	}
}
