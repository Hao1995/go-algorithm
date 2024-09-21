package coursescheduleii

type WordDictionary struct {
	Trie   []*WordDictionary
	IsWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		Trie:   make([]*WordDictionary, 26),
		IsWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.IsWord = true
		return
	}

	idx := word[0] - 'a'

	childTrie := this.Trie[idx]
	if childTrie == nil {
		trie := Constructor()
		childTrie = &trie
	}
	childTrie.AddWord(word[1:])
	this.Trie[idx] = childTrie
}

func (this *WordDictionary) Search(word string) bool {
	if this == nil {
		return false
	}

	if len(word) == 0 {
		return this.IsWord
	}

	if word[0] == '.' {
		for _, childTrie := range this.Trie {
			if childTrie.Search(word[1:]) {
				return true
			}
		}
		return false
	}

	idx := word[0] - 'a'
	childTrie := this.Trie[idx]
	if childTrie == nil {
		return false
	}

	return childTrie.Search(word[1:])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
