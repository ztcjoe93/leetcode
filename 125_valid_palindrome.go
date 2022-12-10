package main

import "fmt"

func isPalindrome(s string) bool {
	fcur := 0
	bcur := len(s) - 1
	found := false

	for fcur < bcur && bcur > fcur {
		var frontVal byte
		var backVal byte

		for !found {
			if s[fcur] >= 'a' && s[fcur] <= 'z' {
				frontVal = s[fcur]
				found = true
			} else if s[fcur] >= 'A' && s[fcur] <= 'Z' {
				frontVal = s[fcur] + 32
				found = true
			} else {
				fcur++
			}
		}

		found = false
		if s[bcur] >= 'a' && s[bcur] <= 'z' {
			backVal = s[bcur]
			found = true
		} else if s[bcur] >= 'A' && s[bcur] <= 'Z' {
			backVal = s[bcur] + 32
			found = true
		} else {
			bcur--
		}

		fmt.Printf("frontcur-%v | backcur-%v\n", fcur, bcur)
		fmt.Printf("frontcur-%v | backcur-%v\n", s[fcur], s[bcur])
		if frontVal != backVal {
			fmt.Printf("%v <-> %v\n", string(frontVal), string(backVal))
			return false
		} else {
			fcur++
			bcur--
		}
	}

	return true
}
