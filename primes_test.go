package euler

import (
	"testing"
)

func BenchmarkPrimeSieve(b *testing.B) {
	N := 100000000
	for i := 0; i < b.N; i++ {
		PrimeSieve(N)
		//log.Print(len(p), NPrimesApprox(N))
	}
}
func BenchmarkPrimeSieveDivisors(b *testing.B) {
	N := 10000000
	for i := 0; i < b.N; i++ {
		PrimeSieveDivisors(N)
		//log.Print(len(p), NPrimesApprox(N))
	}
}

func TestPrimeSieveDivisors(t *testing.T) {
	s0, err := TableIntReadFile("testdata/oies_A000005.txt")
	if err != nil {
		t.Fatal("Read failed")
	}
	s1, err := TableIntReadFile("testdata/oies_A000203.txt")
	if err != nil {
		t.Fatal("Read failed")
	}

	n0 := s0[len(s0)-1][0]
	n1 := s1[len(s1)-1][0]
	n := n0
	if n1 > n0 {
		n = n1
	}

	_, _, dv, dsumv := PrimeSieveDivisors(n + 1)
	for _, test := range s0 {
		if len(dv[test[0]]) != test[1] {
			t.Error("Failed s0 ", test, dv[test[0]])
		}
	}

	for _, test := range s1 {
		if dsumv[test[0]] != test[1] {
			t.Error("Failed s1 ", test, dsumv[test[0]])
		}
	}
}
