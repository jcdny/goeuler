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

	return TrianglePathMax(table) // see prob067_test.go for the actual code.
}

func TestProb018(t *testing.T) {
	out := prob018()
	Validate(t, 18, out)
}

func BenchmarkProb018(b *testing.B) {
	table, _ := TableIntReadFile("testdata/prob018.txt")
	for i := 0; i < b.N; i++ {
		TrianglePathMax(table)
	}
}
