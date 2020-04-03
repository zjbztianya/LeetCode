package problem0114

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans
}{

	// 可以有多个 testcase
}

func Test_flatten(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, flatten(tc.root), "输入:%v", tc)
	}
}

func Benchmark_flatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			flatten(tc.root)
		}
	}
}
