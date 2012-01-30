package euler

import (
	"testing"
	"strconv"
)

func prob004() int {
	mp := 0
	// Start from the top, count down
	for i := 999; i > 99; i-- {
		if i*(i-1) < mp {
			// if we could not find anything bigger than we already
			// have it's time to bail out
			break
		}
		for j := i - 1; j > 99; j-- {
			if i*j > mp && i*j%10 == 9 && Palindrome(strconv.Itoa(i*j)) {
				mp = i * j
			}
		}
	}
	return mp
}

func TestProb004(t *testing.T) {
	out := prob004()
	Validate(t, 4, out)
}

func BenchmarkProb004(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob004()
	}
}
