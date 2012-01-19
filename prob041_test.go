package euler

import (
	"testing"
	"strconv"
)

func prob041() int {
	// we know it can't be more than 7 digits since 
	// 8 and 9 digit pandigitals are divisible by 3
	// so just check them all.

	// It would be a lot faster to check the 5040 7 digit pandigitals
	// but I am sick of these problems so doing the simple version
	pv, _ := PrimeSieve(1000000)
	for j := len(pv) - 1; j > 0; j-- {
		s := strconv.Itoa(pv[j])
		if IsPandigitalN(s, len(s)) {
			return pv[j]
		}
	}

	return 0
}

func TestProb041(t *testing.T) {
	out := prob041()
	t.Log("Problem 041 ", out)
	if out != 7652413 {
		t.Fail()
	}
}

func BenchmarkProb041(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob041()
	}
}
