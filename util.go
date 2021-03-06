package euler

import (
	"math/big"
)

func DigitReverse(n int64) int64 {
	rev := int64(0)
	for n > 0 {
		rev = rev*10 + (n % 10)
		n /= 10
	}
	return rev
}

func SumDigits(z *big.Int) int {
	n := 0
	for _, c := range z.String() {
		n += int(c - '0')
	}

	return n
}

func Palindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

// Number of leading zeroes.
func Nlz(x uint) uint {
	n := uint(32)
	y := x >> 16
	if y != 0 {
		n -= 16
		x = y
	}
	y = x >> 8
	if y != 0 {
		n -= 8
		x = y
	}
	y = x >> 4
	if y != 0 {
		n -= 4
		x = y
	}
	y = x >> 2
	if y != 0 {
		n -= 2
		x = y
	}
	y = x >> 1
	if y != 0 {
		return n - 2
	}

	return n - x
}

func Ttz(x uint) uint {
	if x == 0 {
		return x
	}
	if x&0xFFFF == 0 {
		x >>= 16
	}
	if x&0xFF == 0 {
		x >>= 8
	}
	if x&0x0F == 0 {
		x >>= 4
	}
	if x&0x03 == 0 {
		x >>= 2
	}
	if x&1 == 0 {
		x >>= 1
	}

	return x
}
