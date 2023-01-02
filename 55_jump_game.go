func canJump(nums []int) bool {

	arr := make([]bool, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			arr[i] = true
		} else {
			var ul int
			if i+nums[i] > len(nums)-1 {
				arr[i] = true
			} else {
				ul = i + nums[i]
				found := false
				for j := i; j <= ul; j++ {
					if arr[j] {
						found = true
					}
				}
				arr[i] = found
			}
		}
	}

	return arr[0]
}