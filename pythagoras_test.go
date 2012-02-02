package euler

import (
	"testing"
)

func TestPythagoreanTriples(t *testing.T) {
	out := PythagoreanTriples(5, nil)
	if len(out) != 5 {
		t.Fail()
	}
}

func BenchmarkPythagoreanTriples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PythagoreanTriples(10000, nil)
	}
}
