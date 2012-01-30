package euler

import (
	"testing"
	"log"
	"io/ioutil"
	"strings"
)

func prob042() int {
	MN := 1000
	tf := make([]bool, MN*(MN+1)/2)
	for n := 1; n < MN; n++ {
		tf[n*(n+1)/2] = true
	}

	in := "testdata/prob042.txt"
	buf, err := ioutil.ReadFile(in)
	if err != nil {
		log.Print("read error on ", in, ": ", err)
		return -1
	}

	words := strings.Split(string(buf), ",")
	nt := 0

	for _, w := range words {
		sum := 0
		for _, c := range w[1 : len(w)-1] {
			sum += int(c - 'A' + 1)
		}
		if tf[sum] {
			nt++
		}
	}

	return nt
}

func TestProb042(t *testing.T) {
	out := prob042()
	Validate(t, 42, out)
}

func BenchmarkProb042(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob042()
	}
}
