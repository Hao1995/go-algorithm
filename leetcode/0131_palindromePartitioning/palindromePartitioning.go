package palindromepartitioning

func partition(s string) [][]string {
	var ans [][]string

	var dfs func(i, j int, tmpList []string)
	dfs = func(i, j int, tmpList []string) {
		if !isPalindrome(s[i:j]) {
			return
		}

		tmpList = append(tmpList, s[i:j])

		if j == len(s) {
			ans = append(ans, append([]string{}, tmpList...))
			return
		}

		for length := 1; length <= len(s)-j; length++ {
			dfs(j, j+length, tmpList)
		}
	}

	for length := 1; length <= len(s); length++ {
		dfs(0, length, []string{})
	}

	return ans
}

func isPalindrome(s string) bool {
	var i, j int = 0, len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
