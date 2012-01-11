package euler

import (
	"testing"
)

func prob009() int {
	// inelegant 
	var a, b, c int
OUT:
	for a = 3; a < 500; a++ {
		for b, c = a+1, 1000-a-a-1; b < 1000-a; b, c = b+1, c-1 {
			if a*a+b*b == c*c {
				break OUT
			}
		}
	}
	return a * b * c
}

func TestProb009(t *testing.T) {
	out := prob009()
	t.Log("Problem 009 ", out)
	if out != 31875000 {
		t.Fail()
	}
}

func BenchmarkProb009(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob009()
	}
}
