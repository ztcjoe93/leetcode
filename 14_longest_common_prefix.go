package main

func test_14() {
	fullEval(
		eval([]string{"flower", "flow", "flight"}, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"})),
		eval([]string{"dog", "racecar", "car"}, "", longestCommonPrefix([]string{"dog", "racecar", "car"})),
	)
}

func longestCommonPrefix(strs []string) string {
	cur := 0
	prefix := ""

pfx:
	for {
		curPrefix := ""
		for _, str := range strs {
			if cur >= len(str) {
				break pfx
			}

			if curPrefix == "" {
				curPrefix = string(str[cur])
			}

			if string(str[cur]) != curPrefix {
				break pfx
			}
		}
		prefix += curPrefix
		cur++
	}

	return prefix
}
