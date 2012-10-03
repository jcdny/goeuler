package euler

import (
	"testing"
)

func count(s string, ch chan<- map[rune]bool) {
	k := make(map[rune]bool)
	for _, char := range s {
		k[char] = true
	}
	ch <- k
}

func pcount(s string) int {
	N := len(s)
	c := make(chan map[rune]bool)
	go count(s[0:N/2], c)
	go count(s[N/2:], c)
	k := make(map[rune]bool)
	for i := 0; i < 2; i++ {
		var ko map[rune]bool
		ko = <-c
		for key := range ko {
			k[key] = true
		}
	}

	return len(k)
}

func TestPcount(t *testing.T) {
	type atest struct {
		s string
		n int
	}

	tests := []atest{{"abc", 3}, {"123ab", 5}, {"abcabc", 3}, {"", 0}, {"aaaa", 1}}
	for _, do := range tests {
		got := pcount(do.s)
		if got != do.n {
			t.Error("Expected ", do.n, " got ", got, " for ", do.s)
		}
	}
}
