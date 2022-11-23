package main

import "fmt"

func eval(inputs interface{}, expected interface{}, outputs interface{}) {
	fmt.Printf("inputs=%v,expected=%v,outputs=%v\n", inputs, expected, outputs)

	if outputs == expected {
		fmt.Printf("TEST CASE PASSED\n\n")
	} else {
		fmt.Printf("TEST CASE FAILED\n\n")
	}
}
