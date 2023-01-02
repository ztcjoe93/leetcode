func exist(board [][]byte, word string) bool {
	// visit each cell in the matrix to find start of word
	// keep an identical matrix of booleans to indicate which cell has been visited

	// when we reach a matching cell, we invoke the function with the marked boolean matrix to
	// check UDLR cells that are not visited

	// early terminate TRUE if found, else return false at the end of the matrix

	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == word[0] {
				dup := duplicateSlice(visited)
				dup[y][x] = true
				if traverse(x, y, board, dup, word[1:]) {
					return true
				}
			}
		}
	}
	// fmt.Printf("%v\n", visited)

	return false
}

func duplicateSlice(slice [][]bool) [][]bool {
	duplicate := make([][]bool, len(slice))
	for i := range slice {
		duplicate[i] = make([]bool, len(slice[i]))
		copy(duplicate[i], slice[i])
	}

	return duplicate
}

func traverse(x int, y int, board [][]byte, visited [][]bool, word string) bool {
	if len(word) == 0 {
		return true
	}
	// check left most
	if x > 0 {
		if board[y][x-1] == word[0] && !visited[y][x-1] {
			dup := duplicateSlice(visited)
			dup[y][x-1] = true
			if traverse(x-1, y, board, dup, word[1:]) {
				return true
			}
		}
	}

	// check right most
	if x < len(board[0])-1 {
		if board[y][x+1] == word[0] && !visited[y][x+1] {
			dup := duplicateSlice(visited)
			dup[y][x+1] = true
			if traverse(x+1, y, board, dup, word[1:]) {
				return true
			}
		}
	}

	// check top most
	if y > 0 {
		if board[y-1][x] == word[0] && !visited[y-1][x] {
			dup := duplicateSlice(visited)
			dup[y-1][x] = true
			if traverse(x, y-1, board, dup, word[1:]) {
				return true
			}
		}
	}

	// check bottom most
	if y < len(board)-1 {
		if board[y+1][x] == word[0] && !visited[y+1][x] {
			dup := duplicateSlice(visited)
			dup[y+1][x] = true
			if traverse(x, y+1, board, dup, word[1:]) {
				return true
			}
		}
	}

	return false
}