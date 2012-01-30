package euler

import (
	"testing"
	"strconv"
	"log"
)

func sumreverse(b, bo []byte) []byte {
	bo = bo[:0]
	n := len(b)
	c := uint8(0)
	for i, d := range b {
		nb := d + b[n-i-1] + c
		//log.Print("NB : ", nb, " = ", d, b[n-i-1], c)
		if nb > 9 {
			nb -= 10
			c = 1
		} else {
			c = 0
		}
		bo = append(bo, nb)
	}
	if c == 1 {
		bo = append(bo, c)
	}

	return bo
}

func prob055() int {
	M := 10000
	lychrel := 0
	w := make([]byte, 0, 5000)
	o := make([]byte, 0, 5000)
	for n := 1; n < M; n++ {
		w = w[:0]
		for _, c := range strconv.Itoa(n) {
			w = append(w, uint8(c)-'0')
		}

		var j int
		for j = 0; j < 50; j++ {
			o = sumreverse(w, o)
			if Palindrome(string(o)) {
				break
			}
			o, w = w, o
		}
		if j == 50 {
			lychrel++
		}
	}

	return lychrel
}

// This gives the right answer even though it overflows int64 since
// it only really overflows on lychrel numbers...
//
// It takes 27ms whereas the non memoized and correct version above takes only 16ms
func prob055bad() int {
	M := int64(10000)
	lychrel := 0

	orbit := make([]int64, 0, 51)
	nlorbit := make([]bool, M)
	lorbit := make([]bool, M)
	for n := int64(1); n < M; n++ {
		if nlorbit[n] {
			continue
		}
		if lorbit[n] {
			lychrel++
			continue
		}
		w := n
		orbit = orbit[:1]
		orbit[0] = w
		var j int
		for j = 0; j < 50; j++ {
			w += DigitReverse(w)
			if false && w < 0 {
				log.Print("overflow ", n, j)
			}
			if Palindrome(strconv.Itoa64(w)) {
				for _, o := range orbit {
					nlorbit[o] = true
				}
				break
			}
			if w > 0 && w < M {
				orbit = append(orbit, w)
			}
		}
		if j == 50 {
			for _, o := range orbit {
				lorbit[o] = true
			}
			lychrel++
		}
	}

	return lychrel
}

func TestProb055(t *testing.T) {
	out := prob055()
	Validate(t, 55, out)
}

func BenchmarkProb055(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob055()
	}
}
