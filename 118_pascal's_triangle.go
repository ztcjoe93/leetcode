package main

func generate(numRows int) [][]int {
	var allRows [][]int

	for i := 0; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, allRows[i-1][j-1]+allRows[i-1][j])
			}
		}
		allRows = append(allRows, row)
	}

	return allRows
}
