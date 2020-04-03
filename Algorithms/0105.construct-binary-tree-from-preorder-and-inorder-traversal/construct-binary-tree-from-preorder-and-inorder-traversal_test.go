package problem0105

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	preorder []int
	inorder  []int
	ans      *TreeNode
}{

	// 可以有多个 testcase
}

func Test_buildTree(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, buildTree(tc.preorder, tc.inorder), "输入:%v", tc)
	}
}

func Benchmark_buildTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			buildTree(tc.preorder, tc.inorder)
		}
	}
}
