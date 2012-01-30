package euler

import (
	"testing"
)

func prob019() int {
	md := [][]int{
		{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
		{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	}

	d := 1 // 1900 starts on monday
	ly := 0
	ns := 0
	for y := 0; y < 101; y++ {
		if y%4 == 0 && y != 0 {
			ly = 1
		} else {
			ly = 0
		}
		for m := 0; m < 12; m++ {
			if y > 0 && d%7 == 0 {
				ns++
			}
			d += md[ly][m]
		}
	}

	return ns
}

func TestProb019(t *testing.T) {
	out := prob019()
	Validate(t, 19, out)
}

func BenchmarkProb019(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob019()
	}
}
