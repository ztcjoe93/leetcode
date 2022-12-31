func intersect(nums1 []int, nums2 []int) []int {
	// intersection here means "set intersection"
	// tyvm to whoever who wrote this question down, it's ambiguously vague

	hm := make(map[int]int)
	intersection := make([]int, 0)

	for _, num := range nums1 {
		if _, ok := hm[num]; !ok {
			hm[num] = 0
		}
		hm[num]++
	}

	for _, num := range nums2 {
		if _, ok := hm[num]; ok {
			if hm[num] > 0 {
				intersection = append(intersection, num)
				hm[num]--
			}
		}
	}

	return intersection
}