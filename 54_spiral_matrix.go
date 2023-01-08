package main

func spiralOrder(matrix [][]int) []int {
	/*
		duplicate matrix that indicates if the cell has been visited,
		slice of direction variable [R, D, L, U]

		start cursor at [0][0]

		while direction checks < 4
			check next cell in direction can be visited || reached end of array
				yes,
					reset directional checks to 0
					add value to array
					move cursor at direction
				no,
					change direction to next in list
					increment directional checks
	*/

	// initialize matrix to track which cells are visited
	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	// we start at the first cell, hence it is always true
	visited[0][0] = true

	// directions slice which acts as a circular list
	directions := []string{"R", "D", "L", "U"}
	dirCur := 0
	dcount := 0

	// cursor x & y coordinates
	cx, cy := 0, 0

	// elements of the matrix in spiral
	spiralElements := []int{matrix[cy][cx]}

	for dcount < 4 {

		switch directions[dirCur] {
		case "R":
			if (cx+1 < len(visited[cy])) && !visited[cy][cx+1] {
				dcount = 0
				spiralElements = append(spiralElements, matrix[cy][cx+1])
				visited[cy][cx+1] = true
				cx++
			} else {
				dirCur++
				dcount++
			}
		case "D":
			if (cy+1 < len(visited)) && !visited[cy+1][cx] {
				dcount = 0
				spiralElements = append(spiralElements, matrix[cy+1][cx])
				visited[cy+1][cx] = true
				cy++
			} else {
				dirCur++
				dcount++
			}
		case "L":
			if (cx-1 >= 0) && !visited[cy][cx-1] {
				dcount = 0
				spiralElements = append(spiralElements, matrix[cy][cx-1])
				visited[cy][cx-1] = true
				cx--
			} else {
				dirCur++
				dcount++
			}
		case "U":
			if (cy-1 >= 0) && !visited[cy-1][cx] {
				dcount = 0
				spiralElements = append(spiralElements, matrix[cy-1][cx])
				visited[cy-1][cx] = true
				cy--
			} else {
				dirCur = 0
				dcount++
			}
		}
	}

	return spiralElements
}
