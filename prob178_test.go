package euler

import (
	"testing"
	"log"
)

func addtrailingdigit(pd []uint8) [][]uint8 {
	out := make([][]uint8, 4)
	h := pd[0]
	t := pd[len(pd)-1]
	n := 0
	if h != 9 {
		out[n] = append([]uint8{h + 1}, pd...)
		n++
	} else if h != 1 {
		out[n] = append([]uint8{h - 1}, pd...)
		n++
	}
	if t != 0 {
		b := make([]uint8, 0, len(pd)+1)
		b = append(b, pd...)
		out[n] = append(b, t-1)
		n++
	} else if t != 9 {
		b := make([]uint8, 0, len(pd)+1)
		b = append(b, pd...)
		out[n] = append(b, t+1)
		n++
	}
	out = out[:n]

	for _, s := range out {
		if len(s) != len(pd)+1 {
			log.Print("bad len ", s, " from ", pd)
		}
	}

	return out
}

func interiordigits(pd []uint8) [][]uint8 {
	out := make([][]uint8, 0, (len(pd)-2)*2)

	pre := pd[:1]
	d := pd[1]
	post := pd[1:]
	for len(post) > 1 {
		if d < 9 {
			b := append([]uint8{}, pre...)
			b = append(b, d, d+1)
			b = append(b, post...)
			out = append(out, b)
		}
		if d > 0 {
			b := append([]uint8{}, pre...)
			b = append(b, d, d-1)
			b = append(b, post...)
			out = append(out, b)
		}
		pre = append(pre, d)
		post = post[1:]
		d = post[0]
	}
	for _, s := range out {
		if len(s) != len(pd)+2 {
			log.Print("bad len ", s, " from ", pd)
		}
	}

	return out
}

func prob178() int {
	return 1

	cs := [][]uint8{{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}
	ss := cs
	n := 1
	for d := 11; d < 40; d++ {
		w := make([][]uint8, 0, len(cs)*8)
		if d&1 == 1 {
			ss = cs // on odd ones save the current even number set
		} else {
			// even number steps we insert interior points to the 
			// set two steps ago
			for _, p := range ss {
				w = append(w, interiordigits(p)...)
				// log.Print("interior got ", len(w), ":\n", w)
			}
		}
		for _, p := range cs {
			w = append(w, addtrailingdigit(p)...)
		}
		cs = w
		n += len(w)
		log.Print(d, n, len(w))
		// log.Print(w)
	}

	return 1
}

func TestProb178(t *testing.T) {
	out := prob178()
	t.Log("Problem 178 ", out)
	log.Print("Not done")
	if out != 0 {
		t.Fail()
	}
}

func BenchmarkProb178(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prob178()
	}
}
