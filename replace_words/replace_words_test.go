package replace_words

import (
	"fmt"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	want := "the cat was rat by the bat"
	ans := replaceWords(dict, sentence)
	fmt.Println(ans)
	if ans != want {
		t.Errorf("replaceWords(%q, %q)=%q, ans=%q", dict, sentence, want, ans)
	}

}
