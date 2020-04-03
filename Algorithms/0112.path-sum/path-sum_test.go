package problem0112

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	sum  int
	ans  bool
}{

	// 可以有多个 testcase
}

func Test_hasPathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, hasPathSum(tc.root, tc.sum), "输入:%v", tc)
	}
}

func Benchmark_hasPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasPathSum(tc.root, tc.sum)
		}
	}
}
