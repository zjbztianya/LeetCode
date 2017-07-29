package implement_trie

import (
	"testing"
)

func initTrie() *Trie{
	words:=[]string{"a","bbb","aa","cc"}
	trie:=Constructor()
	for i:=range words{
		trie.Insert(words[i])
	}
	return &trie
}
func TestTrie_Search(t *testing.T) {
	trie:=initTrie()
	words:=[]string{"a","bbb","aa","cc"}
	for i:=range words{
		if !trie.Search(words[i]){
			t.Errorf("can not find %s in trie",words[i])
		}
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie:=initTrie()
	words:=[]string{"b","bb","a","c", "cc"}
	for i:=range words{
		if !trie.StartsWith(words[i]){
			t.Errorf("can not find %s in trie",words[i])
		}
	}
}
