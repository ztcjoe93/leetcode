func findPeakElement(nums []int) int {
	peak := 0
	for idx := range nums {
		if idx == 0 || idx == len(nums)-1 {
			if idx == 0 && idx+1 < len(nums)-1 && nums[idx+1] < nums[idx] {
				return idx
			} else if idx == len(nums)-1 && idx-1 >= 0 && nums[idx-1] < nums[idx] {
				return idx
			}
			continue
		}

		if nums[idx-1] < nums[idx] && nums[idx+1] < nums[idx] {
			peak = idx
			break
		}
	}
	return peak
}