package euler

import (
	"log"
	"testing"
)

func ProdMax(table [][]int, n int) int {
	max := -1
	for r, row := range table {
		for c := range row {
			// across
			if c > n-2 {
				prod := 1
				for i := 0; i < n; i++ {
					prod *= row[c-i]
				}
				if prod > max {
					max = prod
				}
			}
			// down 
			if r > n-2 {
				prod := 1
				for i := 0; i < n; i++ {
					prod *= table[r-i][c]
				}
				if prod > max {
					max = prod
				}
			}
			// neg diag
			if r > n-2 && c > n-2 {
				prod := 1
				for i := 0; i < n; i++ {
					prod *= table[r-i][c-i]
				}
				if prod > max {
					max = prod
				}
			}
			// pos diag 
			if r > n-2 && c < len(row)-n {
				prod := 1
				for i := 0; i < n; i++ {
					prod *= table[r-i][c+i]
				}
				if prod > max {
					max = prod
				}
			}
		}
	}

	return max
}

func prob011() int {
	table, err := TableIntReadFile("testdata/prob011.txt")
	if err != nil {
		log.Print(err)
		return -1
	}

	return ProdMax(table, 4)
}

func TestProb011(t *testing.T) {
	out := prob011()
	Validate(t, 11, out)
}

func BenchmarkProb011(b *testing.B) {
	table, _ := TableIntReadFile("testdata/prob011.txt")
	for i := 0; i < b.N; i++ {
		ProdMax(table, 4)
	}
}
