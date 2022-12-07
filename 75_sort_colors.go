package main

func test_75() {}

func sortColors(nums []int) {
	cur := 0

	for cur < len(nums)-1 {
		if nums[cur+1] < nums[cur] {
			for cur >= 0 && nums[cur] > nums[cur+1] {
				tmp := nums[cur]
				nums[cur] = nums[cur+1]
				nums[cur+1] = tmp
				cur--
			}
		}
		cur++
	}
}
