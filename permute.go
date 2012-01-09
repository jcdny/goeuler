package euler

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
