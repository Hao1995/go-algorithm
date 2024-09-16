package generateparentheses

func generateParenthesis(n int) []string {
	var combs []string
	backtracking(n, 0, 0, "", &combs)
	return combs
}

func backtracking(n int, open, close int, comb string, combs *[]string) {
	if open == n && close == n {
		*combs = append(*combs, comb)
		return
	}

	if open < n {
		backtracking(n, open+1, close, comb+"(", combs)
	}

	if close < open && close < n {
		backtracking(n, open, close+1, comb+")", combs)
	}

	return
}
