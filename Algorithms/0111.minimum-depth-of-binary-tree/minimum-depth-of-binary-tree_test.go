package problem0111

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

func Test_minDepth(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minDepth(tc.root), "输入:%v", tc)
	}
}

func Benchmark_minDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDepth(tc.root)
		}
	}
}
