package euler

import (
	"testing"
	"big"
	"sort"
)

type bigIntSlice []*big.Int

func (s bigIntSlice) Len() int           { return len(s) }
func (s bigIntSlice) Less(i, j int) bool { return s[i].Cmp(s[j]) < 0 }
func (s bigIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func prob029() int {
	M := 101
	z := make([]*big.Int, 0, 100*100)

	za := big.NewInt(0)
	for a := 2; a < M; a++ {
		za.SetInt64(int64(a))
		for b := 2; b < M; b++ {
			t := big.NewInt(int64(b))
			t.Exp(za, t, nil)
			z = append(z, t)
		}
	}

	sort.Sort(bigIntSlice(z))
	zp := z[0]
	n := 1
	for _, zn := range z[1:] {
		if zp.Cmp(zn) == 0 {
			continue
		}
		n++
		zp = zn
	}

	return n
}

func TestProb029(t *testing.T) {
	out := prob029()
	t.Log("Problem 029 ", out)
	if out != 9183 {
		t.Fail()
	}
}

func BenchmarkProb029(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob029()
	}
}
