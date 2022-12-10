package main

func firstUniqChar(s string) int {
	hm := make(map[byte]int)
	num := -1
	for i := len(s) - 1; i >= 0; i-- {
		_, ok := hm[s[i]]
		if !ok {
			hm[s[i]] = i
		} else {
			hm[s[i]] = -1
		}
	}

	for _, v := range hm {
		if v != -1 {
			if num == -1 || v < num {
				num = v
			}
		}
	}

	return num
}
