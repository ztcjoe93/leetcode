package main

func test_13() {
	fullEval(
		eval("III", 3, romanToInt("III")),
		eval("LVIII", 58, romanToInt("LVIII")),
		eval("MCMXCIV", 1994, romanToInt("MCMXCIV")),
	)
}

func romanToInt(s string) int {
	total := 0

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'V' || s[i] == 'X' {
			m['I'] *= -1
		} else if s[i] == 'L' || s[i] == 'C' {
			m['X'] *= -1
		} else if s[i] == 'D' || s[i] == 'M' {
			m['C'] *= -1
		}

		total += m[s[i]]
	}

	return total
}
