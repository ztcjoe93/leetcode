package main

import "fmt"

func fullEval(results ...bool) {
	passed := 0
	for _, result := range results {
		if result {
			passed++
		}
	}

	fmt.Printf("%v/%v TEST CASES PASSED\n", passed, len(results))
}

func eval(inputs interface{}, expected interface{}, outputs interface{}) bool {
	fmt.Printf("inputs=%v,expected=%v,outputs=%v\n", inputs, expected, outputs)

	if outputs == expected {
		fmt.Printf("TEST CASE PASSED\n\n")
		return true
	} else {
		fmt.Printf("TEST CASE FAILED\n\n")
		return false
	}
}
