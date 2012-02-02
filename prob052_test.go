package euler

import (
	"sort"
	"strconv"
	"testing"
)

type bsort []byte

func (s bsort) Len() int           { return len(s) }
func (s bsort) Less(i, j int) bool { return s[i] < s[j] }
func (s bsort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func prob052() int {
	v := make([][]byte, 6, 6)
	m := 100
	u := 1000
	var i int
	for i = m + 2; ; i += 3 {
		// leading digit has to be 1, the number has to be divisible by
		// 3 (since the sum of the digits are a multiple of 3 and at least
		// one number is 3x), number of digits has to be same for x and 6x
		if i*6 >= u {
			u *= 10
			m *= 10
			i = m + 2
			continue
		}
		var j int
		for j = 0; j < 6; j++ {
			v[j] = []byte(strconv.Itoa(i * (j + 1)))
			sort.Sort(bsort(v[j]))
			if j > 0 && string(v[j]) != string(v[j-1]) {
				break
			}
		}
		if j == 6 {
			break
		}
	}
	return i
}

func TestProb052(t *testing.T) {
	out := prob052()
	Validate(t, 52, out)
}

func BenchmarkProb052(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob052()
	}
}
