package euler

import (
	"testing"
	. "math"
	"log"
)

func TestNDivisors(t *testing.T) {
	table, err := TableIntReadFile("testdata/oies_A000005.txt")
	if err != nil {
		t.Fatal("Read failed")
	}
	pv, cv := PrimeSieve(10000)
	for _, test := range table {
		nd := NDivisors(int64(test[0]), pv, cv)
		if nd != test[1] {
			t.Error("Failed ", test, nd)
		}
	}
}

func BenchmarkNDivisors(b *testing.B) {
	N := int64(1000000)
	RN := int(Sqrt(float64(N)))
	for i := 0; i < b.N; i++ {
		pv, cv := PrimeSieve(int(N))
		for j := int64(1); j < 1000000; j++ {
			NDivisors(j, pv, cv)
		}
	}
	log.Print("Used N ", N, " RN ", RN)
}
