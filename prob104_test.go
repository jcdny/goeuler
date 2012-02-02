package euler

import (
	"strconv"
	"testing"
)

// Given that Fk is the first Fibonacci number for which the first
// nine digits AND the last nine digits are 1-9 pandigital, find k.
func prob104() int {
	d := make([]int, 10)
	f0 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	f1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	fn := []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	t0 := int64(1)
	t1 := t0
	tn := t0
	for n := 3; ; n++ {
		// The MSB
		tn = t0 + t1

		// add the LSB and simultaneously check for
		// pandigital-ness
		c := 0
		pd := true
		for i := range f0 {
			fn[i] = f0[i] + f1[i] + c
			if fn[i] > 9 {
				fn[i] -= 10
				c = 1
			} else {
				c = 0
			}
			if pd {
				if fn[i] == 0 || d[fn[i]] == n {
					pd = false
				} else {
					d[fn[i]] = n
				}
			}

		}

		if pd {
			// we have a last 9 pandigital lets see if we have a
			// first 9 as well...
			s := strconv.FormatInt(tn, 10)
			if IsPandigital(s[:9]) {
				return n
			}
		} else {
			//log.Print(n, tn, fn)
		}
		if tn > 1e18-1 {
			// limit the MS to 18 digits; it ought to be plenty.  Its
			// about the limit of int64 as well.
			tn /= 10
			t1 /= 10
		}

		// swappy
		t0, t1 = t1, tn
		f0, f1, fn = f1, fn, f0
	}
	return 0
}

func TestProb104(t *testing.T) {
	out := prob104()
	Validate(t, 104, out)
}

func BenchmarkProb104(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob104()
	}
}
