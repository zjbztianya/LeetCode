package problem0074

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	target int
	ans    bool
}{

	// 可以有多个 testcase
}

func Test_searchMatrix(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, searchMatrix(tc.matrix, tc.target), "输入:%v", tc)
	}
}

func Benchmark_searchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			searchMatrix(tc.matrix, tc.target)
		}
	}
}
