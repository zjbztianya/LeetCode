package replace_words

import (
	"strings"
)

type TrieNode struct {
	val   int
	child [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{root: new(TrieNode)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	u := this.root
	for _, ch := range word {
		idx:=ch -'a'
		if u.child[idx] == nil {
			u.child[idx] = new(TrieNode)
		}
		u = u.child[idx]
	}
	u.val++
}

/** Returns shortest prefix if the word is in the trie. */
func (this *Trie) SearchPrefix(word string) string {
	u := this.root
	ans := ""
	for i, ch := range word {
		idx:=ch -'a'
		if u.child[idx] == nil {
			break
		}
		u = u.child[idx]
		if u.val > 0 {
			ans = word[:i+1]
			break
		}
	}
	return ans
}

func replaceWords(dict []string, sentence string) string {
	sentences := strings.Split(sentence, " ")
	trie := NewTrie()
	for i := range dict {
		trie.Insert(dict[i])
	}
	var words []string
	for i := range sentences {
		word := trie.SearchPrefix(sentences[i])
		if word != "" {
			words = append(words, word)
		} else {
			words = append(words,sentences[i])
		}
	}
	return strings.Join(words, " ")
}
