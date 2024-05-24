package wordsearchii

type TrieNode struct {
	Char   byte
	IsWord bool
	Childs map[byte]*TrieNode
}

func (this *TrieNode) AddWord(word string) {
	currNode := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		_, ok := currNode.Childs[c]
		if !ok {
			currNode.Childs[c] = &TrieNode{
				Char:   c,
				IsWord: false,
				Childs: make(map[byte]*TrieNode),
			}
		}
		currNode = currNode.Childs[c]
	}
	currNode.IsWord = true
}

func findWords(board [][]byte, words []string) []string {
	// Trie
	trieRoot := &TrieNode{
		IsWord: false,
		Childs: make(map[byte]*TrieNode),
	}
	for _, word := range words {
		trieRoot.AddWord(word)
	}

	// For each board
	var ansMap map[string]bool = make(map[string]bool)
	var m, n int = len(board), len(board[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			var visited [][]bool = make([][]bool, m)
			for i := 0; i < m; i++ {
				visited[i] = make([]bool, n)
			}
			dfs(board, r, c, trieRoot, visited, []byte{}, ansMap)
		}
	}

	var ans []string = make([]string, 0, len(ansMap))
	for k := range ansMap {
		ans = append(ans, k)
	}
	return ans
}

func dfs(board [][]byte, r, c int, node *TrieNode, visited [][]bool, str []byte, ansMap map[string]bool) {
	if node == nil {
		return
	}

	if node.IsWord && ansMap[string(str)] == false {
		ansMap[string(str)] = true
	}

	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return
	}

	if visited[r][c] {
		return
	}

	childNode, ok := node.Childs[board[r][c]]
	if !ok {
		return
	}

	visited[r][c] = true

	str = append(str, board[r][c])

	dfs(board, r+1, c, childNode, visited, str, ansMap)
	dfs(board, r, c+1, childNode, visited, str, ansMap)
	dfs(board, r-1, c, childNode, visited, str, ansMap)
	dfs(board, r, c-1, childNode, visited, str, ansMap)

	visited[r][c] = false
	return
}
