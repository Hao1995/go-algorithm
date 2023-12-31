package wordbreak

import "fmt"

type TrieNode struct {
	Next  []*TrieNode
	IsEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Next:  make([]*TrieNode, 26),
		IsEnd: false,
	}
}

func wordBreak(s string, wordDict []string) bool {
	// init trie
	root := NewTrieNode()
	for _, word := range wordDict {
		var node *TrieNode = root
		for i := 0; i < len(word); i++ {
			c := word[i]
			next := node.Next
			if next[c-'a'] == nil {
				next[c-'a'] = NewTrieNode()
			}
			node = next[c-'a']
		}
		node.IsEnd = true
	}
	// printTrie(root, []byte{})

	failedCache := make([]bool, len(s))
	return dfs(s, 0, root, failedCache)
}

func dfs(s string, start int, trie *TrieNode, failedCache []bool) bool {
	if start == len(s) {
		return true
	}

	if failedCache[start] {
		return false
	}

	var node *TrieNode = trie
	for i := start; i < len(s); i++ {
		c := s[i]
		if node.Next[c-'a'] != nil {
			node = node.Next[c-'a']
			if node.IsEnd && dfs(s, i+1, trie, failedCache) {
				return true
			}
		} else {
			break
		}
	}

	failedCache[start] = true
	return false
}

func printTrie(trie *TrieNode, prev []byte) {
	if trie == nil {
		return
	}

	for idx, node := range trie.Next {
		if node != nil {
			c := byte('a' + idx)
			if node.IsEnd {
				for _, prevc := range prev {
					fmt.Printf("%v", string(prevc))
				}
				fmt.Printf("%v, ", string(c))
			}
			printTrie(node, append(prev, c))
		}
	}
}

// s = "catsandog", wordDict = ["cats", "ca", "ts", "and", "cat"]
// "cats", "a", ["n","b"], is not end
// "catsa", "o" is not end >> break
