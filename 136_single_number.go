package main

func test_136() {
	fullEval(
		eval([]int{2, 2, 1}, 1, singleNumber([]int{2, 2, 1})),
		eval([]int{4, 1, 2, 1, 2}, 4, singleNumber([]int{4, 1, 2, 1, 2})),
		eval([]int{1}, 1, singleNumber([]int{1})),
	)
}

func singleNumber(nums []int) int {
	cmp := nums[0]
	for _, num := range nums[1:] {
		cmp ^= num
	}

	return cmp
}
