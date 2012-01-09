package euler

import (
	"testing"
)

func BenchmarkPrimeSieve(b *testing.B) {
	N := 1000000
	for i := 0; i < b.N; i++ {
		PrimeSieve(N)
		//log.Print(len(p), NPrimesApprox(N))
	}
}
