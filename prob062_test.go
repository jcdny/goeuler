package euler

import (
	"testing"
	"strconv"
	"sort"
	"math"
)

func prob062() int64 {
	N := 10000
	mh := make(map[string]int, N)
	mn := make(map[string]int64, N)
	for n := 2; n < N; n++ {
		// compute a minhash of the cube
		cube := int64(n) * int64(n) * int64(n)
		b := []byte(strconv.Itoa64(cube))
		sort.Sort(bsort(b))
		key := string(b)
		mh[key]++
		if mh[key] == 1 {
			// keep track of the min cube per minhash
			mn[key] = cube
		}
	}

	// find the min cube with 5 permutations
	min := int64(math.MaxInt64)
	for k, v := range mn {
		if mh[k] == 5 {
			if v < min {
				min = v
			}
		}
	}

	return min
}

func TestProb062(t *testing.T) {
	out := prob062()
	Validate(t, 62, out)
}

func BenchmarkProb062(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob062()
	}
}
