package euler

import (
	"testing"
	. "math"
	"log"
)

// This is the fast version: only check those with 8+ dup digits.  I
// actually do a bunch of extra work in check since I actually know
// the count by knowing d,d1,d2 below so really only need to check
// primality but I had the machinery so just reused it.  This takes
// about 150ms
func prob111() int64 {
	N := int64(1e10)
	check, add, sum := accumulator(N)

	// create a table to generate all the permutations
	type permtable struct {
		b, m1, m2 int64
	}
	pt := make([]permtable, 0)
	base := (N - 1) / 9
	m1 := int64(1)
	for m1 < N {
		m2 := m1 * 10
		for m2 < N {
			pt = append(pt, permtable{base - m1 - m2, m1, m2})
			m2 *= 10
		}
		m1 *= 10
	}

	seen := make(map[int64]struct{}, 1000)
	for d := int64(0); d < 10; d++ {
		found := 0
		for d1 := int64(0); d1 < 10; d1++ {
			for d2 := int64(0); d2 < 10; d2++ {
				for _, perm := range pt {
					var n int64
					n = perm.b*d + perm.m1*d1 + perm.m2*d2
					if n < N/10 {
						// skip ones with leading 0's
						continue
					}
					if _, ok := seen[n]; !ok {
						if check(n) {
							found++
							add(n)
						}
					}
					seen[n] = struct{}{}
				}
			}
		}
	}

	return sum()
}

// This is the ugly brute force version - takes about 10 min to run.
//
// I realized though that the counts for duplicated digits are large
// you could go the other way and just generate all 10 digit numbers
// with 8+ duplicated digits.
//
// It's actually 9 for all but 0,2,and 8 which are 8.
func prob111brute() int64 {
	N := int64(100000000)
	check, add, sum := accumulator(N)
	for n := N/10 + 1; n < N; n += 2 {
		if n%3 == 0 || n%5 == 0 || n%7 == 0 {
			continue
		}
		if check(n) {
			add(n)
		}
	}

	log.Print(sum())
	return sum()
}

func accumulator(n int64) (func(int64) bool, func(int64), func() int64) {
	max := [10]int{}
	sum := [10]int64{}
	count := [10]int{}

	maxroot := int(Sqrt(float64(n)))
	pv32, _ := PrimeSieve(maxroot + 1)
	pv := make([]int64, 0, len(pv32))
	for _, p := range pv32 {
		pv = append(pv, int64(p))
	}

	r := int64(Sqrt(float64(n/10))) - 1

	check := func(nin int64) bool {
		for i := range count {
			count[i] = 0
		}
		n := nin
		for n > 0 {
			count[n%10]++
			n /= 10
		}

		for i, c := range count {
			if max[i] <= c {
				for r*r < nin {
					r++
				}
				for _, p := range pv {
					if nin%p == 0 {
						return false
					}
					if p > r {
						break
					}
				}
				return true
			}
		}

		return false
	}

	add := func(nin int64) {
		for i, c := range count {
			if max[i] < c {
				sum[i] = nin
				max[i] = c
			} else if max[i] == c {
				sum[i] += nin
			}
		}
	}

	tot := func() int64 {
		//log.Print(max)
		//log.Print(sum)
		tot := int64(0)
		for _, c := range sum {
			tot += c
		}
		return tot
	}

	return check, add, tot
}

func TestProb111(t *testing.T) {
	out := prob111()
	///out = prob111brute()
	Validate(t, 111, out)
}

func BenchmarkProb111(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob111()
	}
}
