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

// Permute takes a slice of ints and returns the permutations of v.
// Transliteration of Algorithm P (plain changes) of TAoCP v4a 7.2.1.2
// modified to use 0 based array...
func PermuteArray(v []int, buf []int) [][]int {
	n := len(v)

	c := make([]int, len(v))
	o := make([]int, len(v))

	// n! permutations, also init o to 1
	nperm := 1
	for i := range v {
		o[i] = 1
		nperm *= (i + 1)
	}

	out := make([][]int, nperm)

	if cap(buf) < nperm*n {
		buf = make([]int, 0, nperm*n)
	} else {
		buf = buf[:0]
	}

	t := make([]int, len(v))

	i, j, s, q := -1, 0, 0, 0
	copy(t, v)
P2:
	buf = append(buf, t...)
	i++
	out[i] = buf[len(buf)-n : len(buf)]
	// P3:
	j = n - 1
	s = 0
P4:
	q = c[j] + o[j]
	if q < 0 {
		goto P7
	}
	if q > j {
		goto P6
	}
	// P5:
	t[j-c[j]+s], t[j-q+s] = t[j-q+s], t[j-c[j]+s]
	c[j] = q
	goto P2
P6:
	if j == 0 {
		return out
	}
	s++
P7:
	o[j] = -o[j]
	j--
	goto P4

	return out
}
