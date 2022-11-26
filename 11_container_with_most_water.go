package main

func test_11() {
	fullEval(
		eval([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})),
		eval([]int{1, 1}, 1, maxArea([]int{1, 1})),
		eval([]int{8, 20, 1, 2, 3, 4, 5, 6}, 42, maxArea([]int{8, 20, 1, 2, 3, 4, 5, 6})),
	)
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	vol := 0

	for l != r {
		var low int
		if height[l] > height[r] {
			low = height[r]
		} else {
			low = height[l]
		}

		volume := (r - l) * low
		if volume > vol {
			vol = volume
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return vol
}
