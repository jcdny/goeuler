package euler

import (
	"testing"
)

func BenchmarkPermuteArray(b *testing.B) {
	perm := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		PermuteArray(perm, []int{})
	}
}
func BenchmarkPermuteArrayAlloc(b *testing.B) {
	perm := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buf := make([]int, 362880*9)
	for i := 0; i < b.N; i++ {
		PermuteArray(perm, buf)
	}
}

func TestPermuteArray(t *testing.T) {
	perm := []int{1}
	np := 1
	for n := 2; n < 6; n++ {
		perm = append(perm, n)
		np *= n
		pa := PermuteArray(perm, []int{})
		if len(pa) != np {
			t.Error("Wrong permute count ", len(pa), np)
		}
		for i, p1 := range pa[:len(pa)-1] {
			for _, p2 := range pa[i+1:] {
				if len(p1) != len(p2) || len(perm) != len(p1) {
					t.Error("Length mismatch", p1, p2)
				}
				var j int
				for j = 0; j < len(p1); j++ {
					if p1[j] != p2[j] {
						break
					}
				}
				if j == len(p1) {
					t.Error("Duplicate permute", p1, p2)
				}
			}
		}
	}
}
