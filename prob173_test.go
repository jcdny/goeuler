package euler

import (
	"testing"
	. "math"
)

func prob173brute() int64 {
	sum := int64(0)
	N := int64(1000000)
	max := N/4 + 1
	for n := int64(3); n <= max; n++ {
		n2 := n * n
		for m := n - 2; m > 0; m -= 2 {
			if n2-m*m > N {
				break
			}
			sum++
		}
	}
	return sum
}

func prob173() int64 {
	sum := int64(0)
	N := int64(1000000)
	max := N/4 + 1
	for n := int64(3); n <= max; n++ {
		if n*n <= N+1 {
			// we can entirely tile the square so we can make
			// all the lamina
			sum += (n - 1) / 2
		} else {
			// we can only do some of them, calculate the
			// first d st n^2 - (2d+o)^2 <= N
			d := int64(Ceil((Sqrt(float64(n*n-N)) - float64(n&1)) / 2))
			sum += n/2 - d
		}
	}
	return sum
}

func TestProb173(t *testing.T) {
	out := prob173()
	t.Log("Problem 173 ", out)
	if out != 1572729 {
		t.Fail()
	}
}
func TestProb173brute(t *testing.T) {
	out := prob173brute()
	t.Log("Problem 173 ", out)
	if out != 1572729 {
		t.Fail()
	}
}

func BenchmarkProb173(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob173()
	}
}
func BenchmarkProb173brute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob173brute()
	}
}
