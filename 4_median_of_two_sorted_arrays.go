package main

import "fmt"

func test_4() {
	inputs := []interface{}{
		[][]int{{1, 3}, {2}},
		[][]int{{1, 2}, {3, 4}},
		[][]int{{2, 4, 8, 10, 15, 23}, {3, 3, 6, 10, 19, 33}},
	}

	outputs := []float64{
		2.0,
		2.5,
		9.0,
	}

	passed, failed := 0, 0
	for i, _ := range inputs {
		res := findMedianSortedArrays(inputs[i].([][]int)[0], inputs[i].([][]int)[1])
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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
	   to retrieve median of both arrays

	   get size of both arrays / 2 == median
	       if size % 2 == 0, even number == index of median + next
	       else, index of median

	   initialize medianCount at median
	   isEven initialization
	   initialize arr1 cursor, arr2 cursor at 0,


	   while medianCount <= index of median
	       check via arr1 cursor value against arr2 cursor value,
	       move cursor for smaller value, add to arr cursor, add 1 to medianCount

	   initialize medianValue
	   if isEven
	       iterate 2 times
	           check via arr1 cursor value against arr2 cursor value,
	           move cursor for smaller value, add value to medianValue
	       medianValue/2
	   else
	       check via arr1 cursor value against arr2 cursor value,
	       set medianValue

	   return medianValue
	*/

	totalLength := len(nums1) + len(nums2)
	isEven := totalLength%2 == 0
	medianCount := totalLength / 2

	cur1, cur2 := 0, 0

	medianValue := 0.0
	for medianCount > 0 {
		medianCount--
		if len(nums1) == cur1 {
			if medianCount == 0 {
				medianValue += float64(nums2[cur2])
			}
			cur2++
		} else if len(nums2) == cur2 {
			if medianCount == 0 {
				medianValue += float64(nums1[cur1])
			}
			cur1++
		} else {
			if nums1[cur1] < nums2[cur2] {
				if medianCount == 0 {
					medianValue += float64(nums1[cur1])
				}
				cur1++
			} else {
				if medianCount == 0 {
					medianValue += float64(nums2[cur2])
				}
				cur2++
			}
		}
	}

	if !isEven {
		medianValue = 0
	}

	if len(nums1) == cur1 {
		medianValue += float64(nums2[cur2])
	} else if len(nums2) == cur2 {
		medianValue += float64(nums1[cur1])
	} else {
		if nums1[cur1] < nums2[cur2] {
			medianValue += float64(nums1[cur1])
		} else {
			medianValue += float64(nums2[cur2])
		}
	}

	if isEven {
		medianValue /= 2
	}

	return medianValue
}
