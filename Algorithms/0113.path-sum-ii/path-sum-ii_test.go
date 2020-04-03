package problem0113

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	sum  int
	ans  [][]int
}{

	// 可以有多个 testcase
}

func Test_pathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, pathSum(tc.root, tc.sum), "输入:%v", tc)
	}
}

func Benchmark_pathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pathSum(tc.root, tc.sum)
		}
	}
}
