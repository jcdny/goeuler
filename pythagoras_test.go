package euler

import (
	"testing"
	"log"
)

func TestPythagoreanTriples(t *testing.T) {
	out := PythagoreanTriples(5, nil)
	log.Print(out)
}

func BenchmarkPythagoreanTriples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PythagoreanTriples(10000, nil)
	}
}
