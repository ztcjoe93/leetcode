package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

func test_1() {

	inputs := []interface{}{
		[]interface{}{[]int{2, 7, 11, 15}, 9},
		[]interface{}{[]int{3, 2, 4}, 6},
		[]interface{}{[]int{3, 3}, 6},
	}

	outputs := []interface{}{
		[]int{0, 1},
		[]int{1, 2},
		[]int{0, 1},
	}

	passed, failed := 0, 0

	for index := range inputs {
		args := inputs[index].([]interface{})
		result := twoSum(args[0].([]int), args[1].(int))
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		if slices.Equal(result, outputs[index].([]int)) {
			fmt.Printf("Test case PASSED input=%v output=%v result=%v\n", args, outputs[index], result)
			passed++
		} else {
			fmt.Printf("Test case FAILED input=%v output=%v result=%v\n", args, outputs[index], result)
			failed++
		}
	}

	fmt.Printf("%v/%v TEST CASES PASSED\n", passed, passed+failed)
}

func twoSum(nums []int, target int) []int {

	/* iterate through int slice
	   check if target - int is in hashmap
	   if found, return value of map[key] (index)
	   if not, add map[key] = index

	   solution is O(n) where n = size of int slice
	*/
	m := make(map[int]int)
	var numIndices []int

	for index, value := range nums {
		otherNum := target - value

		if eleIdx, ok := m[otherNum]; ok {
			numIndices = append(numIndices, index, eleIdx)
			break
		} else {
			m[value] = index
		}
	}

	return numIndices
}
