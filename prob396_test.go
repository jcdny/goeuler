package euler

import (
	"log"
	"testing"
)

func goodstein_sequence(n int) int64 {
	d := make([]int64, 0, 10000)
	// binary expansion of n
	for n > 0 {
		if n&1 != 0 {
			d = append(d, 1)
		} else {
			d = append(d, 0)
		}
		n = n >> 1
	}

	var base int64
	for base = 3; len(d) > 0; base++ {	
		for i, digit := range d {
			if digit > 0 {
				d[i]--
				break
			} else {
				d[i] = base - 1
			}	
		}			
	    	val := int64(0)
		pow := int64(1)
		for _, digit := range d {
			val += digit*pow
			pow *= base
		}			
		log.Print("base, val: ",base, val)
		
		if d[len(d)-1] == 0 {
			log.Print(n, base)
			d = d[:len(d)-1]
		}
	}

	return base - 3
}

func prob396() int {
	tot := int64(0)
	for i := 1; i < 8; i++ {
		d := goodstein_sequence(i)
		tot += d
		log.Print("********** ", i, "=", d)
	}
	log.Print("total ", tot)
	return 0
}

func TestProb396(t *testing.T) {
	out := prob396()
	Validate(t, 396, out)
}

func BenchmarkProb396(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob396()
	}
}
