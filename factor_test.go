package euler

import (
	"log"
	. "math"
	"testing"
)

func TestNDivisors(t *testing.T) {
	table, err := TableIntReadFile("testdata/oies_A000005.txt")
	if err != nil {
		t.Fatal("Read failed")
	}
	pv, cv := PrimeSieve(10000)
	for _, test := range table {
		nd := NDivisors(uint64(test[0]), pv, cv)
		if nd != test[1] {
			t.Error("Failed ", test, nd)
		}
	}
}

func TestDivisors(t *testing.T) {
	table, err := TableIntReadFile("testdata/oies_A000005.txt")
	if err != nil {
		t.Fatal("Read failed")
	}
	pv, cv := PrimeSieve(10000)
	for _, test := range table[4:] {
		pdiv := Divisors(uint64(test[0]), pv, cv)
		if len(pdiv) != test[1] {
			t.Error("Failed ", test, pdiv)
		}
	}
}

func BenchmarkNDivisors(b *testing.B) {
	N := uint64(1000000)
	RN := int(Sqrt(float64(N)))
	for i := 0; i < b.N; i++ {
		pv, cv := PrimeSieve(int(N))
		for j := uint64(1); j < 1000000; j++ {
			NDivisors(j, pv, cv)
		}
	}
	log.Print("Used N ", N, " RN ", RN)
}
