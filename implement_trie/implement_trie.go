package implement_trie

type TrieNode struct {
	val   int
	child [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: new(TrieNode)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	u := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if u.child[idx] == nil {
			u.child[idx] = new(TrieNode)
		}
		u = u.child[idx]
	}
	u.val++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	u := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if u.child[idx] == nil {
			return false
		}
		u = u.child[idx]
	}
	if u.val > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	u := this.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if u.child[idx] == nil {
			return false
		}
		u = u.child[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
