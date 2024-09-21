package coursescheduleii

type WordDictionary struct {
	Trie   map[byte]WordDictionary
	IsWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		Trie:   make(map[byte]WordDictionary),
		IsWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.IsWord = true
		return
	}

	childTrie, ok := this.Trie[word[0]]
	if !ok {
		childTrie = Constructor()
	}
	childTrie.AddWord(word[1:])
	this.Trie[word[0]] = childTrie
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.IsWord
	}

	childTrie, ok := this.Trie[word[0]]
	if ok {
		return childTrie.Search(word[1:])
	}

	if word[0] == '.' {
		for _, childTrie := range this.Trie {
			if childTrie.Search(word[1:]) {
				return true
			}
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
