package main

func test_10() {
	fullEval(
		eval([]string{"a", "a"}, true, isMatch("a", "a")),
		eval([]string{"aa", "a"}, false, isMatch("aa", "a")),
		eval([]string{"aa", "aa"}, true, isMatch("aa", "aa")),
		eval([]string{"a", ".*..a"}, false, isMatch("a", ".*..a")),
		eval([]string{"aa", "a*"}, true, isMatch("aa", "a*")),
		eval([]string{"aaaa", "a*"}, true, isMatch("aaaa", "a*")),
		eval([]string{"ab", ".*"}, true, isMatch("ab", ".*")),
		eval([]string{"aab", "c*a*b*"}, true, isMatch("aab", "c*a*b*")),
		eval([]string{"mississippi", "mis*is*p*."}, false, isMatch("mississippi", "mis*is*p*.")),
		eval([]string{"mississippi", "mis*is*ip*."}, true, isMatch("mississippi", "mis*is*ip*.")),
		eval([]string{"ab", ".*c"}, false, isMatch("ab", ".*c")),
		eval([]string{"aaa", "aaaa"}, false, isMatch("aaa", "aaaa")),
		eval([]string{"aaa", "a*a"}, true, isMatch("aaa", "a*a")),
		eval([]string{"aaa", "a*aa"}, true, isMatch("aaa", "a*aa")),
		eval([]string{"aaab", "a*a"}, false, isMatch("aaab", "a*a")),
		eval([]string{"aaa", "ab*a*c*a"}, true, isMatch("aaa", "ab*a*c*a")),
		eval([]string{"aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*"}, true, isMatch("aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*")),
	)
}

func isMatch(s string, p string) bool {
	var i int = 0
	var pi int = 0

	if pi+1 < len(p) && p[pi+1] == byte('*') {
		for i < len(s) {
			if isMatch(s[i:], p[pi+2:]) {
				return true
			} else if s[i] != p[pi] && p[pi] != byte('.') {
				break
			}
			i++
		}

		pi += 2
		for i == len(s) && pi < len(p) && pi+1 < len(p) {
			if p[pi+1] == byte('*') {
				pi += 2
			} else {
				break
			}
		}
	}

	if i == len(s) && pi == len(p) {
		return true
	} else if (i != len(s) && pi == len(p)) || (i == len(s) && pi != len(p)) {
		return false
	}

	if s[i] != p[pi] && p[pi] != byte('.') {
		return false
	}

	return isMatch(s[i+1:], p[pi+1:])
}
