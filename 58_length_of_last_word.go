func lengthOfLastWord(s string) int {
	start := false
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			count++
			if !start {
				start = true
			}
		} else if string(s[i]) == " " && start {
			break
		}
	}

	return count
}