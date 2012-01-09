package euler

import (
	"testing"
)

func TestNDivisors(t *testing.T) {
	N := int64(100)
	pv, cv := PrimeSieve(100)
	for i := int64(0); i < N; i++ {
		NDivisors(i, pv, cv)
	}
}
