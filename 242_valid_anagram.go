package main

func isAnagram(s string, t string) bool {
	hm := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := hm[s[i]]; ok {
			hm[s[i]]++
		} else {
			hm[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := hm[t[i]]; ok {
			hm[t[i]]--
		} else {
			hm[t[i]] = -1
		}
	}

	for _, val := range hm {
		if val != 0 {
			return false
		}
	}

	return true
}
