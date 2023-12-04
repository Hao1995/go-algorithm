/*
208. Implement Trie (Prefix Tree)
https://leetcode.com/problems/implement-trie-prefix-tree/description/

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/
package trie

type Trie struct {
	nextNodes map[byte]*Node
}

type Node struct {
	Char      byte
	NextNodes map[byte]*Node
	IsAWord   bool
}

func Constructor() Trie {
	return Trie{
		nextNodes: make(map[byte]*Node),
	}
}

func (this *Trie) Insert(word string) {
	current := this.nextNodes
	for i := 0; i < len(word); i++ {
		c := word[i]
		nextNode, ok := current[c]
		if !ok {
			nextNode = &Node{
				Char:      c,
				NextNodes: make(map[byte]*Node),
				IsAWord:   false,
			}
			current[c] = nextNode
		}

		if i == len(word)-1 {
			nextNode.IsAWord = true
		}

		current = nextNode.NextNodes
	}
}

func (this *Trie) searchHelper(word string, checkIsWord bool) bool {
	current := this.nextNodes
	for i := 0; i < len(word); i++ {
		c := word[i]
		node, ok := current[c]
		if !ok {
			return false
		}
		if i == len(word)-1 {
			return !checkIsWord || node.IsAWord
		}
		current = node.NextNodes
	}
	return false
}

func (this *Trie) Search(word string) bool {
	return this.searchHelper(word, true)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchHelper(prefix, false)
}
