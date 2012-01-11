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
	t.Log("Problem XXX ", out)
	log.Print("Not done")
	if out != 0 {
		t.Fail()
	}
}

func BenchmarkProbXXX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		probXXX()
	}
}
