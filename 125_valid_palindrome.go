package main

func isPalindrome(s string) bool {
	isPalindrome := true
	frontCur := 0
	backCur := len(s) - 1
	for frontCur <= backCur {
		var frontVal byte
		var backVal byte

		for frontCur < len(s) {
			if s[frontCur] >= 'A' && s[frontCur] <= 'Z' {
				frontVal = s[frontCur] + 32
				break
			} else if (s[frontCur] >= 'a' && s[frontCur] <= 'z') || (s[frontCur] >= '0' && s[frontCur] <= '9') {
				frontVal = s[frontCur]
				break
			} else {
				frontCur++
			}
		}

		for backCur >= 0 {
			if s[backCur] >= 'A' && s[backCur] <= 'Z' {
				backVal = s[backCur] + 32
				break
			} else if (s[backCur] >= 'a' && s[backCur] <= 'z') || (s[backCur] >= '0' && s[backCur] <= '9') {
				backVal = s[backCur]
				break
			} else {
				backCur--
			}
		}

		if frontVal != backVal {
			return false
		} else {
			frontCur++
			backCur--
		}
	}

	return isPalindrome
}
