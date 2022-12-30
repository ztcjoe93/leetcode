func rotate(matrix [][]int) {
	times := len(matrix) / 2
	// we swap TL with TR, TL with BL, BL with BR do complete a rotation
	// repeat for each element
	for i := 0; i < times; i++ {
		elements := len(matrix[0]) - 1 - (i * 2)
		for j := 0; j < elements; j++ {
			tl := &matrix[i][j+i]
			tr := &matrix[i+j][len(matrix[i])-(1+i)]
			bl := &matrix[len(matrix)-j-i-1][i]
			br := &matrix[len(matrix)-i-1][len(matrix[i])-j-i-1]
			fmt.Printf("On element %v\n", j+1)
			fmt.Printf("Swapping %v with %v\n", *tl, *tr)
			*tl, *tr = *tr, *tl
			fmt.Printf("Swapping %v with %v\n", *tl, *bl)
			*tl, *bl = *bl, *tl
			fmt.Printf("Swapping %v with %v\n", *bl, *br)
			*bl, *br = *br, *bl
		}
	}
}
