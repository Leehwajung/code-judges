package main

func reverseBits(num uint32) uint32 {
	i := 1
	result := uint32(0)
	for digit := uint32(1); digit != 0; digit <<= 1 {
		if num&digit == digit {
			result |= 1 << (32 - i)
		}
		i++
	}
	return result
}
