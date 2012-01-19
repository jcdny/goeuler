package euler

import (
	"testing"
	"strconv"
	"sort"
	"log"
)

func iminhash(i int) string {
	s := strconv.Itoa(i)
	b := []byte(s)
	sort.Sort(bsort(b))
	s = string(b)
	return s
}

func prob049() int64 {
	mmap := make(map[string][]int, 10000)
	pv, cv := PrimeSieve(10000)
	for _, p := range pv {
		if p < 1000 {
			continue
		}
		s := iminhash(p)
		if len(mmap[s]) > 0 {
			for _, v := range mmap[s] {
				as := 2*p - v
				if as > len(cv) {
					break
				}
				if !cv[as] {
					ash := iminhash(as)
					// looking for the sequence not 1487...
					if ash == s && ash != "1478" {
						return int64(v)*100000000 + int64(p)*10000 + int64(as)
					}
				}
			}
		}
		mmap[s] = append(mmap[s], p)
	}
	return 1
}

func TestProb049(t *testing.T) {
	out := prob049()
	t.Log("Problem 049 ", out)
	log.Print("Not done")
	if out != 296962999629 {
		t.Fail()
	}
}

func BenchmarkProb049(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob049()
	}
}
