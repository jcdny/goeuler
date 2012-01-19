    // compute bits per byte lut
	var bits [256]int
	for i := 0; i < 256; i++ {
		for s := uint(0); s < 8; s++ {
			bits[i] += (i >> s) & 1
		}
	}
