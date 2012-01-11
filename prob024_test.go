package euler

import (
	"testing"
)

func prob024() string {
	return PermuteString("0123456789", 1000000)
}

func TestProb024(t *testing.T) {
	out := prob024()
	t.Log("Problem 024 ", out)
	if out != "2783915460" {
		t.Fail()
	}
}

func BenchmarkProb024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob024()
	}
}
