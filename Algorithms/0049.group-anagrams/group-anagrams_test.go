package problem0049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	strs []string
	ans  [][]string
}{

	// 可以有多个 testcase
}

func Test_groupAnagrams(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, groupAnagrams(tc.strs), "输入:%v", tc)
	}
}

func Benchmark_groupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			groupAnagrams(tc.strs)
		}
	}
}
