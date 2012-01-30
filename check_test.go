package euler

import (
	"testing"
)

type checkTest struct {
	p       int
	ai      int
	as      string
	solved  bool
	correct bool
}

func TestCheck(t *testing.T) {
	var test = []checkTest{
		{1, 233168, "233168", true, true},
		{1, 23316, "23316", true, false},
		{999, 0, "0", false, false},
	}

	for _, c := range test {
		correct, solved := Check(c.p, c.ai)
		if c.solved != solved || c.correct != correct {
			t.Error("int", t)
		}
		correct, solved = Check(c.p, c.as)
		if c.solved != solved || c.correct != correct {
			t.Error("string", t)
		}
	}
}
