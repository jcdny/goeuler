package euler

import (
	"testing"
)

// again with the pandigital...
// anyway it would be much faster to work from
// end to start with end a multiple of 17 and < 3 digits.
func prob043() int64 {
	sum := int64(0)
	pa := PermuteArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{})
	div := []int{17, 13, 11, 7, 5, 3, 2}
	for _, p := range pa {
		dp := p[8]*10 + p[9]
		var k int
		for k = 7; k > 0; k-- {
			dp += p[k] * 100
			if dp%div[7-k] != 0 {
				break
			}
			dp /= 10
		}
		if k == 0 {
			tpd := int64(0)
			for _, d := range p {
				tpd = tpd*10 + int64(d)
			}
			sum += tpd
		}
	}

	return sum
}

func TestProb043(t *testing.T) {
	out := prob043()
	Validate(t, 43, out)
}

func BenchmarkProb043(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob043()
	}
}
