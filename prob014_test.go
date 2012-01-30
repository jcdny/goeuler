package euler

import (
	"testing"
	"log"
)

func collatzMemo(lvec []int, sc []int64, off int) {
	//log.Print(off, sc)
	lvn := int64(len(lvec))

	for i, n := range sc {
		if n < lvn {
			lvec[n] = len(sc) + off - i
		}
	}
}

func Collatz(n int64, lvec []int, sc []int64) (int, []int64) {
	if lvec[n] != 0 {
		return lvec[n], sc[:0]
	}

	lvn := int64(len(lvec))
	sc = sc[:1]
	sc[0] = n

	for n != 1 {
		if n < lvn && lvec[n] != 0 {
			collatzMemo(lvec, sc, lvec[n]-1)
			return lvec[sc[0]], sc
		}

		if n&1 == 0 {
			n >>= 1
		} else {
			//n = n<<1 + n + 1
			n = 3*n + 1
			if n < 0 {
				log.Panic("overflow")
			}
		}

		sc = append(sc, n)
	}

	// If we fell out we never encountered a memoized seq.
	collatzMemo(lvec, sc, 0)
	return len(sc), sc
}

func prob014() int {
	N := int64(1000000)
	lvec := make([]int, N+1)
	sc := make([]int64, 1000)
	l := 0
	imax, max := 0, 0
	for i := int64(1); i < N; i++ {
		l, sc = Collatz(i, lvec, sc)
		if l > max {
			max = l
			imax = int(i)
		}
	}

	return imax
}

func prob014nomem() int {
	N := 1000000
	imax, max := 0, 0
	for i := 1; i < N; i++ {
		l := 1
		n := int64(i)
		for n != 1 {
			l++
			if n&1 == 0 {
				n >>= 1
			} else {
				//n = n<<1 + n + 1
				n = 3*n + 1
				if n < 0 {
					log.Panic("overflow")
				}
			}
		}
		if l > max {
			max = l
			imax = i
		}
	}

	return imax
}

func TestProb014(t *testing.T) {
	out := prob014()
	Validate(t, 14, out)
}

func BenchmarkProb014(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob014()
	}
}

func TestProb014nomem(t *testing.T) {
	out := prob014nomem()
	Validate(t, 14, out)
}

func BenchmarkProb014nomem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob014nomem()
	}
}
