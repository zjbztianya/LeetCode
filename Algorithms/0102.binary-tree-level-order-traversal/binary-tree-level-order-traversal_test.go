package problem0102

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  [][]int
}{

	// 可以有多个 testcase
}

func Test_levelOrder(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, levelOrder(tc.root), "输入:%v", tc)
	}
}

func Benchmark_levelOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			levelOrder(tc.root)
		}
	}
}
