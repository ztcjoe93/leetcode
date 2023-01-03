func uniquePaths(m int, n int) int {
	board := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
	}

	for y, _ := range board {
		for x, _ := range board[y] {
			if y == 0 || x == 0 {
				board[y][x] = 1
			} else {
				board[y][x] = board[y-1][x] + board[y][x-1]
			}
		}
	}

	return board[m-1][n-1]
}