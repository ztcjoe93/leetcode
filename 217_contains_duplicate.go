package main

func containsDuplicate(nums []int) bool {
	hm := make(map[int]int)

	for _, num := range nums {
		if _, ok := hm[num]; ok {
			return true
		}

		hm[num] = 1
	}

	return false
}
