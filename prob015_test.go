package euler

import (
	"math/big"
	"testing"
)

func prob015() string {
	// Just use math for this.
	// its (2N)!/(N!N!) so in this case 40!/20!/20!
	// All praise Pascal...
	z := big.NewInt(0)
	z = z.MulRange(1, 40)
	z2 := big.NewInt(0)
	z2 = z2.MulRange(1, 20)
	z = z.Div(z, z2)
	z = z.Div(z, z2)

	return z.String()

}

func TestProb015(t *testing.T) {
	out := prob015()
	Validate(t, 15, out)
}

func BenchmarkProb015(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob015()
	}
}
