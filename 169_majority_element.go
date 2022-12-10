package main

func majorityElement(nums []int) int {
	hm := make(map[int]int)
	var target int

	if len(nums) == 1 {
		return nums[0]
	}

	for _, num := range nums {
		if _, ok := hm[num]; ok {
			hm[num]++
			if hm[num] > len(nums)/2 {
				target = num
				break
			}
		} else {
			hm[num] = 1
		}
	}

	return target
}
