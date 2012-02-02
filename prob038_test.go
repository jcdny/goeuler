package euler

import (
	"sort"
	"strconv"
	"testing"
)

func IsPandigital(s string) bool {
	b := []byte(s)
	sort.Sort(bsort(b))
	if string(b) == "123456789" {
		return true
	}
	return false
}

func IsPandigitalN(s string, n int) bool {
	b := []byte(s)
	pd := "123456789"
	sort.Sort(bsort(b))
	if string(b) == pd[:n] {
		return true
	}
	return false
}

func prob038() int {
	sm := "918273645"
	// we know the number has to start with 9 and be greater than 918273645
	// 9x and 9xx would have the wrong # of digits so look from 9xxx to 9999
	for i := 9000; i < 10000; i++ {
		if i%10 == 9 || i%10 == 0 {
			continue
		}
		s := ""
		k := 1
		for len(s) < 9 {
			s += strconv.Itoa(i * k)
			k++
		}
		if len(s) == 9 && s > sm && IsPandigital(s) {
			sm = s
		}
	}
	m, _ := strconv.Atoi(sm)
	return m
}

func TestProb038(t *testing.T) {
	out := prob038()
	Validate(t, 38, out)
}

func BenchmarkProb038(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob038()
	}
}
