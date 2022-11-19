package main

import "fmt"

func test_8() {
	inputs := []string{
		"42", "-1312", "2147483647", "-2147483648",
		" w+42", " +312 w", " +w312 w", ".42",
		".4324.231", "341.23", " b11228552307",
	}

	outputs := []int{42, -1312, 2147483647, -2147483648, 0, 312, 0, 0, 0, 341, 0}

	passed, failed := 0, 0

	for i := range inputs {
		res := myAtoi(inputs[i])
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

func myAtoi(s string) int {

	whitespace := true
	positive := true

	checkChar := false

	digits := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	atoiStr := ""
	atoiInt := 0

	// retrieving atoi string
	for _, val := range s {
		valStr := string(val)
		if whitespace && valStr == " " {
			continue
		}

		whitespace = false
		if !checkChar {
			checkChar = true
			if _, ok := digits[valStr]; ok {
				atoiStr += valStr
				continue
			}

			if valStr == "+" {
				positive = true
			} else if valStr == "-" {
				positive = false
			} else {
				break
			}
			continue
		}

		_, ok := digits[valStr]

		if ok {
			atoiStr += valStr
		} else {
			break
		}
	}

	// early termination for non-valid atoiStrs
	if atoiStr == "" {
		return atoiInt
	}

	UPPER_LIMIT := 2147483647
	LOWER_LIMIT := -2147483648
	// converting atoiStr into actual integer
	if positive {
		for _, chrx := range atoiStr {
			if atoiInt > UPPER_LIMIT/10 || atoiInt == UPPER_LIMIT/10 && digits[string(chrx)] > UPPER_LIMIT%10 {
				atoiInt = UPPER_LIMIT
				break
			}
			atoiInt *= 10
			atoiInt += digits[string(chrx)]
		}
	} else {
		for _, chrx := range atoiStr {
			if atoiInt < LOWER_LIMIT/10 || atoiInt == LOWER_LIMIT/10 && digits[string(chrx)] > -(LOWER_LIMIT%10) {
				atoiInt = LOWER_LIMIT
				break
			}
			atoiInt *= 10
			atoiInt += digits[string(chrx)] * -1
		}
	}

	return atoiInt
}
