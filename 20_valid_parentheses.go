package main

func test_20() {
	eval("()", true, isValid("()"))
	eval("()[]{}", true, isValid("()[]{}"))
}

func isValid(s string) bool {
	roundStart := rune('(')
	roundEnd := rune(')')
	curlyStart := rune('{')
	curlyEnd := rune('}')
	squareStart := rune('[')
	squareEnd := rune(']')

	var queue []rune = make([]rune, len(s))
	queue = queue[:0]

	for _, symbol := range s {
		if symbol == roundStart || symbol == curlyStart || symbol == squareStart {
			queue = append(queue, symbol)
		} else {
			if len(queue) == 0 {
				return false
			}

			x := queue[len(queue)-1]
			if (symbol == roundEnd && x != roundStart) ||
				(symbol == curlyEnd && x != curlyStart) ||
				(symbol == squareEnd && x != squareStart) {
				return false
			}

			queue = queue[:len(queue)-1]
		}
	}

	return len(queue) == 0
}
