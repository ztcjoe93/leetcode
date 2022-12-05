package main

func test_66() {
	eval([]int{1, 2, 3}, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	eval([]int{4, 3, 2, 1}, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	eval([]int{9}, []int{1, 0}, plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	arr := make([]int, len(digits)+1)
	arr = arr[:0]

	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		result := 0
		if i == len(digits)-1 || carry {
			carry = false
			result = digits[i] + 1
		} else {
			result = digits[i]
		}

		if result == 10 {
			carry = true
			result = 0
		}

		arr = append(arr, 0)
		copy(arr[1:], arr)
		arr[0] = result
	}

	if carry {
		arr = append(arr, 0)
		copy(arr[1:], arr)
		arr[0] = 1
	}

	return arr
}
