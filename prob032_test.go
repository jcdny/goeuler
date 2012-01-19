package euler

import (
	"testing"
)

func prob032() int {
	pan := []int{} // the actual pandigital numbers we found

	// pairs of initial digits for multipliers given final digit
	// of product
	pairs := [][][2]int{
		1: {{7, 3}},
		2: {{9, 8}, {7, 6}, {8, 4}, {4, 3}},
		3: {{9, 7}},
		4: {{9, 6}, {8, 3}, {7, 2}},
		6: {{9, 4}, {8, 7}, {3, 2}, {8, 2}},
		7: {{9, 3}},
		8: {{9, 2}, {7, 4}, {3, 6}, {2, 4}},
	}

	dz := make([]int, 10)
	dz[0] = 1
	dc := make([]int, 10)

	// use these below for permuting possible final 3 digits
	try := make([]int, 3)
	perms := [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

SKIP:
	// numbers have to be between 1354*2 and 9876
	for p := 2708; p < 9877; p++ {
		d1 := p % 10
		if d1 == 0 || d1 == 5 || d1 == 9 {
			// 0, 5 or 9 can't be first digit in product
			continue
		}

		// check its a valid partial pandigital.
		copy(dc, dz)
		pt := p
		d := d1
		dc[d]++
		for pt > 0 {
			if dc[d] > 1 {
				// anything with a duplicate digit is discarded
				continue SKIP
			}
			pt /= 10
			d = pt % 10
			dc[d]++
		}

	FOUND:
		// now check each valid initial digit pair until we find a
		// pandigital product
		for _, pair := range pairs[d1] {
			if dc[pair[0]] > 0 || dc[pair[1]] > 0 {
				continue
			}
			dc[pair[0]]++
			dc[pair[1]]++

			// at this point there are 3 digits unallocated
			try = try[:0]
			//factor := p%pair[0] == 0 || p%pair[1] == 0
			for i, c := range dc {
				if c == 0 {
					try = append(try, i)
					//factor = factor || p%(i*10+pair[0]) == 0 || p%(i*10+pair[1]) == 0
				}
			}
			if true { //factor {
				for _, pd := range perms {
					if (try[pd[0]]*1000+try[pd[1]]*100+try[pd[2]]*10+pair[0])*pair[1] == p ||
						(try[pd[0]]*1000+try[pd[1]]*100+try[pd[2]]*10+pair[1])*pair[0] == p ||
						(try[pd[0]]*100+try[pd[1]]*10+pair[1])*(try[pd[2]]*10+pair[0]) == p ||
						(try[pd[0]]*100+try[pd[1]]*10+pair[0])*(try[pd[2]]*10+pair[1]) == p {
						pan = append(pan, p)
						break FOUND
					}
				}
			}
			dc[pair[0]]--
			dc[pair[1]]--
		}
	}

	sum := 0
	for _, pd := range pan {
		sum += pd
	}

	return sum
}

func TestProb032(t *testing.T) {
	out := prob032()
	t.Log("Problem 032 ", out)
	if out != 45228 {
		t.Fail()
	}
}

func BenchmarkProb032(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob032()
	}
}
