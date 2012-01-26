package euler

type PTriple struct {
	A, B, C int
}

// generate the primitive pythagorean triples
// cf http://mathworld.wolfram.com/PythagoreanTriple.html
func PythagoreanTriples(n int, reject func(PTriple) bool) []PTriple {
	// generate n primitive triples
	pt := make([]PTriple, 1, n+3)
	pt[0] = PTriple{3, 4, 5}

	return genPythagoreanTriples(n, 0, pt, reject)
}

func genPythagoreanTriples(n, j int, pt []PTriple, reject func(PTriple) bool) []PTriple {
	k := len(pt)
	for _, p := range pt[j:] {
		pnew := PTriple{p.A - 2*p.B + 2*p.C, 2*p.A - p.B + 2*p.C, 2*p.A - 2*p.B + 3*p.C} // U
		if len(pt) < n && (reject == nil || !reject(pnew)) {
			pt = append(pt, pnew)
		}
		pnew = PTriple{p.A + 2*p.B + 2*p.C, 2*p.A + p.B + 2*p.C, 2*p.A + 2*p.B + 3*p.C} // A
		if len(pt) < n && (reject == nil || !reject(pnew)) {
			pt = append(pt, pnew)
		}
		pnew = PTriple{2*p.B - p.A + 2*p.C, p.B - 2*p.A + 2*p.C, 2*p.B - 2*p.A + 3*p.C} // D
		if len(pt) < n && (reject == nil || !reject(pnew)) {
			pt = append(pt, pnew)
		}
	}

	if len(pt) == n || j == k {
		return pt
	}

	return genPythagoreanTriples(n, k, pt, reject)
}
