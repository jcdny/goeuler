package euler

import (
	"testing"
	"log"
	"sort"
	. "math"
)

// Junk for sorting composite sums
type csum struct {
	c     int
	delta int
}

type csumslice []csum

func (s csumslice) Len() int           { return len(s) }
func (s csumslice) Less(i, j int) bool { return s[i].delta > s[j].delta }
func (s csumslice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

const verb355 = false

// Return a list of composites which are larger than the sum of prime factors.
func composites(n int, pl []int, fv [][]int) []csum {
	cl := make([]csum, 0, len(pl)*2)

	for i := 4; i <= n; i++ {
		if len(fv[i]) == 0 {
			continue
		}

		sum := 0
		for _, pn := range fv[i] {
			sum += pl[pn]
		}

		if i > sum {
			cl = append(cl, csum{i, i - sum})
		}
	}

	return cl
}

func sumcoprime(fv [][]int, pl []int, cl []csum, drop []bool) (uint64, []int) {
	ptaken := make([]bool, len(pl))
	cused := make([]int, 0, len(pl))
	sum := uint64(1)

	// Go through and compute the sum of composites while updating
	// which primes have been used.
	for _, cs := range cl {
		if !drop[cs.c] {
			omit := false
			for _, pn := range fv[cs.c] {
				if ptaken[pn] {
					omit = true
					break
				}
			}
			if !omit {
				for _, pn := range fv[cs.c] {
					ptaken[pn] = true
				}
				sum += uint64(cs.c)
				cused = append(cused, cs.c)
			}
		}
	}

	// now sum unused prime numbers
	npused := 0
	for pn, used := range ptaken {
		if !used {
			sum += uint64(pl[pn])
			npused++
		}
	}

	return sum, cused
}

func prob355(n int) uint64 {
	rn := int(Ceil(Sqrt(float64(n))))
	pv, _, fv := PrimeSieveFactors(n + 1)

	if verb355 {
		log.Print("For ", n, " # of Primes: ", len(pv))
		//log.Print("Factors: ", fv)
	}

	// initially assume all primes are as big as can be
	// and make a list of them
	pl := make([]int, 0, len(pv))
	for _, p := range pv {
		pp := p
		if pp <= rn { // check this first to avoid overflow.
			for pp*p < n {
				pp *= p
			}
		}
		pl = append(pl, pp)
	}

	// generate a list of composites which are larger than the sum of their prime factors
	cl := composites(n, pl, fv)
	// sort in order of decreasing delta
	sort.Sort(csumslice(cl))

	// Generate the greedy set
	drop := make([]bool, len(fv))
	sum, cused := sumcoprime(fv, pl, cl, drop)

	if verb355 {
		log.Print("Greedy sum ", sum)
		if len(cl) < 20 {
			log.Print("Composites: ", cl)
		} else {
			log.Print("# of Composites: ", len(cl), " used ", len(cused))
		}
	}

	// Now iterate through the list and test dropping each composite
	// and any time we find one where the sum is greater drop it and
	// repeat.
	for {
		found := false
		for _, c := range cused {
			drop[c] = true
			newsum, newcused := sumcoprime(fv, pl, cl, drop)
			if newsum <= sum {
				drop[c] = false
			} else {
				found = true
				sum = newsum
				cused = newcused
				break
			}
		}

		if !found {
			break
		} else if verb355 {
			log.Print("Iteration sum ", sum)
		}
	}

	return sum
}

type test355 struct {
	n   int
	sum uint64
}

func TestProb355(t *testing.T) {
	for _, test := range []test355{{30, 193}, {100, 1356}, {200000, 0}} {
		// {1000000, 37717171223} takes almost a minute though
		sum := prob355(test.n)
		if test.n == 200000 {
			Validate(t, 355, sum)
		} else if sum != test.sum {
			t.Error("Prob355  for ", test, " got ", sum)
		}
	}
}

func BenchmarkProb355(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob355(200000)
	}
}
