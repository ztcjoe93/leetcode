package main

import "fmt"

func test_3() {
	inputs := []string{"abcabcbb", "bbbbb", "pwwkew", "abccaacbacabcbacbacbwuibcqqyibckasbdclkjasgbchjkvawghjkcvqwhcbashjcbkhjasvltjgbqwkljbnwe,nmfxaskdjlfslkjfl"}
	outputs := []int{3, 1, 3, 13}

	passed, failed := 0, 0

	for i, _ := range inputs {
		res := lengthOfLongestSubstring(inputs[i])
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

func lengthOfLongestSubstring(s string) int {
	/*
	   iterate through string
	       if character not in map,
	           add to base string, map, add to count as it's not a repeating character
	       if character in map,
	           check count, if is biggest count so far replace biggest count
	           add to base string, map
	           iterate base string from front to back
	               if character count in map is 2, remove char, count, break out of this
	               else, remove char, count, continue
	   return biggest count
	*/

	m := make(map[rune]int)
	baseStr := ""
	biggestCount := 0
	count := 0

	for _, character := range s {
		_, ok := m[character]
		if !ok {
			baseStr += string(character)
			m[character] = 1
			count++
		} else {
			baseStr += string(character)
			m[character] += 1
			for i, v := range baseStr {
				if m[v] < 2 {
					delete(m, v)
					count--
				} else {
					m[character] -= 1
					baseStr = baseStr[i+1:]
					break
				}
			}
		}

		if count > biggestCount {
			biggestCount = count
		}
	}

	return biggestCount
}
