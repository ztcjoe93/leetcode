package main

import "fmt"

func test_5() {

	inputs := []string{
		"aaaaa",
		"aaaa",
		"ccc",
		"bcb",
		"cbb",
		"cba",
		"caba",
		"babad",
		"cbbd",
		"cb",
		"bb",
		"a",
		"ccbracecarpopracecarbcc",
		"ccbracecarpopracebcc",
		"azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc",
	}

	outputs := []string{
		"aaaaa",
		"aaaa",
		"ccc",
		"bcb",
		"bb",
		"c",
		"aba",
		"bab",
		"bb",
		"c",
		"bb",
		"a",
		"ccbracecarpopracecarbcc",
		"ecarpoprace",
		"sooos",
	}

	passed, failed := 0, 0
	for i, _ := range inputs {
		res := longestPalindrome(inputs[i])
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

func longestPalindrome(s string) string {
	/*
		if all letters are the same, both odd/even combinations will be a valid palindrome
		i.e oooo, ooooo
		however, if the letters are not, we have to check
		1: odd palindrome
			that the left character matches the right character from the middle (rac<-e->car) at each iteration
		2: even palindrome
			that the left cursor matches the right cursor (red<-->der) at each iteration
		we can't use this 1/2 algorithm without first checking for repeated characters as it's impossible to determine where it ends

		bazoozk
		left - 2, right - 5
		5+1 - 2 = 4


		leftIdx, rightIdx
		iterate curIdx through string till len-1,
			check if string[curIdx] == string[curIdx+1]
			size = 0
			leftIdx = curIdx
			rightIdx = curIdx
			while curIdx < len(string)-1 && string[curIdx] == string[curIdx+1]:
				rightIdx = curIdx+1
				curIdx++
			while leftIdx > 0 && rightIdx < len(string)-1 && string[leftIdx] == string[rightIdx]:
				leftIdx--
				rightIdx++
			if (rightIdx - leftIdx + 1) < len(longestSubStr):
				longestSubStr = string[leftIdx : rightIdx+1]

		   iterate through string
		       check if next rune is the same (for a even-palindrome)
		           if yes, we expand from current index backwards, next index forwards
		               start checking <->
		           if no, check if previous index == forward index
		               if yes, start checking <->
		               if no, move to next rune

		   checking <-> // expand(backCur, forwardCur, str)
		       if backCur == 0 (reached front of string) or forwardCur == len(s)-1 (back of string)
		           if forwardCur - backCur + 1 > longestSize // check that it's longer ;)
		               longestSubStr = s[backCur:forwardCur]
		       if s[backCur-1] == s[forwardCur+1],
		           expand(backCur-1, forwardCur+1, str)
	*/

	if len(s) == 1 {
		return s
	}

	leftIdx, rightIdx, longestSubStr := 0, 0, ""

	for curIdx := 0; curIdx < len(s)-1; curIdx++ {
		leftIdx = curIdx
		rightIdx = curIdx
		for curIdx < len(s)-1 && s[curIdx] == s[curIdx+1] {
			rightIdx = curIdx + 1
			curIdx++
		}

		for leftIdx > 0 && rightIdx < len(s)-1 && s[leftIdx-1] == s[rightIdx+1] {
			leftIdx--
			rightIdx++
		}

		if (rightIdx - leftIdx + 1) > len(longestSubStr) {
			longestSubStr = s[leftIdx : rightIdx+1]
		}
	}

	return longestSubStr
}
