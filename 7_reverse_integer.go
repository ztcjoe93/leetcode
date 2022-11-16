package main

import "fmt"

func test_7() {
	inputs := []int{
		123, -123, 120, -2147483648,
		2147483647, 947483647, -47483647,
		1534236469,
	}

	outputs := []int{
		321, -321, 21, 0, 0, 746384749, -74638474, 0,
	}

	passed, failed := 0, 0
	for i, _ := range inputs {
		res := reverse(inputs[i])
		if res == outputs[i] {
			fmt.Printf("Test case PASSED input=%v output=%v expected=%v\n", inputs[i], res, outputs[i])
			passed++
		} else {
			fmt.Printf("Test case FAILED input=%v output=%v expected=%v\n", inputs[i], res, outputs[i])
			failed++
		}
	}
	fmt.Printf("%v/%v TEST CASES PASSED\n", passed, passed+failed)
}

func reverse(x int) int {
	val := 0
	UPPER_LIMIT := 2147483647
	LOWER_LIMIT := -2147483648

	if x > 0 {
		for x >= 10 {
			val *= 10
			digit := x % 10
			val += digit
			x /= 10
		}

		if val > UPPER_LIMIT/10 || val == UPPER_LIMIT/10 && x > UPPER_LIMIT%10 {
			return 0
		} else {
			val *= 10
			val += x
		}
	} else if x < 0 {
		for x <= -10 {
			val *= 10
			digit := x % -10
			val += digit
			x /= 10
		}

		if val < LOWER_LIMIT/10 || val == LOWER_LIMIT/10 && x > -(LOWER_LIMIT/10) {
			return 0
		} else {
			val *= 10
			val += x
		}
	}

	return val
}
