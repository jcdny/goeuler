package euler

import (
	"testing"
	"log"
)

func TrianglePathMax(table [][]int) int {
	for r := len(table) - 2; r >= 0; r-- {
		for c := range table[r] {
			if table[r+1][c] > table[r+1][c+1] {
				table[r][c] = table[r][c] + table[r+1][c]
			} else {
				table[r][c] = table[r][c] + table[r+1][c+1]
			}
		}
	}

	return table[0][0]
}

func prob067() int {
	table, err := TableIntReadFile("testdata/prob067.txt")
	//table, err := TableIntReadFile("testdata/prob067sample.txt") // produces 23
	if err != nil {
		log.Print(err)
		return -1
	}

	return TrianglePathMax(table)
}

func TestProb067(t *testing.T) {
	out := prob067()
	Validate(t, 67, out)
}

func BenchmarkProb067(b *testing.B) {
	table, _ := TableIntReadFile("testdata/prob067.txt")
	for i := 0; i < b.N; i++ {
		TrianglePathMax(table)
	}
}
