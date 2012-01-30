package euler

import (
	"testing"
)

func prob006() int {
	ssq, sqs := 0, 0
	for i := 1; i <= 100; i++ {
		ssq += i * i
		sqs += i
	}
	sqs *= sqs
	return sqs - ssq
}

func TestProb006(t *testing.T) {
	out := prob006()
	Validate(t, 6, out)
}

func BenchmarkProb006(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob006()
	}
}
