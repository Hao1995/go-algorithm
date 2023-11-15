package lexicalorder

func lexicalOrder(n int) []int {
	var res []int
	for i := 1; i < 10; i++ {
		res = append(res, dfs(i, n)...)
	}
	return res
}

func dfs(i int, n int) []int {
	if i > n {
		return []int{}
	}

	var res []int = []int{i}
	for j := i * 10; j < (i+1)*10 && j <= n; j++ {
		res = append(res, dfs(j, n)...)
	}

	return res
}
