package euler

import (
	"big"
)

func SumDigits(z *big.Int) int {
	s := z.String()
	n := 0
	for _, c := range s {
		n += c - '0'
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
