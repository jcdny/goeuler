package euler

import (
	"testing"
	"log"
)

// Connectedness of a network.

// The Prime Minister's phone number is 524287. After how many
// successful calls, not counting misdials, will 99% of the users
// (including the PM) be a friend, or a friend of a friend etc., of
// the Prime Minister?

func LaggedFibGen(j, k, m int) func() int {
	if j >= k {
		return nil
	}

	// initialize sn
	sn := make([]int, k)
	for i := range sn {
		n := int64(i + 1)
		s := (100003 - 200003*n + 300007*n*n*n) % int64(m)
		sn[i] = int(s)
	}

	n := -1
	return func() int {
		n++
		i := n % k
		if n >= k {
			sn[i] = (sn[(n-j)%k] + sn[i]) % m
		}
		return sn[i]
	}
}

func prob186() int {
	root := func(n int, tree []int) int {
		for tree[n] != n {
			n = tree[n]
		}
		return n
	}

	// problem setup: 1e6 numbers, 90% connected to PM at 524287
	M := 1000000
	T := M - M/100
	pmr := 524287

	fg := LaggedFibGen(24, 55, M)

	// initially everyone is in a network rooted on themselves
	// with size 1
	nr := make([]int, M) // parent for each caller.
	nc := make([]int, M) // number of calls, correct for a given root
	for i := range nr {
		nr[i] = i
		nc[i] = 1
	}

	c := 0

	// make calls til the PM is connected enough
	for nc[pmr] < T {
		d1 := fg()
		d2 := fg()
		if d1 == d2 {
			continue
		} else {
			c++
		}

		// Join disjoint groups and accumulate tree size.
		// Joining on the min root speeds this up substatially
		// since the mean depth is then appreciably lower.
		dr1 := root(d1, nr)
		dr2 := root(d2, nr)
		if dr1 < dr2 {
			nc[dr1] += nc[dr2]
			nr[dr2] = dr1
		} else if dr2 < dr1 {
			nc[dr2] += nc[dr1]
			nr[dr1] = dr2
		} else {
			nr[d1] = dr1
			nr[d2] = dr1
		}

		pmr = root(pmr, nr)
		if false && nc[pmr] > 1 {
			log.Print(c, nc[pmr])
		}
	}

	return c
}

func TestProb186(t *testing.T) {
	out := prob186()
	Validate(t, 186, out)
}

func BenchmarkProb186(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob186()
	}
}
