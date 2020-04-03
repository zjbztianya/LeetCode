package problem0124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  int
}{

	// 可以有多个 testcase
}

func Test_maxPathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxPathSum(tc.root), "输入:%v", tc)
	}
}

func Benchmark_maxPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxPathSum(tc.root)
		}
	}
}
