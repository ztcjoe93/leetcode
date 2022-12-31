func isHappy(n int) bool {
	visited := make(map[int]bool)
	found := false

	for !found {
		arr := make([]int, 0)
		for n > 0 {
			arr = append(arr, n%10)
			n /= 10
		}

		sum := 0
		for _, num := range arr {
			sum += num * num
		}
		if sum == 1 {
			found = true
			break
		}

		if _, ok := visited[sum]; !ok {
			visited[sum] = true
		} else {
			return false
		}

		n = sum
	}

	return found
}