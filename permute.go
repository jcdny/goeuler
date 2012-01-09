package euler

// PermuteString returns the n-th permutation of s.
// s is not sorted so assumes s is in the order you
// want the permutations sorted by.
func PermuteString(s string, n uint64) string {
	out := ""
	base := uint64(1)
	for i := 2; i < len(s); i++ {
		if base*uint64(i) < base {
			return "**OVERFLOW**"
		}
		base *= uint64(i)
	}
	// if n > base we just cycle.
	n = (n - 1) % (base * uint64(len(s)))

	for {
		out += string(s[n/base])
		// log.Print(n/base, n, base, " s ", s, " out ", out)
		if len(s) > 2 {
			s = s[0:n/base] + s[n/base+1:]
			n -= (n / base) * base
			// log.Print(base, len(s))
			base /= uint64(len(s))
		} else {
			out += string(s[(n/base+1)%2])
			break
		}
	}

	return out
}

// PermuteList takes a list of states and generates the permutations...
func PermuteList(N []int) [][]int {

	n := len(N)
	cnt := make([]int, n)

	nperm := 1
	for i := range N {
		nperm *= N[i]
	}

	out := make([][]int, nperm)
	buf := make([]int, nperm*n)
	for c, i := 0, 0; i < nperm; i++ {
		out[i] = buf[i*n : (i+1)*n]
		for j := 0; j < n; j++ {
			if c > 0 {
				cnt[j] += 1
				c--
			}
			if cnt[j] == N[j] {
				cnt[j] = 0
				c = 1
			}
		}
		c = 1
		copy(out[i], cnt)
	}

	return out
}
