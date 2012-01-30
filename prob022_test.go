package euler

import (
	"testing"
	"log"
	"sort"
	"io/ioutil"
	"strings"
)

func prob022() int {
	in := "testdata/prob022.txt"
	buf, err := ioutil.ReadFile(in)
	if err != nil {
		log.Print("read error on ", in, ": ", err)
		return -1
	}

	names := strings.Split(string(buf), ",")
	sort.Sort(sort.StringSlice(names))
	sum := 0
	for i, n := range names {
		for _, c := range n[1 : len(n)-1] {
			sum += (i + 1) * int(c-'A'+1)
		}
	}
	return sum
}

func TestProb022(t *testing.T) {
	out := prob022()
	Validate(t, 22, out)
}

func BenchmarkProb022(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob022()
	}
}
