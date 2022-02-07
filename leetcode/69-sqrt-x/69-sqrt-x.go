package main

// See https://en.wikipedia.org/wiki/Integer_square_root.
func mySqrt(x int) int {
	k := x / 2
	if k == 0 {
		return x
	}
	m := (k + x/k) / 2
	for m < k {
		k = m
		m = (k + x/k) / 2
	}
	return k
}
