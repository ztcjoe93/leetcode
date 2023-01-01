func reverseBits(num uint32) uint32 {
	b := uint32(2147483648)
	c := uint32(1)
	var sum uint32 = 0
	for i := 0; i < 32; i++ {
		if (num & c) != 0 {
			sum += b
		}
		c *= 2
		b /= 2
	}

	return sum
}