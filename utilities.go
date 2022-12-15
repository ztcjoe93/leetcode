package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

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

	if reflect.DeepEqual(outputs, expected) {
		fmt.Printf("TEST CASE PASSED\n\n")
		return true
	} else {
		fmt.Printf("TEST CASE FAILED\n\n")
		return false
	}
}

// generate a slice of int with >= slb and <= sub size
// value is scoped using vlb and vub
func generateIntSlice(slb int, sub int, vlb int, vub int) []int {
	if sub == 0 {
		sub = 10
	}

	size := slb + rand.Intn(sub-slb)
	arr := make([]int, size)
	arr = arr[:0]

	for i := 0; i < size; i++ {
		arr = append(arr, vlb+rand.Intn(vub-vlb))
	}

	return arr
}
