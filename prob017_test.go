package euler

import (
	"testing"
	"log"
)

var words = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func prob017() int {
	s := ""

	for i := 1; i < 1000; i++ {
		if w, ok := words[i]; ok {
			s += w
		} else {
			ii := i
			if i > 99 {
				s += words[i/100] + "hundred"
				ii %= 100
				if ii != 0 {
					s += "and"
				}
			}
			if w, ok := words[ii]; ok {
				s += w
				continue
			}
			s += words[(ii/10)*10] + words[ii%10]
		}
	}
	s += "onethousand"

	return len(s)
}

func TestProb017(t *testing.T) {
	if prob017() != 21124 {
		log.Print(prob017())
		t.Error("Prob017 failed")
	}
}

func BenchmarkProb017(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob017()
	}
}
