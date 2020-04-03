package problem0090

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  [][]int
}{

	// 可以有多个 testcase
}

func Test_subsetsWithDup(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, subsetsWithDup(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_subsetsWithDup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subsetsWithDup(tc.nums)
		}
	}
}
