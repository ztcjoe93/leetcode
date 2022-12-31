func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	// 3 ** 0
	i := 1
	for i < n {
		i *= 3
	}

	return i == n
}