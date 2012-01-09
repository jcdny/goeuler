package euler

import (
	"testing"
	"big"
	"log"
)

func prob181() string {
	// Just use math for this.
	// its (2N)!/(N!N!) so in this case 40!/20!/20!
	// All praise Pascal...
	z := big.NewInt(0)
	z = z.MulRange(1, 60)
	z2 := big.NewInt(0)
	z2 = z2.MulRange(1, 40)
	z = z.Mul(z, z2)
	return z.String()
}

func TestProb181(t *testing.T) {
	log.Print("Prob181 not done")
	out := prob181()
	if out == "0" {
		t.Error("Prob181 ", out)
	}
}

func BenchmarkProb181(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob181()
	}
}
