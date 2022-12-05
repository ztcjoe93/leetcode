package main

import "fmt"

func test_26() {
	eval([]int{1, 1, 2}, 2, removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	idx := 1
	cur := nums[0]

	for _, num := range nums[1:] {
		if num != cur {
			nums[idx] = num
			cur = num
			idx++
		}
	}

	fmt.Printf("%v\n", nums)

	return idx
}
