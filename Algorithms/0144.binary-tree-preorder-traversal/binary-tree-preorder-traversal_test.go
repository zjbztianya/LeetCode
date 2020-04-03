package problem0144

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []int
}{

	// 可以有多个 testcase
}

func Test_preorderTraversal(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, preorderTraversal(tc.root), "输入:%v", tc)
	}
}

func Benchmark_preorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			preorderTraversal(tc.root)
		}
	}
}
