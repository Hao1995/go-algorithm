package generateparentheses

func generateParenthesis(n int) []string {
	store := []int{n, n}
	var ans []string
	help(n, store, "", &ans)
	return ans
}

func help(n int, store []int, tmp string, ans *[]string) {

	// insert ans
	if len(tmp) == n*2 {
		*ans = append(*ans, tmp)
		return
	}

	// put an open parenthesis
	if store[0] > 0 && store[0] <= store[1] {
		newStore := make([]int, 2)
		copy(newStore, store)
		newStore[0]--
		help(n, newStore, tmp+"(", ans)
	}

	// put a close parenthesis
	if store[1] > 0 && store[1] >= store[0] {
		newStore := make([]int, 2)
		copy(newStore, store)
		newStore[1]--
		help(n, newStore, tmp+")", ans)
	}
}
