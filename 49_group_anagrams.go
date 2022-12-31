func groupAnagrams(strs []string) [][]string {
	anagramHm := make(map[string][]string)
	for _, str := range strs {
		hm := make(map[string]int)
		for _, char := range str {
			if _, ok := hm[string(char)]; !ok {
				hm[string(char)] = 0
			}
			hm[string(char)]++
		}
		keys := make([]string, len(str))
		count := 0
		for key := range hm {
			keys[count] = key + strconv.Itoa(hm[key])
			count++
		}
		sort.Strings(keys)
		combinedKey := strings.Join(keys, "")
		if _, ok := anagramHm[combinedKey]; !ok {
			anagramHm[combinedKey] = make([]string, 0)
		}
		anagramHm[combinedKey] = append(anagramHm[combinedKey], str)

	}

	anagrams := make([][]string, 0)
	for _, v := range anagramHm {
		anagrams = append(anagrams, v)
	}

	return anagrams
}