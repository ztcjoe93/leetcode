func generateParenthesis(n int) []string {
	combinations := make([]string, 0)
	dfs(0, 0, "", n, &combinations)

	return combinations
}

func dfs(left int, right int, str string, n int, combinations *[]string) {
	if len(str) == n*2 {
		*combinations = append(*combinations, str)
		return
	}

	if left < n {
		dfs(left+1, right, str+"(", n, combinations)
	}

	if right < left {
		dfs(left, right+1, str+")", n, combinations)
	}
}