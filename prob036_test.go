package euler

import (
	"strconv"
	"testing"
)

// see Hacker's Delight chapter 7 for this
func bitreverse(x uint) uint {
	x = (x&0x55555555)<<1 | (x&0xAAAAAAAA)>>1
	x = (x&0x33333333)<<2 | (x&0xCCCCCCCC)>>2
	x = (x&0x0F0F0F0F)<<4 | (x&0xF0F0F0F0)>>4
	x = (x&0x00FF00FF)<<8 | (x&0xFF00FF00)>>8
	x = (x&0x0000FFFF)<<16 | (x&0xFFFF0000)>>16

	return x
}

func prob036() int {
	N := uint(1000000)
	sum := uint(0)
	for i := uint(1); i < N; i += 2 {
		r := bitreverse(i)
		r >>= Nlz(i)
		if i == r {
			if Palindrome(strconv.Itoa(int(i))) {
				sum += i
			}
		}
	}

	return int(sum)
}

func TestProb036(t *testing.T) {
	out := prob036()
	Validate(t, 36, out)
}

func BenchmarkProb036(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob036()
	}
}
