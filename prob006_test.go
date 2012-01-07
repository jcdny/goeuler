package euler

import (
	"testing"
	"log"
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
	if prob006() != 25164150 {
		log.Print(prob006())
		t.Error("Prob006 failed")
	}
}

func BenchmarkProb006(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob006()
	}
}
