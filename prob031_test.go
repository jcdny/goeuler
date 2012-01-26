package euler

import (
	"testing"
)

var NWK = make(map[int]int64, 10000)

func W(a int, c []int) int64 {
	k := a*10 + len(c)
	if n, ok := NWK[k]; ok {
		return n
	}

	var n int64
	if len(c) == 1 || a == 0 {
		return 1
	} else if c[0] > a {
		n = W(a, c[1:])
	} else {
		n = W(a-c[0], c) + W(a, c[1:])
	}
	NWK[k] = n
	return n
}

func prob031() int64 {
	c := []int{200, 100, 50, 20, 10, 5, 2, 1}
	a := 200

	return W(a, c)
}

func TestProb031(t *testing.T) {
	out := prob031()
	t.Log("Problem 031 ", out)
	if out != 73682 {
		t.Fail()
	}
}

func BenchmarkProb031(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// have to wipe this out or we are getting the memoized answer in 180 ns :)
		NWK = make(map[int]int64, 10000)
		prob031()
	}
}
