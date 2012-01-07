package euler

import (
	"testing"
	"log"
)

func prob018() int {
	table, err := TableIntReadFile("testdata/prob018.txt")
	if err != nil {
		log.Print(err)
		return -1
	}

	return TrianglePathMax(table) // cf prob067_test.go for the actual code.
}

func TestProb018(t *testing.T) {
	if prob018() != 1074 {
		log.Print(prob018())
		t.Error("Prob018 failed")
	}
}

func BenchmarkProb018(b *testing.B) {
	table, _ := TableIntReadFile("testdata/prob018.txt")
	for i := 0; i < b.N; i++ {
		TrianglePathMax(table)
	}
}