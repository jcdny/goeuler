package euler

import (
	"testing"
	"log"
)

func probXXX() int {
	return 0
}

func TestProbXXX(t *testing.T) {
	out := probXXX()
	if out != 0 {
		log.Print("ProbXXX ", out)
		t.Error("ProbXXX failed ", out)
	}
}

func BenchmarkProbXXX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		probXXX()
	}
}
