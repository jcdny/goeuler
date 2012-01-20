package euler

import (
	"testing"
	"log"
)

func prob079() int {
	// I did this one by hand 
	// sort -u < keylog.txt gives:
	keys := []int{
		129, 160, 162, 168, 180, 289, 290,
		316, 318, 319, 362, 368, 380, 389,
		620, 629, 680, 689, 690, 710, 716,
		718, 719, 720, 728, 729, 731, 736,
		760, 762, 769, 790, 890}
	// Make a table per digit like so
	//  before:D:after
	// 9876321 0
	//      73 1 02689
	//    7631 2 089
	//       7 3 012689
	//     731 6 289
	//         7 0123689
	//   76321 8 09
	//  876321 9 0
	//
	// from the above you can see 7 is the first digit removing 7 from
	// before all digits leaves 3 as the second digit and similiarly
	// removing 3 leaves 1 etc.
	//
	// once you have placed all the digits you end up with 
	// 73162890.
	//
	// going back to the original keylog data you can see all the
	// triples are satisfied by this sequence and each digit occurs
	// once so it is a minimal length sequence.
	if false {
		log.Print(keys)
	}
	return 73162890
}

func TestProb079(t *testing.T) {
	out := prob079()
	t.Log("Problem 079 ", out)
	if out != 73162890 {
		t.Fail()
	}
}

func BenchmarkProb079(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob079()
	}
}
