package problem0095

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans []*TreeNode
}{

	// 可以有多个 testcase
}

func Test_generateTrees(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, generateTrees(tc.n), "输入:%v", tc)
	}
}

func Benchmark_generateTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			generateTrees(tc.n)
		}
	}
}
