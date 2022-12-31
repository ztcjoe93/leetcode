func hammingWeight(num uint32) int {
	count := 0

	bcount := uint32(1)
	for i := 0; i < 32; i++ {
		if (num & bcount) != 0 {
			count++
		}
		bcount *= 2
	}

	return count
}